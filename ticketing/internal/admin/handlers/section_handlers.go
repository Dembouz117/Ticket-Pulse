// @tags admin
// Admin sections

package handlers

import (
	"log"
	"net/http"
	"ticketing/ent/section"
	"ticketing/ent/ticket"
	"ticketing/internal/common/repository"
	_ "ticketing/internal/common/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ConcertHandler handles concert-related operations.
type SectionHandler struct {
	Repository       repository.SectionRepository
	TicketRepository repository.TicketRepository
}

// NewSectionHandler creates a new instance of SectionHandler.
func NewSectionHandler(repo repository.SectionRepository, ticketRepo repository.TicketRepository) *SectionHandler {
	return &SectionHandler{Repository: repo, TicketRepository: ticketRepo}
}

// @Summary Create a section with tickets
// @Description Create a new section and its associated tickets
// @Accept  json
// @Produce  json
// @Tags admin
// @Param sessionId body string true "Session ID"
// @Param name body string true "Section Name"
// @Param capacity body int true "Section Capacity"
// @Param category body string true "Section Category"
// @Param price body int true "Ticket Price"
// @Success 200 {object} types.SuccessResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /admin/sections [post]
func (h *SectionHandler) CreateSectionWithTickets() gin.HandlerFunc {
	var requestBody struct {
		SessionID string `json:"sessionId"`
		Name      string `json:"name"`
		Capacity  int    `json:"capacity"`
		Category  string `json:"category"`
		Price     int    `json:"price"`
	}

	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			log.Println("Invalid JSON request")
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON request"})
			return
		}

		sessionID, err := uuid.Parse(requestBody.SessionID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
			return
		}

		// Ensure that the provided category value is one of the allowed enum values
		allowedCategories := map[string]section.Category{
			"CAT1": section.CategoryCAT1,
			"CAT2": section.CategoryCAT2,
			"CAT3": section.CategoryCAT3,
			"CAT4": section.CategoryCAT4,
			"CAT5": section.CategoryCAT5,
		}

		category, ok := allowedCategories[requestBody.Category]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid category value"})
			return
		}

		// Create a new Section entity
		section, err := h.Repository.
			CreateSection(c, sessionID,
				requestBody.Name,
				requestBody.Capacity, 0, 0,
				category,
				requestBody.Price)

		if err != nil {
			log.Printf("Failed to create Section: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Section"})
			return
		}

		// Creates the number of tickets based on capacity
		for i := 0; i < requestBody.Capacity; i++ {
			_, err := h.TicketRepository.CreateTicket(c, section.ID, i+1, ticket.StatusAVAILABLE)
			if err != nil {
				log.Printf("Failed to create Ticket for Section: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Ticket"})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "Section and Tickets created successfully", "id": section.ID})
	}
}

// @Summary Get details of a section
// @Description Retrieve details of a specific section by its ID
// @Accept  json
// @Produce  json
// @Tags admin
// @Param id path string true "Section ID"
// @Success 200 {object} types.SectionResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /admin/sections/{id} [get]
func (h *SectionHandler) GetSectionByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		sectionID := c.Param("id") // Assuming the section ID is passed as a route parameter

		// Parse the section ID as a UUID
		sectionUUID, err := uuid.Parse(sectionID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section ID"})
			return
		}

		// Retrieve the section by its ID
		section, err := h.Repository.GetSectionByID(c, sectionUUID)

		if err != nil {
			log.Printf("Failed to fetch section: %v", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Section not found"})
			return
		}

		// Return the section as a JSON response
		c.JSON(http.StatusOK, section)
	}
}

// @Summary Update a section
// @Description Update details of a specific section by its ID
// @Accept  json
// @Produce  json
// @Tags admin
// @Param id path string true "Section ID"
// @Param name body string true "Section Name"
// @Param capacity body int true "Section Capacity"
// @Param reserved body int true "Reserved Seats"
// @Param bought body int true "Bought Seats"
// @Param category body string true "Section Category"
// @Param price body int true "Ticket Price"
// @Success 200 {object} types.SectionResponse
// @Failure 400 {object} types.SuccessResponse
// @Router /admin/sections/{id} [put]
func (h *SectionHandler) UpdateSection() gin.HandlerFunc {
	// Parse the request body
	var requestBody struct {
		Name     string `json:"name"`
		Capacity int    `json:"capacity"`
		Reserved int    `json:"reserved"`
		Bought   int    `json:"bought"`
		Category string `json:"category"`
		Price    int    `json:"price"`
	}

	return func(c *gin.Context) {
		sectionID := c.Param("id") // Assuming the section ID is passed as a route parameter

		// Parse the section ID as a UUID
		sectionUUID, err := uuid.Parse(sectionID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section ID"})
			return
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			log.Println("Invalid JSON request")
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON request"})
			return
		}

		// Ensure that the provided category value is one of the allowed enum values
		allowedCategories := map[string]section.Category{
			"CAT1": section.CategoryCAT1,
			"CAT2": section.CategoryCAT2,
			"CAT3": section.CategoryCAT3,
			"CAT4": section.CategoryCAT4,
			"CAT5": section.CategoryCAT5,
		}

		category, ok := allowedCategories[requestBody.Category]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid category value"})
			return
		}

		// Update the section entity
		section, err := h.Repository.UpdateSection(c, sectionUUID, requestBody.Name,
			requestBody.Capacity, requestBody.Reserved, requestBody.Bought,
			category, requestBody.Price)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update section"})
			return
		}

		// Return the updated section as a JSON response
		c.JSON(http.StatusOK, gin.H{"section": section})
	}
}

// @Summary Get tickets by section
// @Description Retrieve all tickets associated with a specific section by its ID
// @Accept  json
// @Produce  json
// @Tags admin
// @Param id path string true "Section ID"
// @Success 200 {object} types.SectionWithTicketsListResponse
// @Failure 400 {object} types.SuccessResponse
// @Router /admin/sections/{id}/tickets [get]
func (h *SectionHandler) GetTicketsBySection() gin.HandlerFunc {
	return func(c *gin.Context) {
		sectionIDStr := c.Param("id") // Assuming the section ID is passed as a route parameter
		// Parse the section ID as a UUID
		sectionID, err := uuid.Parse(sectionIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section ID"})
			return
		}

		// Delete the section by its ID
		sectionWithTickets, err := h.Repository.GetTicketsBySection(c, sectionID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete section"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"tickets": sectionWithTickets})
	}
}

// @Summary Delete a section
// @Description Delete a specific section by its ID
// @Accept  json
// @Produce  json
// @Tags admin
// @Param id path string true "Section ID"
// @Success 200 {object} types.SuccessResponse
// @Failure 400 {object} types.SuccessResponse
// @Router /admin/sections/{id} [delete]
func (h *SectionHandler) DeleteSection() gin.HandlerFunc {
	return func(c *gin.Context) {
		sectionID := c.Param("id") // Assuming the section ID is passed as a route parameter

		// Parse the section ID as a UUID
		sectionUUID, err := uuid.Parse(sectionID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section ID"})
			return
		}
		err = h.Repository.DeleteSection(c, sectionUUID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Section deleted successfully"})

	}
}
