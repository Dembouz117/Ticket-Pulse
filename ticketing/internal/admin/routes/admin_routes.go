package routes

import (
	"ticketing/ent"
	"ticketing/internal/admin/handlers"
	"ticketing/internal/common/cache"
	"ticketing/internal/common/repository"

	"github.com/gin-gonic/gin"
)

// SetupAdminRoutes sets up admin-specific routes.
func SetupAdminRoutes(routerGroup *gin.RouterGroup, client *ent.Client, redisClient *cache.RedisCache) {
	concertRoutes(routerGroup, client, redisClient)
	sectionRoutes(routerGroup, client, redisClient)
	sessionRoutes(routerGroup, client)
	ticketRoutes(routerGroup, client, redisClient)
}

// concertRoutes sets up routes for managing concerts.
func concertRoutes(routerGroup *gin.RouterGroup, client *ent.Client, redisClient *cache.RedisCache) {
	// Define your concert routes here
	concertHandler := handlers.NewConcertHandler(repository.NewConcertRepository(client, redisClient))
	ConcertRoutes(routerGroup, *concertHandler)
}

// sessionRoutes sets up routes for managing sessions.
func sessionRoutes(routerGroup *gin.RouterGroup, client *ent.Client) {
	// Define your session routes here
	sessionHandler := handlers.NewSessionHandler(repository.NewSessionRepository(client))
	SessionRoutes(routerGroup, *sessionHandler)
}

// sectionRoutes sets up routes for managing sections.
func sectionRoutes(routerGroup *gin.RouterGroup, client *ent.Client, redisClient *cache.RedisCache) {
	// Define your section routes here
	sectionHandler := handlers.NewSectionHandler(
		repository.NewSectionRepository(client, redisClient),
		repository.NewTicketRepository(client, redisClient),
	)
	SectionRoutes(routerGroup, *sectionHandler)
}

// ticketRoutes sets up routes for managing tickets.
func ticketRoutes(routerGroup *gin.RouterGroup, client *ent.Client, redisClient *cache.RedisCache) {
	ticketHandler := handlers.NewTicketHandler(repository.NewTicketRepository(client, redisClient))
	TicketRoutes(routerGroup, *ticketHandler)
}
