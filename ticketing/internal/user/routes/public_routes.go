package routes

import (
	"ticketing/ent"
	"ticketing/internal/common/cache"
	"ticketing/internal/common/repository"
	"ticketing/internal/user/handlers"

	"github.com/gin-gonic/gin"
)

func SetupPublicRoutes(routerGroup *gin.RouterGroup, client *ent.Client, redisClient *cache.RedisCache) {
	concertRepo := repository.NewConcertRepository(client, redisClient)
	sessionRepo := repository.NewSessionRepository(client)
	sectionRepo := repository.NewSectionRepository(client, redisClient)
	ticketRepo := repository.NewTicketRepository(client, redisClient)

	concertHandler := handlers.NewConcertHandler(concertRepo)
	sessionHandler := handlers.NewSessionHandler(sessionRepo, sectionRepo, ticketRepo)
	sectionHandler := handlers.NewSectionHandler(sectionRepo)
	ticketHandler := handlers.NewTicketHandler(ticketRepo)

	concertGroup := routerGroup.Group("/concerts")
	{
		concertGroup.GET(".", concertHandler.GetAllConcertsHandler())

		concertGroup.GET("/:id", concertHandler.GetConcertByIDHandler())

		concertGroup.GET("/artist/:artistName", concertHandler.GetConcertsByArtistHandler())

		concertGroup.GET("/:id/sessions", concertHandler.GetSessionsOfConcertHandler())

		concertGroup.GET("/featured", concertHandler.GetFeaturedConcerts())
	}

	sessionGroup := routerGroup.Group("/sessions")
	{
		sessionGroup.GET("/:id/sections", sessionHandler.GetSectionsOfAConcertSessionHandler())
	}

	sectionGroup := routerGroup.Group("/sections")
	{
		sectionGroup.GET("/:id/available", sectionHandler.GetAvailableSeatsHandler())

		// Retrieve a specific section by ID.
		sectionGroup.GET("/:id", sectionHandler.GetSectionByID())

		// List all tickets for a specific section.
		sectionGroup.GET("/:id/tickets", sectionHandler.GetAllTicketsBySection())
	}

	ticketGroup := routerGroup.Group("/tickets")
	{
		// Retrieve a specific ticket by ID.
		ticketGroup.GET("/:id", ticketHandler.GetTicketByIdHandler())
	}

}
