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

func TestGetAllConcertsHandler(t *testing.T) {
	// Initialize the mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock repository
	mockRepo := mocks.NewMockConcertRepository(ctrl)

	// Set up your expected behavior here
	concerts := []*ent.Concert{} // Assume types.Concert is the correct type
	mockRepo.EXPECT().GetAllConcerts(gomock.Any()).Return(concerts, nil).Times(1)

	// Create your handler with the mock repo
	handler := handlers.NewConcertHandler(mockRepo)

	// Set up the Gin engine
	router := gin.Default()
	router.GET("/concerts", handler.GetAllConcertsHandler())

	// Record the HTTP response
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/concerts", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetConcertByIDHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockConcertRepository(ctrl)
	concertHandler := handlers.NewConcertHandler(mockRepo)

	// Use the exact UUID from the error message in your test
	testUUID, _ := uuid.Parse("b26f5c81-5ec8-4d64-98b4-d38c52f34035")

	// Setting up the expected call with the exact UUID
	mockRepo.EXPECT().
		GetConcertByID(gomock.Any(), testUUID).
		Return(&ent.Concert{}, nil). // Replace with the actual expected return type
		Times(1)

	// Creating a fake HTTP context with the expected UUID
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: testUUID.String()}}

	// Call the handler method
	concertHandler.GetConcertByIDHandler()(c)
}

func TestGetConcertsByArtistHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockConcertRepository(ctrl)
	artistName := "Taylor Swift"
	expectedConcerts := []*ent.Concert{
		{
			Artist: "Taylor Swift",
		},
	}
	mockRepo.EXPECT().
		GetConcertsByArtist(gomock.Any(), artistName).
		Return(expectedConcerts, nil).
		Times(1)

	handler := handlers.NewConcertHandler(mockRepo)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "artistName", Value: artistName},
	}

	handler.GetConcertsByArtistHandler()(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetFeaturedConcerts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockConcertRepository(ctrl)
	expectedConcerts := []*ent.Concert{}
	mockRepo.EXPECT().
		GetFeaturedConcerts(gomock.Any()).
		Return(expectedConcerts, nil).
		Times(1)

	handler := handlers.NewConcertHandler(mockRepo)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	handler.GetFeaturedConcerts()(c)
	assert.Equal(t, http.StatusOK, w.Code)
}
