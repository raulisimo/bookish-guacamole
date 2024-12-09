package repositories

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"
	"time"

	"backend/internal/models"
)

// MockCache is a mock implementation of cache.Cache
type MockCache struct {
	data map[string]interface{}
}

func (m *MockCache) Set(key string, value interface{}, ttl time.Duration) {
	if m.data == nil {
		m.data = make(map[string]interface{})
	}
	m.data[key] = value
}

func (m *MockCache) Get(key string) (interface{}, bool) {
	value, found := m.data[key]
	return value, found
}

// MockHTTPClient is a mock implementation of HTTP client
type MockHTTPClient struct {
	ResponseMap map[string]*http.Response
	Err         error
}

func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	if resp, found := m.ResponseMap[url]; found {
		return resp, nil
	}
	return nil, errors.New("unexpected URL")
}

func TestSwapiRepository_GetPlanets(t *testing.T) {
	// Mock SWAPI response
	mockPlanetsResponse := `{
		"results": [{"name": "Tatooine", "climate": "arid"}],
		"next": ""
	}`

	client := &MockHTTPClient{
		ResponseMap: map[string]*http.Response{
			"https://swapi.dev/api/planets": {
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(mockPlanetsResponse)),
			},
		},
	}

	cache := &MockCache{}
	repo := NewSwapiRepository("https://swapi.dev/api", client, cache)

	planets, err := repo.GetPlanets()
	if err != nil {
		t.Fatalf("GetPlanets returned error: %v", err)
	}

	if len(planets) != 1 || planets[0].Name != "Tatooine" {
		t.Errorf("Unexpected planets data: %+v", planets)
	}

	// Verify caching
	cachedData, found := cache.Get("planets")
	if !found {
		t.Errorf("Expected planets data to be cached")
	}

	cachedPlanets, ok := cachedData.([]models.Planet)
	if !ok || len(cachedPlanets) != 1 || cachedPlanets[0].Name != "Tatooine" {
		t.Errorf("Unexpected cached planets data: %+v", cachedData)
	}
}

func TestSwapiRepository_GetPeople(t *testing.T) {
	// Mock SWAPI response
	mockPeopleResponse := `{
		"results": [{"name": "Luke Skywalker", "gender": "male"}],
		"next": ""
	}`

	client := &MockHTTPClient{
		ResponseMap: map[string]*http.Response{
			"https://swapi.dev/api/people": {
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(mockPeopleResponse)),
			},
		},
	}

	cache := &MockCache{}
	repo := NewSwapiRepository("https://swapi.dev/api", client, cache)

	people, err := repo.GetPeople()
	if err != nil {
		t.Fatalf("GetPeople returned error: %v", err)
	}

	if len(people) != 1 || people[0].Name != "Luke Skywalker" {
		t.Errorf("Unexpected people data: %+v", people)
	}

	// Verify caching
	cachedData, found := cache.Get("people")
	if !found {
		t.Errorf("Expected people data to be cached")
	}

	cachedPeople, ok := cachedData.([]models.Person)
	if !ok || len(cachedPeople) != 1 || cachedPeople[0].Name != "Luke Skywalker" {
		t.Errorf("Unexpected cached people data: %+v", cachedData)
	}
}

func TestSwapiRepository_ErrorHandling(t *testing.T) {
	client := &MockHTTPClient{
		Err: errors.New("network error"),
	}

	cache := &MockCache{}
	repo := NewSwapiRepository("https://swapi.dev/api", client, cache)

	_, err := repo.GetPlanets()
	if err == nil {
		t.Errorf("Expected error from GetPlanets, got nil")
	}

	_, err = repo.GetPeople()
	if err == nil {
		t.Errorf("Expected error from GetPeople, got nil")
	}
}
