package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// getDetailsFromJWTToken extracts email and role from the JWT token.
func getDetailsFromJWTToken(c *gin.Context, secretKey string) (userID uuid.UUID, role string, err error) {
	jwtToken, err := c.Cookie("token")

	if err != nil || jwtToken == "" {
		return uuid.Nil, "", err
	}

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return uuid.Nil, "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return uuid.Nil, "", err
	}

	userIDStr, ok := claims["userId"].(string)

	if !ok {
		return uuid.Nil, "", err
	}

	userID, err = uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, "", err
	}

	role, ok = claims["role"].(string)

	if !ok {
		return uuid.Nil, "", err
	}

	return userID, role, nil
}

// AuthMiddleware is your authentication middleware.
func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, role, err := getDetailsFromJWTToken(c, secretKey)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized access"})
			c.Abort()
			return
		}

		c.Set("userId", userID)
		c.Set("role", role)
		c.Next()
	}
}

// AdminMiddleware is your admin authentication middleware.
func AdminMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, role, err := getDetailsFromJWTToken(c, secretKey)

		if err != nil || role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized access"})
			c.Abort()
			return
		}

		c.Set("userId", userID)
		c.Set("role", role)
		c.Next()
	}
}
