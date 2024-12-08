package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"

	"backend/internal/models"
)

// SwapiRepository defines the interface for interacting with SWAPI
type SwapiRepository interface {
	GetPlanets(limit int) ([]models.Planet, error) // Accept limit as parameter
	GetPeople(limit int) ([]models.Person, error)  // Accept limit as parameter
}

// swapiRepository is the concrete implementation of SwapiRepository
type swapiRepository struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewSwapiRepository creates a new instance of swapiRepository
func NewSwapiRepository(baseURL string, client *http.Client) SwapiRepository {
	return &swapiRepository{BaseURL: baseURL, HTTPClient: client}
}

// GetPlanets fetches up to 'limit' planets from SWAPI, handling pagination
func (r *swapiRepository) GetPlanets(limit int) ([]models.Planet, error) {
	var allPlanets []models.Planet
	url := fmt.Sprintf("%s/planets", r.BaseURL)
	for {
		// Fetch the next page of results
		var response struct {
			Results []models.Planet `json:"results"`
			Next    string          `json:"next"` // URL of the next page, if any
		}

		if err := r.fetchResource(url, &response); err != nil {
			return nil, fmt.Errorf("failed to fetch planets: %w", err)
		}

		allPlanets = append(allPlanets, response.Results...)

		// If we have enough items, return them
		if len(allPlanets) >= limit {
			return allPlanets[:limit], nil
		}

		// If there's a next page, set the URL to the next page; otherwise, stop
		if response.Next == "" {
			break
		}
		url = response.Next
	}

	// Return the fetched planets (may be fewer than limit if no more pages)
	return allPlanets, nil
}

// GetPeople fetches up to 'limit' people from SWAPI, handling pagination
func (r *swapiRepository) GetPeople(limit int) ([]models.Person, error) {
	var allPeople []models.Person
	url := fmt.Sprintf("%s/people", r.BaseURL)
	for {
		// Fetch the next page of results
		var response struct {
			Results []models.Person `json:"results"`
			Next    string          `json:"next"` // URL of the next page, if any
		}

		if err := r.fetchResource(url, &response); err != nil {
			return nil, fmt.Errorf("failed to fetch people: %w", err)
		}

		allPeople = append(allPeople, response.Results...)

		// If we have enough items, return them
		if len(allPeople) >= limit {
			return allPeople[:limit], nil
		}

		// If there's a next page, set the URL to the next page; otherwise, stop
		if response.Next == "" {
			break
		}
		url = response.Next
	}

	// Return the fetched people (may be fewer than limit if no more pages)
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
