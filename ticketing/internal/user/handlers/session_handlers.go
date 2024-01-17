package handlers

import (
	"log"
	"net/http"
	"ticketing/internal/common/repository"
	seatallocation "ticketing/internal/common/seatAllocation"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ConcertHandler handles concert-related operations.
type SessionHandler struct {
	Repository        repository.SessionRepository
	SectionRepository repository.SectionRepository
	TicketRepository  repository.TicketRepository
}

// NewSessionHandler creates a new instance of SessionHandler.
func NewSessionHandler(
	repo repository.SessionRepository,
	sectionRepo repository.SectionRepository,
	ticketRepo repository.TicketRepository) *SessionHandler {
	return &SessionHandler{Repository: repo, SectionRepository: sectionRepo, TicketRepository: ticketRepo}
}

type SessionResponse struct {
	ID              uuid.UUID `json:"id"`
	SessionDateTime int       `json:"sessionDateTime"`
}

// Get sections of a concert session
func (h *SessionHandler) GetSectionsOfAConcertSessionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		sectionID := c.Param("id") // Parse the section ID as a UUID
		sectionUUID, err := uuid.Parse(sectionID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section ID"})
			return
		}

		sections, err := h.Repository.GetSectionsOfAConcertSession(c, sectionUUID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sections"})
			return
		}

		c.JSON(http.StatusOK, sections)
	}
}

// Based on the session and payload that contains the section + quantity
func (h *SessionHandler) GetAndReserveTickets() gin.HandlerFunc {
	var requestBody struct {
		SeatRequirements []seatallocation.SeatRequirement `json:"seats"`
	}

	return func(c *gin.Context) {
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

		var availableRequiredTickets []seatallocation.SeatAvailable

		// There should be a seat allocation algorithm
		for _, seat := range requestBody.SeatRequirements {
			sectionID := seat.SectionID
			quantity := seat.Quantity

			err := h.SectionRepository.ReleaseExpiredTicketsBySectionID(c, sectionID)
			if err != nil {
				log.Printf("Cannot release tickets")
			}

			// Check if there are seats
			tickets, err := h.SectionRepository.GetAvailableTicketsBySectionID(c, sectionID)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Invalid ID"})
				return
			}
			numberOfTickets := len(tickets)
			log.Printf("%d", numberOfTickets)

			if numberOfTickets < quantity {
				c.JSON(http.StatusNotFound, gin.H{"error": "Desired tickets are not available"})
				return
			}

			var seatsAvailable []seatallocation.Seat

			for _, ticket := range tickets {
				section := ticket.Edges.WithinSection
				seatsAvailable = append(seatsAvailable, seatallocation.Seat{
					SeatID:          ticket.ID,
					SeatNumber:      ticket.SeatNumber,
					SectionCategory: section.Category.String(),
					SectionName:     section.Name,
					SectionID:       sectionID,
				})
			}

			sectionTickets := seatallocation.SeatAvailable{
				SectionID: sectionID,
				Seats:     seatsAvailable,
			}

			availableRequiredTickets = append(availableRequiredTickets, sectionTickets)
		}

		allocatedSeats := seatallocation.AllocateSeatsConcurrently(requestBody.SeatRequirements, availableRequiredTickets)
		if len(allocatedSeats) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Desired tickets are not available"})
			return
		}

		// Reserve
		for _, ticketsToReserve := range allocatedSeats {
			err := h.TicketRepository.ReserveTicket(c, ticketsToReserve.SectionID, ticketsToReserve.SeatID, userIDUUID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Ticket cannot be reserve"})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"tickets": allocatedSeats})
	}
}

func (h *SessionHandler) GetAllSession() gin.HandlerFunc {
	return func(c *gin.Context) {

		sessions, err := h.Repository.GetAllSessions(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sessions"})
			return
		}
		// Map the sessions to the limited response struct
		response := make([]SessionResponse, len(sessions))
		for i, session := range sessions {
			response[i] = SessionResponse{
				ID:              session.ID,
				SessionDateTime: session.SessionDateTime,
			}
		}

		c.JSON(http.StatusOK, response)
	}
}
