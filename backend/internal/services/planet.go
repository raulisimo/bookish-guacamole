package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/sorting"
)

// PlanetService handles the business logic for planets
type PlanetService struct {
	Repository repositories.SwapiRepository
}

// NewPlanetService creates a new instance of PlanetService
func NewPlanetService(repository repositories.SwapiRepository) *PlanetService {
	return &PlanetService{Repository: repository}
}

// GetPlanets fetches and sorts the planets based on the provided sorter
func (s *PlanetService) GetPlanets(sorter sorting.Sorter, order string) ([]models.Planet, error) {
	planets, err := s.Repository.GetPlanets()  // Fetch all planets
	if err != nil {
		return nil, err
	}

	// Sort planets if a sorter is provided.
	if sorter != nil {
		err := sorter.Sort(planets, order)
		if err != nil {
			return nil, err
		}
	}

	return planets, nil
}
