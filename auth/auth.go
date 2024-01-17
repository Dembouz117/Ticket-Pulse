package main

import (
	"auth/ent"
	"auth/ent/user"
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"entgo.io/ent/dialect"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	redisAddress       string
	postgresAddress    string
	secretKey          string
	cookieDomain       string
	awsAccessKeyID     string
	awsSecretAccessKey string
	environment        string
	ctx                = context.Background()
	redisClient        *redis.Client
	awsSession         *session.Session
)

func main() {
	godotenv.Load("../.env")
	redisAddress = os.Getenv("REDIS_ADDRESS")
	postgresAddress = os.Getenv("POSTGRES_ADDRESS")
	secretKey = os.Getenv("SECRET_KEY")
	cookieDomain := os.Getenv("COOKIE_DOMAIN")
	if cookieDomain == "" {
		cookieDomain = "localhost"
	}
	awsAccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	environment = os.Getenv("ENVIRONMENT")

	opts, err := redis.ParseURL(redisAddress)

	if err != nil {
		panic(err)
	}

	redisClient = redis.NewClient(opts)

	awsSession = session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""),
	}))

	r := gin.Default()

	// Allow CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://frontend:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/auth/login", loginHandler)
	r.POST("/auth/otp", otpHandler)
	r.GET("/auth/user", authMiddleware(), getUserInfoHandler)
	r.POST("/auth/user", createUserHandler)
	r.POST("/auth/admin", createAdminHandler)
	r.DELETE("/auth/user", authMiddleware(), deleteUserHandler)
	r.POST("/auth/logout", authMiddleware(), logoutHandler)
	r.GET("/auth/checkLogin", authMiddleware(), checkLoginHandler)
	r.GET("/auth/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Authentication microservice is running"})
	})

	client, err := ent.Open(dialect.Postgres, postgresAddress)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
		return
	}

	defer client.Close()

	fmt.Println("Authentication microservice is running on :8080")
	r.Run(":8080")
}

func loginHandler(c *gin.Context) {
	var requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println("Invalid JSON request")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON request"})
		return
	}

	phone, err := checkUsernameAndPasswordReturnNumber(requestBody.Email, requestBody.Password)

	if err != nil || phone == "" {
		log.Println("Invalid credentials")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	otp := generateOTP(6)
	addToRedis("otp:"+requestBody.Email, otp, 60)
	sendOTP(phone, otp)

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent to user"})
}

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

		userID, ok := claims["userId"].(string)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid User ID in JWT claims"})
			c.Abort()
			return
		}

		email, ok := claims["email"].(string)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Email in JWT claims"})
			c.Abort()
			return
		}

		// Check if the token is in Redis
		_, err = redisClient.Get(ctx, "token:"+jwtToken).Result()

		if err != nil {
			if err == redis.Nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired JWT token"})
			} else {
				log.Fatalf("Error getting token from Redis for %s: %v", userID, err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Error checking JWT token"})
			}
			c.Abort()
			return
		}

		c.Set("userId", userID) // Store the user id in the context for later use
		c.Set("email", email)   // Store the email in the context for later use
		c.Next()
	}
}

func otpHandler(c *gin.Context) {
	var requestBody struct {
		Email   string `json:"email"`
		OtpCode string `json:"otpCode"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println("Invalid JSON request")
		log.Println(requestBody.Email + " " + requestBody.OtpCode)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON request"})
		return
	}

	storedOTP, err := redisClient.Get(ctx, "otp:"+requestBody.Email).Result()

	if err != nil {
		if err == redis.Nil {
			log.Printf("OTP not found for %s", requestBody.Email)
			c.JSON(http.StatusUnauthorized, gin.H{"message": "OTP expired"})
		} else {
			log.Fatalf("Error getting OTP from Redis for %s: %v", requestBody.Email, err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error checking OTP"})
		}
		return
	}

	if requestBody.OtpCode != storedOTP {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid OTP"})
		return
	}

	// Get user role
	userInfo, err := getUserInfoByEmail(requestBody.Email)
	if err != nil {
		log.Printf("Error fetching user information: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching user information"})
		return
	}

	// put user role into jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userInfo.ID,
		"role":   userInfo.Role,
		"email":  userInfo.Email,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		log.Fatalf("Error generating JWT token for %s: %v", requestBody.Email, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error generating JWT token"})
		return
	}

	cookieMaxAge := 60 * 60 * 24 // 1 day
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("token", tokenString, cookieMaxAge, "/", cookieDomain, false, true)
	addToRedis("token:"+tokenString, "", cookieMaxAge)

	c.JSON(http.StatusOK, gin.H{"message": "Authentication successful"})
}

func createUserHandler(c *gin.Context) {
	var requestBody struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Phone    string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println("Invalid JSON request")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON request"})
		return
	}

	if !isValidEmail(requestBody.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email format"})
		return
	} else if !isValidPassword(requestBody.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character"})
		return
	} else if !isValidPhoneNumber(requestBody.Phone) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid phone number format"})
		return
	}

	client, err := ent.Open(dialect.Postgres, postgresAddress)

	if err != nil {
		log.Printf("failed opening connection to postgres: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating user"})
		return
	}

	defer client.Close()

	// Check if the email already exists in the database
	exists, err := client.User.Query().Where(user.EmailEQ(requestBody.Email)).Exist(ctx)

	if err != nil {
		log.Printf("failed querying user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating user"})
		return
	}

	if exists {
		log.Printf("Email %s already exists", requestBody.Email)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email already exists"})
		return
	}

	// Hash and salt the password
	// bcrypt generates a random salt and stores together
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Printf("failed hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating user"})
		return
	}

	_, err = client.User.
		Create().
		SetName(requestBody.Name).
		SetEmail(requestBody.Email).
		SetPassword(string(hashedPassword)).
		SetPhone(requestBody.Phone).
		SetRole("user").
		Save(ctx)

	if err != nil {
		log.Printf("failed creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func createAdminHandler(c *gin.Context) {
	var requestBody struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Phone    string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println("Invalid JSON request")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON request"})
		return
	}

	if !isValidEmail(requestBody.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email format"})
		return
	} else if !isValidPassword(requestBody.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character"})
		return
	} else if !isValidPhoneNumber(requestBody.Phone) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid phone number format"})
		return
	}

	client, err := ent.Open(dialect.Postgres, postgresAddress)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating user"})
		return
	}

	defer client.Close()

	// Check if the email already exists in the database
	exists, err := client.User.Query().Where(user.EmailEQ(requestBody.Email)).Exist(ctx)

	if err != nil {
		log.Fatalf("failed querying user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating user"})
		return
	}

	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email already exists"})
		return
	}

	// Hash and salt the password
	// bcrypt generates a random salt and stores together
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Printf("failed hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating user"})
		return
	}

	_, err = client.User.
		Create().
		SetName(requestBody.Name).
		SetEmail(requestBody.Email).
		SetPassword(string(hashedPassword)).
		SetPhone(requestBody.Phone).
		SetRole("admin").
		Save(ctx)

	if err != nil {
		log.Fatalf("failed creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func getUserInfoHandler(c *gin.Context) {
	// Retrieve the email from the context, which was set in the authMiddleware
	userIDStr, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to retrieve user information"})
		return
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid User ID in context"})
		return
	}

	// Assuming you have a function to fetch user information from your database
	userInfo, err := getUserInfoByUserID(userID)
	if err != nil {
		log.Printf("Error fetching user information: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching user information"})
		return
	}

	// Return the user information as JSON
	c.JSON(http.StatusOK, gin.H{"user": userInfo})
}

func deleteUserHandler(c *gin.Context) {
	// Retrieve the user's email from the context
	userIDStr, exists := c.Get("userId")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "User ID not found in context"})
		return
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid User ID in context"})
		return
	}

	client, err := ent.Open(dialect.Postgres, postgresAddress)

	if err != nil {
		log.Printf("Error opening Ent client: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting user"})
		return
	}

	defer client.Close()

	// Look up the user by email
	u, err := client.User.Query().
		Where(user.ID(userID)).
		First(ctx)

	if err != nil {
		log.Printf("Error querying user by User ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting user"})
		return
	}

	if u == nil {
		log.Printf("User with User ID %s not found for deletion.", userID)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting user"})
		return
	}

	if err := client.User.DeleteOne(u).Exec(ctx); err != nil {
		log.Printf("Error deleting user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting user"})
		return
	}

	log.Printf("User with User ID %s has been deleted.", userID)
	c.JSON(http.StatusOK, gin.H{"message": "User account deleted"})
}

// Remove the token from Redis
func logoutHandler(c *gin.Context) {
	jwtToken, err := c.Cookie("token")

	if err != nil || jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing JWT token"})
		return
	}

	_, err = redisClient.Del(ctx, "token:"+jwtToken).Result()

	if err != nil {
		log.Printf("Error deleting token from Redis: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error logging out"})
		return
	}

	// Clear the cookie in the client by setting its expiration to a past time.
	c.SetCookie("token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// Check if the user is logged in
func checkLoginHandler(c *gin.Context) {
	jwtToken, err := c.Cookie("token")

	if err != nil || jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing JWT token"})
		return
	}

	_, err = redisClient.Get(ctx, "token:"+jwtToken).Result()

	if err != nil {
		if err == redis.Nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired JWT token"})
		} else {
			log.Fatalf("Error getting token from Redis: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error checking JWT token"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User is logged in"})
}

func sendOTP(phone string, otp string) {
	log.Printf("OTP for %s is %s", phone, otp)

	if environment != "production" {
		return
	}

	// Create an SNS client.
	svc := sns.New(awsSession)

	// Define the message and phone number.
	message := "Hello from Ticketing! Your OTP is " + otp
	phoneNumber := phone

	// Send the SMS.
	result, err := svc.Publish(&sns.PublishInput{
		Message:     aws.String(message),
		PhoneNumber: aws.String(phoneNumber),
	})
	if err != nil {
		log.Printf("Error sending SMS: %v", err)
		return
	}

	log.Printf("SMS sent with message ID: %s\n", *result.MessageId)
}

func generateRandomID() string {
	return uuid.New().String()
}

// Check if the username and password are valid and return phone number if valid
func checkUsernameAndPasswordReturnNumber(email string, password string) (string, error) {
	client, err := ent.Open(dialect.Postgres, postgresAddress)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	defer client.Close()

	user, err := client.User.Query().
		Where(user.Email(email)).
		Only(context.Background())

	if err != nil {
		return "", err
	}

	// Compare the input password with the hashed password stored in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	return user.Phone, nil
}

func addToRedis(key string, value string, seconds int) {
	err := redisClient.Set(ctx, key, value, 0).Err()

	if err != nil {
		log.Printf("Error writing to Redis for %s: %v", key, value)
	}

	if seconds > 0 {
		err = redisClient.Expire(ctx, key, time.Duration(seconds)*time.Second).Err()

		if err != nil {
			log.Printf("Error setting expiry for %s: %v", key, value)
		}
	}
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func generateOTP(length int) string {
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	return string(b)
}

func isValidEmail(email string) bool {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(emailPattern).MatchString(email)
}

func isValidPassword(password string) bool {
	// Check for at least one uppercase letter
	upper := regexp.MustCompile("[A-Z]")
	if !upper.MatchString(password) {
		return false
	}

	// Check for at least one lowercase letter
	lower := regexp.MustCompile("[a-z]")
	if !lower.MatchString(password) {
		return false
	}

	// Check for at least one digit
	digit := regexp.MustCompile("[0-9]")
	if !digit.MatchString(password) {
		return false
	}

	// Check for at least one special character
	special := regexp.MustCompile(`[@#$%^&+=!]`)
	if !special.MatchString(password) {
		return false
	}

	// Check for a minimum length of 8 characters
	if len(password) < 8 {
		return false
	}

	// Check for no whitespace
	if regexp.MustCompile(`\s`).MatchString(password) {
		return false
	}

	// If all checks passed, the password is valid
	return true
}

func isValidPhoneNumber(phone string) bool {
	// At least 8 digits, optionally starting with a '+'
	phonePattern := `^\+?\d{8,}$`
	return regexp.MustCompile(phonePattern).MatchString(phone)
}

func getUserInfoByEmail(email string) (*ent.User, error) {
	client, err := ent.Open(dialect.Postgres, postgresAddress)

	if err != nil {
		log.Printf("Error opening Ent client: %v", err)
		return nil, err
	}

	defer client.Close()

	// Query the User entity by email and exclude the fields including password
	user, err := client.User.Query().
		Where(user.Email(email)).
		Select(user.FieldName, user.FieldEmail, user.FieldPhone, user.FieldRole).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			// Handle case where the user was not found.
			return nil, fmt.Errorf("user not found")
		}
		// Handle other errors.
		return nil, err
	}

	return user, nil
}

func getUserInfoByUserID(userID uuid.UUID) (*ent.User, error) {
	client, err := ent.Open(dialect.Postgres, postgresAddress)

	if err != nil {
		log.Printf("Error opening Ent client: %v", err)
		return nil, err
	}

	defer client.Close()

	// Query the User entity by email and exclude the fields including password
	user, err := client.User.Query().
		Where(user.ID(userID)).
		Select(user.FieldName, user.FieldEmail, user.FieldPhone, user.FieldRole).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			// Handle case where the user was not found.
			return nil, fmt.Errorf("user not found")
		}
		// Handle other errors.
		return nil, err
	}

	return user, nil
}
