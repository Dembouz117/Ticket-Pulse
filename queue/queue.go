package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var (
	redisAddress string
	secretKey    string
	ctx          = context.Background()
	redisClient  *redis.Client

	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

var runLoopToRemove = false

func main() {
	godotenv.Load("../.env")
	redisAddress = os.Getenv("REDIS_ADDRESS")
	secretKey = os.Getenv("SECRET_KEY")

	opts, err := redis.ParseURL(redisAddress)

	if err != nil {
		panic(err)
	}

	redisClient = redis.NewClient(opts)

	setRedisValue("concert-status", "waiting")

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.GET("/queue/randomise", randomiseHandler)
	r.GET("/queue/reset", resetHandler)

	r.POST("/queue/join", authMiddleware(), joinWaitingOrQueueRoomHandler)
	r.GET("/queue/queue", authMiddleware(), queueHandler)
	r.GET("/queue/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	log.Println("Ticketing microservice is running on :8500")
	r.Run(":8500")
}

// Middleware to check if the user is authenticated
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken, err := c.Cookie("token")

		if err != nil || jwtToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing JWT token"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired JWT token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid JWT claims"})
			c.Abort()
			return
		}

		email, ok := claims["email"].(string)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email in JWT claims"})
			c.Abort()
			return
		}

		c.Set("email", email) // Store the email in the context for later use
		c.Next()
	}
}

func randomiseHandler(c *gin.Context) {
	setRedisValue("concert-status", "sorting")

	generateRandomNameAndAddToRedis(100)
	names := getNamesFromWaitingRoom()
	shuffle(names)

	start := time.Now()

	// Add names to queue
	addNamesToRedisQueue(names)

	elapsed := time.Since(start)

	log.Printf("Added %d elements in %s\n", len(names), elapsed)

	setRedisValue("concert-status", "queueing")

	// Create a sorted set to store the expiration timestamp of each member when they are buying tickets after being selected with
	redisClient.ZAdd(ctx, "permitted-buyers")

	runLoopToRemove = true
	// Start the loop in a goroutine
	go func() {
		for runLoopToRemove {
			time.Sleep(20 * time.Second)
			removeInActiveMember()
			log.Println("Removed inactive members")

			minScore := float64(time.Now().Unix() - 600) // Keep elements from the last 10 minutes
			// Check if the member already exists in the sorted set
			redisClient.ZRemRangeByScore(context.Background(), "permitted-buyers", "-inf", fmt.Sprintf("%.f", minScore))
		}
	}()

	c.JSON(http.StatusOK, gin.H{"message": "Randomise successful"})
}

func resetHandler(c *gin.Context) {
	setRedisValue("concert-status", "waiting")
	redisClient.Del(ctx, "waiting-room")
	redisClient.Del(ctx, "queue")
	redisClient.Del(ctx, "queue:expiration")
	redisClient.Del(ctx, "permitted-buyers")

	runLoopToRemove = false

	c.JSON(http.StatusOK, gin.H{"message": "Reset successful"})
}

func addNamesToRedisQueue(names []string) {
	var zElements []*redis.Z

	timestamp := 1.0

	for _, name := range names {
		z := &redis.Z{Member: name, Score: timestamp}
		zElements = append(zElements, z)

		timestamp += 0.000001
	}

	result, err := redisClient.ZAdd(ctx, "queue", zElements...).Result()

	if err != nil {
		log.Println("Error:", err)
		return
	}

	log.Print(result)
}

func getNamesFromWaitingRoom() []string {
	var names []string

	// Get all names from waiting room
	names, err := redisClient.SMembers(ctx, "waiting-room").Result()
	if err != nil {
		panic(err)
	}

	return names
}

func generateRandomNameAndAddToRedis(size int) {
	var names []string

	for i := 0; i < size; i++ {
		name := gofakeit.Email()
		names = append(names, name)
	}

	redisClient.SAdd(ctx, "waiting-room", names).Result()
}

func setRedisValue(key string, value string) {
	err := redisClient.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

// Shuffle shuffles a string slice randomly
func shuffle(slice []string) {
	// Initialize the random number generator with a unique seed
	// rand.Seed(time.Now().UnixNano())

	// Fisher-Yates shuffle algorithm
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func removeElementsFromQueue() {
	var counter = 0
	// Remove and print elements from the sorted set
	for {
		result, err := redisClient.ZPopMin(ctx, "queue").Result()
		if err != nil {
			log.Println(err)
			return // Exit the function on error
		}

		if len(result) > 0 {
			member := result[0].Member.(string)
			score := result[0].Score

			log.Printf("Name: %s, Score: %f\n", member, score)
		} else {
			log.Println("Queue is empty")
			return // Exit the function if the queue is empty
		}

		counter++
		if counter == 1000 {
			break
		}
	}
}

func removeExpiredMembers(setKey string) {
	currentTime := time.Now().Unix()
	members := redisClient.SMembers(ctx, setKey).Val()

	for _, member := range members {
		expirationTimestamp, _ := redisClient.HGet(ctx, setKey+":expiration", member).Int64()
		if expirationTimestamp <= currentTime {
			redisClient.SRem(ctx, setKey, member)
			redisClient.HDel(ctx, setKey+":expiration", member)
			log.Printf("Removed expired member: %s\n", member)
		}
	}
}

func removeInActiveMember() {
	setValues, err := redisClient.SMembers(ctx, "waiting-room").Result()
	if err != nil {
		log.Println("Error retrieving set values:", err)
		return
	}

	hashKeys, err := redisClient.HKeys(ctx, "queue:expiration").Result()
	if err != nil {
		log.Println("Error retrieving hash values:", err)
		return
	}

	// Convert hash values to a map for easy comparison
	hashKeyMap := make(map[string]struct{})
	for _, val := range hashKeys {
		hashKeyMap[val] = struct{}{}
	}

	// Calculate the difference between the two sets
	var difference []string
	for _, val := range setValues {
		if _, exists := hashKeyMap[val]; !exists {
			difference = append(difference, val)
		}
	}

	log.Println("Difference:", difference)

	if len(difference) > 0 {
		loopCounter := 0

		if len(difference) > 10 {
			loopCounter = 10
		}

		for i := 0; i <= loopCounter; i++ {
			// Get first element from sorted set without removing it
			result, err := redisClient.ZRange(ctx, "queue", 0, 0).Result()
			if err != nil {
				log.Println("Error retrieving sorted set value:", err)
				return
			}

			// Check if the first element is in the difference
			if _, exists := hashKeyMap[result[0]]; !exists {
				// Remove the difference from the sorted set
				redisClient.ZRem(ctx, "queue", result)
				redisClient.SRem(ctx, "waiting-room", result)
				log.Printf("Removed %s from queue\n", result)
			}
		}
	}
}

func joinWaitingOrQueueRoomHandler(c *gin.Context) {
	// Get the user ID from the JWT token
	email, exists := c.Get("email")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Email not found in context"})
		return
	}

	// check if key in redis is "waiting"
	val, err := redisClient.Get(ctx, "concert-status").Result()

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	if val == "waiting" {
		// add user to waiting list
		// add user to redis set
		redisClient.SAdd(ctx, "waiting-room", email)
		c.JSON(http.StatusOK, gin.H{"message": "Added to waiting list"})
		return
	}

	for val == "sorting" {
		time.Sleep(2000)

		val, err = redisClient.Get(ctx, "concert-status").Result()

		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}
	}

	if val == "queueing" {
		// Check if the member already exists in the sorted set
		_, err := redisClient.ZScore(ctx, "queue", email.(string)).Result()

		if err == redis.Nil {
			// generate current timestamp
			timestamp := float64(time.Now().UnixNano())

			// add user to waiting room (important because we use it to determine if user is active in queue)
			redisClient.SAdd(ctx, "waiting-room", email)
			redisClient.ZAdd(ctx, "queue", &redis.Z{Member: email, Score: timestamp})

			log.Printf("Added member to queue: %v", email)
			c.JSON(http.StatusOK, gin.H{"message": "Added to queue"})
			return
		} else if err != nil {
			log.Fatalf("Error checking user existence: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		}

		redisClient.SAdd(ctx, "waiting-room", email)
		log.Printf("Member %v already exists in queue", email)

		c.JSON(http.StatusOK, gin.H{"message": "Added to queue"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Concert is not open yet"})
}

func queueHandler(c *gin.Context) {
	// Get the user email from the JWT token
	email, exists := c.Get("email")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Email not found in context"})
		return
	}

	// Check if member exists in the "queue" sorted set
	_, err := redisClient.ZRank(ctx, "queue", email.(string)).Result()

	if err != nil {
		log.Printf("Error checking member existence: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Member does not exist in queue"})
		return
	}

	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	stop := make(chan struct{})
	defer func() {
		// Ensure the WebSocket connection is closed when the function exits
		conn.Close()
		// Signal the goroutine to stop by closing the stop channel
		close(stop)
	}()

	// Continuously send updates of the position in the queue to the WebSocket client
	go func() {
		for {
			select {
			case <-stop:
				return // Stop the goroutine when the WebSocket connection is closed
			default:
				// Get the position of the user in the queue
				rank, err := redisClient.ZRank(ctx, "queue", email.(string)).Result()

				if err != nil {
					log.Printf("Error checking member existence: %v", err)
					conn.WriteMessage(websocket.TextMessage, []byte("Internal server error"))
					return
				}

				log.Printf("Rank: %v", rank)

				// Send a redirect message to the user if their rank is 0
				if rank == 0 {
					// Check if sorted set exist
					val, err := redisClient.Exists(ctx, "permitted-buyers").Result()

					if err != nil {
						log.Printf("Error checking permitted-buyers existence: %v", err)
						conn.WriteMessage(websocket.TextMessage, []byte("Internal server error"))
						return
					}

					if val == 0 {
						// Remove the user from the "queue" and redirect them to the "buy tickets" page
						redisClient.ZRem(ctx, "queue", email.(string))
						timeNow := float64(time.Now().Unix())
						redisClient.ZAdd(ctx, "permitted-buyers", &redis.Z{Member: email, Score: timeNow})
						conn.WriteMessage(websocket.TextMessage, []byte("Redirect to buy tickets"))
						return
					}

					// Get size of sorted set
					size, err := redisClient.ZCard(ctx, "permitted-buyers").Result()

					if err != nil {
						log.Printf("Error checking size of permitted-buyers: %v", err)
						conn.WriteMessage(websocket.TextMessage, []byte("Internal server error"))
						return
					}

					if size < 50 {
						// Remove the user from the "queue" and redirect them to the "buy tickets" page
						redisClient.ZRem(ctx, "queue", email.(string))
						timeNow := float64(time.Now().Unix())
						redisClient.ZAdd(ctx, "permitted-buyers", &redis.Z{Member: email, Score: timeNow})
						conn.WriteMessage(websocket.TextMessage, []byte("Redirect to buy tickets"))
						return
					}
				}

				// Add the user back to the "queue" with a 60-second expiration
				addMemberWithExpiration("queue", email.(string), 60)

				// Send the position in the queue to the WebSocket client
				msg := fmt.Sprintf("Position in queue is: %d", rank)
				conn.WriteMessage(websocket.TextMessage, []byte(msg))

				// Sleep for a while before sending the next update (e.g., every 5 seconds)
				time.Sleep(5 * time.Second)
			}
		}
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			return // Exit the handler when the WebSocket connection is closed
		}
	}
}

func addMemberWithExpiration(setKey string, member string, expirationSeconds int) {
	// Calculate the expiration timestamp
	expirationTimestamp := time.Now().Add(time.Duration(expirationSeconds) * time.Second).Unix()

	// Store the expiration time in a Redis hash
	redisClient.HSet(ctx, setKey+":expiration", member, expirationTimestamp)
}

func checkIfBuyerIsPermitted(email string) bool {
	_, err := redisClient.ZScore(ctx, "permitted-buyers", email).Result()

	if err != nil {
		log.Printf("Error checking member existence: %v", err)
		return false
	}

	return true
}
