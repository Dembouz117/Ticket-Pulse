// @tags admin
// Admin concerts

package handlers

import (
	"net/http"
	"ticketing/internal/common/repository"
	_ "ticketing/internal/common/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// ConcertHandler handles concert-related operations.
type ConcertHandler struct {
	Repository repository.ConcertRepository
}

// NewConcertHandler creates a new instance of ConcertHandler.
func NewConcertHandler(repo repository.ConcertRepository) *ConcertHandler {
	return &ConcertHandler{Repository: repo}
}

// @Summary Get all Concerts
// @Description Retrieves a list of all concerts.
// @Accept json
// @Produce json
// @Tags admin
// @Success 200 {object} types.ConcertListResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /admin/concerts [get]
func (h *ConcertHandler) GetAllConcertsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		concerts, err := h.Repository.GetAllConcerts(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch concerts"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"concerts": concerts})
	}
}

// @Summary Get a Concert by ID
// @Description Retrieves details of a concert by its ID.
// @Accept json
// @Produce json
// @Tags admin
// @Param id path string true "Concert ID"
// @Success 200 {object} types.ConcertResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Router /admin/concerts/{id} [get]
func (h *ConcertHandler) GetConcertByIDHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		concertIDStr := c.Param("id")
		concertID, err := uuid.Parse(concertIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid concert ID"})
			return
		}

		concert, err := h.Repository.GetConcertByID(c, concertID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Concert not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"concerts": concert})
	}
}

// @Summary Get Concerts by Artist
// @Description Retrieves concerts by artist.
// @Accept json
// @Produce json
// @Tags admin
// @Param artistName path string true "Artist Name"
// @Success 200 {object} types.ConcertListResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Router /admin/concerts/artist/{artistName} [get]
func (h *ConcertHandler) GetConcertsByArtistHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		artistName := c.Param("artistName")
		if artistName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid artist name"})
			return
		}

		concerts, err := h.Repository.GetConcertsByArtist(c, artistName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch concerts"})
			return
		}

		if len(concerts) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No concerts found for the artist"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"concerts": concerts})
	}
}

// @Summary Create a Concert
// @Description Creates a new concert.
// @Accept json
// @Produce json
// @Tags admin
// @Param title body string true "Concert Title"
// @Param artist body string true "Artist Name"
// @Param imageUrl body string true "Image URL"
// @Param description body string true "Description"
// @Param headline body string true "Headline"
// @Param featured body boolean true "Featured"
// @Success 201 {object} types.ConcertResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /admin/concerts [post]
func (h *ConcertHandler) CreateConcertHandler() gin.HandlerFunc {
	var requestBody struct {
		Title       string `json:"title"`
		Artist      string `json:"artist"`
		ImageUrl    string `json:"imageUrl"`
		Description string `json:"description"`
		Headline    string `json:"headline"`
		Featured    bool   `json:"featured"`
	}
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
			return
		}

		newConcert, err := h.Repository.CreateConcert(c, requestBody.Title, requestBody.Artist, requestBody.ImageUrl, requestBody.Description, requestBody.Headline, requestBody.Featured)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating concert"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"concert": newConcert})
	}
}

// @Summary Update a Concert by ID
// @Description Updates a concert by its ID.
// @Accept json
// @Produce json
// @Tags admin
// @Param id path string true "Concert ID"
// @Param title body string true "Concert Title"
// @Param artist body string true "Artist Name"
// @Param imageUrl body string true "Image URL"
// @Param description body string true "Description"
// @Param headline body string true "Headline"
// @Param featured body boolean true "Featured"
// @Success 200 {object} types.ConcertResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /admin/concerts/{id} [put]
func (h *ConcertHandler) UpdateConcertHandler() gin.HandlerFunc {
	var requestBody struct {
		Title       string `json:"title"`
		Artist      string `json:"artist"`
		ImageURL    string `json:"imageUrl"`
		Description string `json:"description"`
		Headline    string `json:"headline"`
		Featured    bool   `json:"featured"`
	}

	return func(c *gin.Context) {
		concertIDStr := c.Param("id")
		concertID, err := uuid.Parse(concertIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid concert ID"})
			return
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
			return
		}

		concert, err := h.Repository.UpdateConcert(c, concertID, requestBody.Title, requestBody.Artist, requestBody.ImageURL, requestBody.Description, requestBody.Headline, requestBody.Featured)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating concert"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"concert": concert})
	}
}

// @Summary Delete a Concert by ID
// @Description Deletes a concert by its ID.
// @Accept json
// @Produce json
// @Tags admin
// @Param id path string true "Concert ID"
// @Success 200 {object} types.SuccessResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /admin/concerts/{id} [delete]
func (h *ConcertHandler) DeleteConcertHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		concertIDStr := c.Param("id")
		concertID, err := uuid.Parse(concertIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid concert ID"})
			return
		}

		err = h.Repository.DeleteConcert(c, concertID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting concert"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Concert deleted successfully"})
	}
}

// @Summary Get Sessions of a Concert
// @Description Retrieves sessions of a concert by its ID.
// @Accept json
// @Produce json
// @Tags admin
// @Param id path string true "Concert ID"
// @Success 200 {object} types.ConcertWithSessionListResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /admin/concerts/{id}/sessions [get]
func (h *ConcertHandler) GetSessionsOfConcertHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the concert ID from the URL parameters
		concertIDStr := c.Param("id")

		// Parse the Concert ID as a UUID
		concertID, err := uuid.Parse(concertIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid concert ID"})
			return
		}

		// Query the Concert by its ID
		concert, err := h.Repository.GetConcertByID(c, concertID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Concert not found"})
			return
		}

		sessions, err := h.Repository.GetSessionsOfConcert(c, concert.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve ConcertSessions"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"concert": concert, "sessions": sessions})
	}
}
