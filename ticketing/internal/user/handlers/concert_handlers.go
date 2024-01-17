// @tags user
// User concerts
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
// @Tags user
// @Success 200 {object} types.ConcertListResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /concerts [get]
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
// @Tags user
// @Param id path string true "Concert ID"
// @Success 200 {object} types.ConcertResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Router /concerts/{id} [get]
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

		c.JSON(http.StatusOK, gin.H{"concert": concert})
	}
}

// @Summary Get Concerts by Artist
// @Description Retrieves concerts by artist.
// @Accept json
// @Produce json
// @Tags user
// @Param artistName path string true "Artist Name"
// @Success 200 {object} types.ConcertListResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Router /concerts/artist/{artistName} [get]
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
			c.JSON(http.StatusNotFound, gin.H{"message": "No concerts found for the artist"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"concerts": concerts})
	}
}

// @Summary Get Sessions of a Concert
// @Description Retrieves sessions of a concert by its ID.
// @Accept json
// @Produce json
// @Tags user
// @Param id path string true "Concert ID"
// @Success 200 {object} types.ConcertWithSessionListResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /concerts/{id}/sessions [get]
func (h *ConcertHandler) GetSessionsOfConcertHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the concert ID from the URL parameters
		concertIDStr := c.Param("id")
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

		// Query the sessions associated with the concert using the repository
		sessions, err := h.Repository.GetSessionsOfConcert(c, concert.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve ConcertSessions"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"concert": concert, "sessions": sessions})
	}
}

// @Summary Get Featured Concerts
// @Description Retrieves featured concerts
// @Accept json
// @Produce json
// @Tags user
// @Success 200 {object} types.ConcertListResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /concerts/featured [get]
func (h *ConcertHandler) GetFeaturedConcerts() gin.HandlerFunc {
	return func(c *gin.Context) {
		concerts, err := h.Repository.GetFeaturedConcerts(c)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "No concerts found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"concerts": concerts})
	}
}
