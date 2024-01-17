package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"ticketing/ent"
	"ticketing/internal/user/handlers"
	mocks "ticketing/mocks/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// TestSessionHandler_GetSectionsOfAConcertSessionHandler tests the GetSectionsOfAConcertSessionHandler method.
func TestSessionHandler_GetSectionsOfAConcertSessionHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSessionRepo := mocks.NewMockSessionRepository(ctrl)
	mockSectionRepo := mocks.NewMockSectionRepository(ctrl)
	mockTicketRepo := mocks.NewMockTicketRepository(ctrl)

	sessionHandler := handlers.NewSessionHandler(mockSessionRepo, mockSectionRepo, mockTicketRepo)

	// Setting up the expected session UUID
	sessionUUID := uuid.New()

	// Mocking the repository call
	mockSessionRepo.EXPECT().
		GetSectionsOfAConcertSession(gomock.Any(), sessionUUID).
		Return([]*ent.Section{{ID: sessionUUID}}, nil).
		Times(1)

	// Setting up the HTTP request and recorder
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/concert-session/"+sessionUUID.String()+"/sections", nil)
	_, r := gin.CreateTestContext(w)

	// Adding the handler to the Gin router and running the request
	r.GET("/concert-session/:id/sections", sessionHandler.GetSectionsOfAConcertSessionHandler())
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var sectionsResponse []*ent.Section
	err := json.Unmarshal(w.Body.Bytes(), &sectionsResponse)
	assert.NoError(t, err)
	assert.NotEmpty(t, sectionsResponse)
	assert.Equal(t, sessionUUID, sectionsResponse[0].ID)
}

// TestSessionHandler_GetAllSession tests the GetAllSession method.
func TestSessionHandler_GetAllSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSessionRepo := mocks.NewMockSessionRepository(ctrl)
	sessionHandler := handlers.NewSessionHandler(mockSessionRepo, nil, nil)

	// Mock a response slice
	mockSessions := []*ent.ConcertSession{
		{ID: uuid.New(), SessionDateTime: 123456789},
		{ID: uuid.New(), SessionDateTime: 987654321},
	}

	// Mock the expected interaction with the repository
	mockSessionRepo.EXPECT().
		GetAllSessions(gomock.Any()).
		Return(mockSessions, nil).
		Times(1)

	// Setup the HTTP request and recorder
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/sessions", nil)
	_, r := gin.CreateTestContext(w)

	// Adding the handler to the Gin router and running the request
	r.GET("/sessions", sessionHandler.GetAllSession())
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var sessionsResponse []handlers.SessionResponse
	err := json.Unmarshal(w.Body.Bytes(), &sessionsResponse)
	assert.NoError(t, err)
	assert.Len(t, sessionsResponse, len(mockSessions))
	for i, sessionResponse := range sessionsResponse {
		assert.Equal(t, mockSessions[i].ID, sessionResponse.ID)
		assert.Equal(t, mockSessions[i].SessionDateTime, sessionResponse.SessionDateTime)
	}
}
