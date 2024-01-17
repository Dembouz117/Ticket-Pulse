// @tags admin
// Admin sessions

package handlers

import (
	"net/http"
	"ticketing/internal/common/repository"
	_ "ticketing/internal/common/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ConcertHandler handles concert-related operations.
type SessionHandler struct {
	Repository repository.SessionRepository
}

// NewSessionHandler creates a new instance of SessionHandler.
func NewSessionHandler(repo repository.SessionRepository) *SessionHandler {
	return &SessionHandler{Repository: repo}
}

// @Summary Get all Concert Sessions
// @Description Retrieves a list of all concert sessions.
// @Accept json
// @Produce json
// @Tags admin
// @Success 200 {object} types.SessionListResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /admin/sessions	 [get]
func (h *SessionHandler) GetAllConcertSessionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Query all concerts sessions
		sessions, err := h.Repository.GetAllSessions(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sessions"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"sessions": sessions})
	}
}

// @Summary Create a new Concert Session
// @Description Creates a new concert session with the provided details.
// @Accept json
// @Produce json
// @Tags admin
// @Param sessionDateTime body int true "Date and time of the session"
// @Param concertId body string true "ID of the concert"
// @Success 200 {object} types.SessionResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /admin/sessions [post]
func (h *SessionHandler) CreateConcertSessionHandler() gin.HandlerFunc {
	var requestBody struct {
		SessionDateTime int    `json:"sessionDateTime"`
		ConcertID       string `json:"concertId"`
	}
	return func(c *gin.Context) {

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON request"})
			return
		}

		// Parse the Concert ID as a UUID
		concertID, err := uuid.Parse(requestBody.ConcertID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid concert ID"})
			return
		}

		session, err := h.Repository.CreateConcertSession(c, concertID, requestBody.SessionDateTime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch concerts"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "ConcertSession created successfully", "id": session.ID})
	}
}

// @Summary Get a Concert Session by ID
// @Description Retrieves a concert session by its ID.
// @Accept json
// @Produce json
// @Tags admin
// @Param id path string true "Session ID"
// @Success 200 {object} types.SessionResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Router /admin/sessions/{id} [get]
func (h *SessionHandler) GetSessionByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionIDStr := c.Param("id")
		sessionID, err := uuid.Parse(sessionIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
			return
		}

		session, err := h.Repository.GetSessionByID(c, sessionID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
			return
		}

		c.JSON(http.StatusOK, session)
	}
}

// @Summary Update a Concert Session
// @Description Updates the date and time of a concert session.
// @Accept json
// @Produce json
// @Tags admin
// @Param id path string true "Session ID"
// @Param sessionDateTime body int true "Date and time of the session"
// @Success 200 {object} types.SessionResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /admin/sessions/{id} [put]
func (h *SessionHandler) UpdateConcertSessionHandler() gin.HandlerFunc {
	var requestBody struct {
		SessionDateTime int `json:"sessionDateTime"`
	}

	return func(c *gin.Context) {
		sessionIDStr := c.Param("id")
		sessionID, err := uuid.Parse(sessionIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
			return
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON request"})
			return
		}

		session, err := h.Repository.UpdateConcertSession(c, sessionID, requestBody.SessionDateTime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update ConcertSession"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"session": session})
	}
}

// @Summary Delete a Concert Session
// @Description Deletes a concert session and its children by its ID.
// @Accept json
// @Produce json
// @Tags admin
// @Param id path string true "Session ID"
// @Success 200 {object} types.SuccessResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /admin/sessions/{id} [delete]
func (h *SessionHandler) DeleteConcertSessionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionIDStr := c.Param("id")

		// Parse the Session ID as a UUID
		sessionID, err := uuid.Parse(sessionIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
			return
		}

		// Query the ConcertSession by its ID
		err = h.Repository.DeleteConcertSession(c, sessionID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ConcertSession not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "ConcertSession deleted successfully"})
	}
}

// @Summary Get Sections by Session ID
// @Description Retrieves sections associated with a concert session by its ID.
// @Accept json
// @Produce json
// @Tags admin
// @Param id path string true "Session ID"
// @Success 200 {object} types.SessionWithSectionListResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Router /admin/sessions/{id}/sections [get]
func (h *SessionHandler) GetSectionsBySessionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the session ID from the URL parameters
		sessionIDStr := c.Param("id")
		sessionID, err := uuid.Parse(sessionIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
			return
		}

		sections, err := h.Repository.GetSectionsOfAConcertSession(c, sessionID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Concert Session not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"sections": sections})
	}
}
