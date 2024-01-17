package routes

import (
	"ticketing/internal/admin/handlers"

	"github.com/gin-gonic/gin"
)

func SessionRoutes(routerGroup *gin.RouterGroup, sessionHandler handlers.SessionHandler) {
	sessionGroup := routerGroup.Group("/sessions")
	{
		// GET ALL
		sessionGroup.GET(".", sessionHandler.GetAllConcertSessionHandler())
		sessionGroup.GET("/:id", sessionHandler.GetSessionByIdHandler())

		// POST
		sessionGroup.POST(".", sessionHandler.CreateConcertSessionHandler())

		// UPDATE
		sessionGroup.PUT("/:id", sessionHandler.UpdateConcertSessionHandler())

		// DELETE
		sessionGroup.DELETE("/:id", sessionHandler.DeleteConcertSessionHandler())

		// List all sections for a specific session.
		sessionGroup.GET("/:id/sections", sessionHandler.GetSectionsBySessionHandler())
	}
}
