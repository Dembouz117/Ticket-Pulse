package routes

import (
	"ticketing/internal/admin/handlers"

	"github.com/gin-gonic/gin"
)

func ConcertRoutes(routerGroup *gin.RouterGroup, concertHandler handlers.ConcertHandler) {
	concertGroup := routerGroup.Group("/concerts")
	{
		concertGroup.GET(".", concertHandler.GetAllConcertsHandler())

		concertGroup.GET("/:id", concertHandler.GetConcertByIDHandler())

		concertGroup.GET("/artist/:artistName", concertHandler.GetConcertsByArtistHandler())

		concertGroup.GET("/:id/sessions", concertHandler.GetSessionsOfConcertHandler())

		concertGroup.POST(".", concertHandler.CreateConcertHandler())

		concertGroup.PUT("/:id", concertHandler.UpdateConcertHandler())

		concertGroup.DELETE("/:id", concertHandler.DeleteConcertHandler())
	}
}
