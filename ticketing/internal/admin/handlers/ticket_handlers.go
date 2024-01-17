// @tags admin
// Admin tickets

package handlers

import (
	"fmt"
	"net/http"
	"ticketing/ent/ticket"
	"ticketing/internal/common/mapping"
	"ticketing/internal/common/repository"

	_ "ticketing/internal/common/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TicketHandler handles ticket-related operations.
type TicketHandler struct {
	Repository repository.TicketRepository
}

// NewTicketHandler creates a new instance of TicketHandler.
func NewTicketHandler(repo repository.TicketRepository) *TicketHandler {
	return &TicketHandler{Repository: repo}
}

// @Summary Get all Tickets
// @Description Retrieves a list of all tickets.
// @Accept json
// @Produce json
// @Tags admin
// @Success 200 {object} ent.Ticket
// @Failure 500 {object} types.ErrorResponse
// @Router /admin/tickets [get]
func (h *TicketHandler) GetAllTicketsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tickets, err := h.Repository.GetAllTickets(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tickets"})
			return
		}

		mappedTickets, err := mapping.FromEntTicketList(tickets)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tickets"})
			return
		}

		c.JSON(http.StatusOK, mappedTickets)
	}
}

// @Summary Get Ticket by ID
// @Description Retrieves a ticket by its ID.
// @Accept json
// @Produce json
// @Tags admin
// @Param id path string true "Ticket ID"
// @Success 200 {object} types.Ticket
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /admin/tickets/{id} [get]
func (h *TicketHandler) GetTicketByIdHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		ticketIDStr := c.Param("id") // Assuming the section ID is passed as a route parameter
		ticketID, err := uuid.Parse(ticketIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
			return
		}

		tickets, err := h.Repository.GetTicketById(c, ticketID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tickets"})
			return
		}

		c.JSON(http.StatusOK, tickets)
	}
}

// @Summary Update Ticket
// @Description Updates a ticket by its ID.
// @Accept json
// @Produce json
// @Tags admin
// @Param id path string true "Ticket ID"
// @Param seatNumber body int true "Seat Number"
// @Param status body string true "Status"
// @Param userId body string true "User ID"
// @Success 200 {object} types.Ticket
// @Failure 400 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Router /admin/tickets/{id} [put]
func (h *TicketHandler) UpdateTicketHandler() gin.HandlerFunc {
	var requestBody struct {
		SeatNumber int    `json:"seatNumber"`
		Status     string `json:"status"`
		UserID     string `json:"userId"`
	}

	return func(c *gin.Context) {
		ticketIDStr := c.Param("id")

		// Parse the section ID as a UUID
		ticketID, err := uuid.Parse(ticketIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
			return
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON request"})
			return
		}

		userId, err := uuid.Parse(requestBody.UserID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid User ID"})
			return
		}

		// Ensure that the provided category value is one of the allowed enum values
		allowedCategories := map[string]ticket.Status{
			"AVAILABLE": ticket.StatusAVAILABLE,
			"BOUGHT":    ticket.StatusBOUGHT,
			"RESERVED":  ticket.StatusRESERVED,
		}

		status, ok := allowedCategories[requestBody.Status]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid status value"})
			return
		}

		ticket, err := h.Repository.UpdateTicketStatus(c, ticketID, status, userId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
			return
		}

		// Return the updated section as a JSON response
		c.JSON(http.StatusOK, gin.H{"ticket": ticket})
	}
}
