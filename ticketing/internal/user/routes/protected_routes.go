package routes

import (
	"ticketing/ent"
	"ticketing/internal/common/cache"
	"ticketing/internal/common/repository"
	"ticketing/internal/user/handlers"

	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(routerGroup *gin.RouterGroup, client *ent.Client, redisClient *cache.RedisCache) {
	sessionRepo := repository.NewSessionRepository(client)
	sectionRepo := repository.NewSectionRepository(client, redisClient)
	ticketRepo := repository.NewTicketRepository(client, redisClient)

	sessionHandler := handlers.NewSessionHandler(sessionRepo, sectionRepo, ticketRepo)
	ticketHandler := handlers.NewTicketHandler(ticketRepo)

	sessionGroup := routerGroup.Group("/sessions")
	{
		// Check available seats
		sessionGroup.POST("/available", sessionHandler.GetAndReserveTickets())
	}

	ticketGroup := routerGroup.Group("/tickets")
	{
		// Get ticket based on user email
		ticketGroup.GET("/user", ticketHandler.GetTicketsByUserIDHandler())
	}

}
