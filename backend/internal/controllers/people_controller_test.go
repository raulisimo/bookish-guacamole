package controllers

import (
	"backend/internal/mocks"
	"backend/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPeople_Success(t *testing.T) {
	// Mock the PeopleService
	mockService := mocks.NewMockPeopleService()
	mockService.PeopleData = []models.Person{
		{Name: "Luke", Created: "2024-01-01T00:00:00Z"},
		{Name: "Leia", Created: "2023-01-01T00:00:00Z"},
	}

	// Create the controller with the mock service
	controller := NewPeopleController(mockService)

	// Create a test router and define the route
	r := gin.Default()
	r.GET("/people", controller.GetPeople)

	// Create a mock HTTP request
	req, _ := http.NewRequest("GET", "/people", nil)
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Luke")
	assert.Contains(t, w.Body.String(), "Leia")
}

func TestGetPeople_Failure(t *testing.T) {
	// Mock the PeopleService with an error
	mockService := mocks.NewMockPeopleService()
	mockService.Err = assert.AnError // Simulate an error

	// Create the controller with the mock service
	controller := NewPeopleController(mockService)

	// Create a test router and define the route
	r := gin.Default()
	r.GET("/people", controller.GetPeople)

	// Create a mock HTTP request
	req, _ := http.NewRequest("GET", "/people", nil)
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}
