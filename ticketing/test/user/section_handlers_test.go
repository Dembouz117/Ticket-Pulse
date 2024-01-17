package handlers_test

import (
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

func TestGetAvailableSeatsHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSectionRepository(ctrl)
	sectionHandler := handlers.NewSectionHandler(mockRepo)

	sectionID := uuid.New()
	mockRepo.EXPECT().GetAvailableTicketsBySectionID(gomock.Any(), sectionID).Return([]*ent.Ticket{}, nil)

	router := gin.Default()
	router.GET("/sections/:id/available-seats", sectionHandler.GetAvailableSeatsHandler())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/sections/"+sectionID.String()+"/available-seats", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetSectionByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSectionRepository(ctrl)
	sectionHandler := handlers.NewSectionHandler(mockRepo)

	sectionID := uuid.New()
	mockRepo.EXPECT().GetSectionByID(gomock.Any(), sectionID).Return(&ent.Section{}, nil)

	router := gin.Default()
	router.GET("/sections/:id", sectionHandler.GetSectionByID())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/sections/"+sectionID.String(), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAllTicketsBySection(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSectionRepository(ctrl)
	sectionHandler := handlers.NewSectionHandler(mockRepo)

	sectionID := uuid.New()
	mockRepo.EXPECT().GetTicketsBySection(gomock.Any(), sectionID).Return([]*ent.Ticket{}, nil)

	router := gin.Default()
	router.GET("/sections/:id/tickets", sectionHandler.GetAllTicketsBySection())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/sections/"+sectionID.String()+"/tickets", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
