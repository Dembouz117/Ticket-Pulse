package handlers

import (
	"fmt"
	"net/http"
	"ticketing/ent/ticket"
	"ticketing/internal/common/repository"

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

// GetAllTicketsHandler handles the route for retrieving all tickets.
// Not needed
func (h *TicketHandler) GetAllTicketsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tickets, err := h.Repository.GetAllTickets(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tickets"})
			return
		}

		c.JSON(http.StatusOK, tickets)
	}
}

func (h *TicketHandler) GetTicketByIdHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		ticketIDStr := c.Param("id")
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

func (h *TicketHandler) GetTicketsByUserIDHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userId")

		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "User ID not found in context"})
			return
		}

		userIDUUID, ok := userID.(uuid.UUID)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid User ID in context"})
			return
		}

		tickets, err := h.Repository.GetTicketsByUserID(c, userIDUUID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tickets"})
			return
		}

		c.JSON(http.StatusOK, tickets)
	}
}

func (h *TicketHandler) ReserveTickets() gin.HandlerFunc {
	// takes in an array of ticket
	var requestBody struct {
		Tickets []string `json:"tickets"`
	}

	return func(c *gin.Context) {
		// verify the json is legit
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON request"})
			return
		}
		userID, exists := c.Get("userId")

		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "User ID not found in context"})
			return
		}

		userIDUUID, ok := userID.(uuid.UUID)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid User ID in context"})
			return
		}

		// for loop and update to reserve
		for _, ticketToReserve := range requestBody.Tickets {

			// check if the ticket ID is legit
			ticketID, err := uuid.Parse(ticketToReserve)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid ticket ID: %s", ticketToReserve)})
				return
			}

			foundTicket, err := h.Repository.GetTicketById(c, ticketID)
			if err != nil {
				// Handle the error (e.g., return an error response)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve ticket"})
				return
			}

			// check if the tickets are available
			if foundTicket.Status != ticket.StatusAVAILABLE {
				c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("This ticket is not available: %s", ticketToReserve)})
				return
			}

			// Use the repository method to update the ticket status
			_, err = h.Repository.UpdateTicketStatus(c, ticketID, ticket.StatusRESERVED, userIDUUID)
			if err != nil {
				// Handle the error (e.g., return an error response)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reserve ticket"})
				return
			}
		}

		// Return the updated section as a JSON response
		c.JSON(http.StatusOK, gin.H{"message": "Tickets have been verified"})
	}
}
