package handlers

import (
	"net/http"
	"ticketing/internal/common/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ConcertHandler handles concert-related operations.
type SectionHandler struct {
	Repository repository.SectionRepository
}

// NewSectionHandler creates a new instance of SectionHandler.
func NewSectionHandler(repo repository.SectionRepository) *SectionHandler {
	return &SectionHandler{Repository: repo}
}

func (h *SectionHandler) GetAvailableSeatsHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		sectionID := c.Param("id")

		// Parse the section ID as a UUID
		sectionUUID, err := uuid.Parse(sectionID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section ID"})
			return
		}

		tickets, err := h.Repository.GetAvailableTicketsBySectionID(c, sectionUUID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch concerts"})
			return
		}

		c.JSON(http.StatusOK, tickets)
	}
}

func (h *SectionHandler) GetSectionByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		sectionID := c.Param("id") // Assuming the section ID is passed as a route parameter

		// Parse the section ID as a UUID
		sectionUUID, err := uuid.Parse(sectionID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section ID"})
			return
		}

		section, err := h.Repository.GetSectionByID(c, sectionUUID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch section"})
			return
		}

		// Return the section as a JSON response
		c.JSON(http.StatusOK, section)
	}
}

func (h *SectionHandler) GetAllTicketsBySection() gin.HandlerFunc {
	return func(c *gin.Context) {
		sectionIDStr := c.Param("id")
		// Parse the section ID as a UUID
		sectionID, err := uuid.Parse(sectionIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section ID"})
			return
		}

		tickets, err := h.Repository.GetTicketsBySection(c, sectionID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tickets"})
			return
		}

		c.JSON(http.StatusOK, tickets)
	}
}
