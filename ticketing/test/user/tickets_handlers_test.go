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

func setupRouter(t *testing.T) (*gin.Engine, *mocks.MockTicketRepository) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockTicketRepository(ctrl)
	ticketHandler := handlers.NewTicketHandler(mockRepo)

	router := gin.Default()

	router.GET("/tickets", ticketHandler.GetAllTicketsHandler())
	router.GET("/tickets/:id", ticketHandler.GetTicketByIdHandler())
	router.GET("/user/tickets", ticketHandler.GetTicketsByUserIDHandler())
	router.POST("/tickets/reserve", ticketHandler.ReserveTickets())

	return router, mockRepo
}

func TestGetAllTicketsHandler(t *testing.T) {
	router, mockRepo := setupRouter(t)

	expectedTickets := []*ent.Ticket{{ /* ... fields ... */ }, { /* ... fields ... */ }}
	mockRepo.EXPECT().GetAllTickets(gomock.Any()).Return(expectedTickets, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tickets", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// You would also assert that the body of the response matches the expectedTickets
}

func TestGetTicketByIdHandler(t *testing.T) {
	router, mockRepo := setupRouter(t)

	ticketID := uuid.New()
	expectedTicket := &ent.Ticket{ /* ... fields ... */ }
	mockRepo.EXPECT().GetTicketById(gomock.Any(), ticketID).Return(expectedTicket, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tickets/"+ticketID.String(), nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
