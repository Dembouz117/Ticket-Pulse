package middleware

import (
	"context"
	"net/http"
	"ticketing/internal/common/database"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// DatabaseConnectionMiddleware is a middleware function that establishes a database connection
// and adds the client to the request context.
func DatabaseConnectionMiddleware(postgresAddress string, ctx context.Context) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Initialize the database connection
        client, err := database.ConnectDatabase(postgresAddress, ctx)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
            c.Abort()
            return
        }
        defer client.Close()

        // ctx := context.WithValue(c.Request.Context(), "db", client)
        // c.Request = c.Request.WithContext(ctx)
        c.Set("db", client)
        c.Next()
    }
}
