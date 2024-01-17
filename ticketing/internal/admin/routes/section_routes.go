package routes

import (
	"ticketing/internal/admin/handlers"

	"github.com/gin-gonic/gin"
)

func SectionRoutes(routerGroup *gin.RouterGroup, sectionHandler handlers.SectionHandler) {
	sectionGroup := routerGroup.Group("/sections")
	{
		// Retrieve a specific section by ID.
		sectionGroup.GET("/:id", sectionHandler.GetSectionByID())

		sectionGroup.POST(".", sectionHandler.CreateSectionWithTickets())

		// Update a section by ID.
		sectionGroup.PUT("/:id", sectionHandler.UpdateSection())

		// Delete a section by ID.
		sectionGroup.DELETE("/:id", sectionHandler.DeleteSection())

		// List all tickets for a specific section.
		sectionGroup.GET("/:id/tickets", sectionHandler.GetTicketsBySection())
	}
}
