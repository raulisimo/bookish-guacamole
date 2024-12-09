package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"backend/internal/cache"
	"backend/internal/models"
)


// Doer defines the interface for making HTTP requests (allows mocking)
type Doer interface {
	Get(url string) (*http.Response, error)
}

// SwapiRepository defines the interface for interacting with SWAPI
type SwapiRepository interface {
	GetPlanets() ([]models.Planet, error)
	GetPeople() ([]models.Person, error)
}

// swapiRepository is the concrete implementation of SwapiRepository
type swapiRepository struct {
	BaseURL    string
	HTTPClient Doer
	Cache      cache.Cache      

}

// NewSwapiRepository creates a new instance of swapiRepository
func NewSwapiRepository(baseURL string, client Doer, cache cache.Cache) SwapiRepository {
	return &swapiRepository{BaseURL: baseURL, HTTPClient: client, Cache: cache}
}

// GetPlanets fetches all planets from SWAPI, handling pagination
func (r *swapiRepository) GetPlanets() ([]models.Planet, error) {
	cacheKey := "planets"
	if cachedData, found := r.Cache.Get(cacheKey); found {
		return cachedData.([]models.Planet), nil
	}

	var allPlanets []models.Planet
	url := fmt.Sprintf("%s/planets", r.BaseURL)
	for {
		var response struct {
			Results []models.Planet `json:"results"`
			Next    string          `json:"next"`
		}

		if err := r.fetchResource(url, &response); err != nil {
			return nil, fmt.Errorf("failed to fetch planets: %w", err)
		}

		allPlanets = append(allPlanets, response.Results...)

		if response.Next == "" {
			break
		}
		url = response.Next
	}

	// Cache the data for 15 minutes
	r.Cache.Set(cacheKey, allPlanets, 15*time.Minute)

	return allPlanets, nil
}

func (r *swapiRepository) GetPeople() ([]models.Person, error) {
	cacheKey := "people"
	if cachedData, found := r.Cache.Get(cacheKey); found {
		return cachedData.([]models.Person), nil
	}

	var allPeople []models.Person
	url := fmt.Sprintf("%s/people", r.BaseURL)
	for {
		var response struct {
			Results []models.Person `json:"results"`
			Next    string          `json:"next"`
		}

		if err := r.fetchResource(url, &response); err != nil {
			return nil, fmt.Errorf("failed to fetch people: %w", err)
		}

		allPeople = append(allPeople, response.Results...)

		if response.Next == "" {
			break
		}
		url = response.Next
	}

	// Cache the data for 15 minutes
	r.Cache.Set(cacheKey, allPeople, 15*time.Minute)

	return allPeople, nil
}

// fetchResource is a helper function for making GET requests and decoding the response
func (r *swapiRepository) fetchResource(url string, target interface{}) error {
	resp, err := r.HTTPClient.Get(url)
	if err != nil {
		return fmt.Errorf("error making GET request to %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d from %s", resp.StatusCode, url)
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("error decoding JSON response from %s: %w", url, err)
	}
	return nil
}
