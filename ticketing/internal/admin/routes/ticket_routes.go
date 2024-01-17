package routes

import (
	adminHandler "ticketing/internal/admin/handlers"

	"github.com/gin-gonic/gin"
)

func TicketRoutes(routerGroup *gin.RouterGroup, ticketHandler adminHandler.TicketHandler) {
	ticketGroup := routerGroup.Group("/tickets")
	{
		ticketGroup.GET(".", ticketHandler.GetAllTicketsHandler())

		// Retrieve a specific ticket by ID.
		ticketGroup.GET("/:id", ticketHandler.GetTicketByIdHandler())

		// Update a ticket by ID.
		ticketGroup.PUT("/:id", ticketHandler.UpdateTicketHandler())
	}
}
