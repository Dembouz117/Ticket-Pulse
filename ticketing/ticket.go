package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var (
	redisAddress    string
	postgresAddress string
	secretKey       string
	ctx             = context.Background()
	redisClient     *redis.Client

	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func main() {
	godotenv.Load("../.env")
	redisAddress = os.Getenv("REDIS_ADDRESS")
	postgresAddress = os.Getenv("POSTGRES_ADDRESS")
	secretKey = os.Getenv("SECRET_KEY")

	opts, err := redis.ParseURL(redisAddress)

	if err != nil {
		panic(err)
	}

	redisClient = redis.NewClient(opts)

	r := gin.Default()
	r.GET("/tickets", authMiddleware(), viewTicketsHandler)
	r.POST("/join", authMiddleware(), joinWaitingOrQueueRoomHandler)
	r.GET("/queue", authMiddleware(), queueHandler)

	log.Println("Ticketing microservice is running on :8083")
	r.Run(":8083")
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

func viewTicketsHandler(c *gin.Context) {
	// Handle ticket viewing logic here

	c.JSON(http.StatusOK, gin.H{"message": "List of tickets"})
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

		if err != nil {
			log.Printf("Error checking member existence: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		redisClient.SAdd(ctx, "waiting-room", email)

		if err == redis.Nil {
			// generate current timestamp
			timestamp := float64(time.Now().UnixNano())

			// add user to waiting room (important because we use it to determine if user is active in queue)
			redisClient.SAdd(ctx, "waiting-room", email)
			redisClient.ZAdd(ctx, "queue", &redis.Z{Member: email, Score: timestamp})

			log.Printf("Added member to queue: %v", email)
		} else {
			log.Printf("Member already exists in queue: %v", err)
		}

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
