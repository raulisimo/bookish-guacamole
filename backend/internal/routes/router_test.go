package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// Mock Controllers
type MockPlanetController struct {
	mock.Mock
}

func (m *MockPlanetController) GetPlanets(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "mocked planets"})
}

type MockPeopleController struct {
	mock.Mock
}

func (m *MockPeopleController) GetPeople(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "mocked people"})
}

func TestSetupRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mock controllers
	mockPlanetController := &MockPlanetController{}
	mockPeopleController := &MockPeopleController{}

	// Create the router
	router := SetupRouter(mockPlanetController, mockPeopleController)

	tests := []struct {
		method       string
		route        string
		expectedCode int
		expectedBody string
	}{
		{
			method:       "GET",
			route:        "/",
			expectedCode: http.StatusOK,
			expectedBody: `{"message":"Welcome to the Star Wars API!"}`,
		},
		{
			method:       "GET",
			route:        "/api/planets",
			expectedCode: http.StatusOK,
			expectedBody: `{"data":"mocked planets"}`,
		},
		{
			method:       "GET",
			route:        "/api/people",
			expectedCode: http.StatusOK,
			expectedBody: `{"data":"mocked people"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.route, func(t *testing.T) {
			// Create a test HTTP request
			req := httptest.NewRequest(test.method, test.route, nil)
			resp := httptest.NewRecorder()

			// Serve the request
			router.ServeHTTP(resp, req)

			// Check status code
			if resp.Code != test.expectedCode {
				t.Errorf("Expected status %d, got %d", test.expectedCode, resp.Code)
			}

			// Check response body
			if resp.Body.String() != test.expectedBody {
				t.Errorf("Expected body %s, got %s", test.expectedBody, resp.Body.String())
			}
		})
	}
}
