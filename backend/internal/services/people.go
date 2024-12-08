package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/sorting"
)

// PeopleService handles the business logic for people.
type PeopleService struct {
	Repository repositories.SwapiRepository
}

// NewPeopleService creates a new instance of PeopleService.
func NewPeopleService(repository repositories.SwapiRepository) *PeopleService {
	return &PeopleService{Repository: repository}
}

// GetPeople fetches and sorts the people based on the provided sorter.
func (s *PeopleService) GetPeople(sorter sorting.Sorter, order string) ([]models.Person, error) {
	people, err := s.Repository.GetPeople(15)
	if err != nil {
		return nil, err
	}

	// Sort people if a sorter is provided.
	if sorter != nil {
		err := sorter.Sort(people, order)
		if err != nil {
			return nil, err
		}
	}

	return people, nil
}
