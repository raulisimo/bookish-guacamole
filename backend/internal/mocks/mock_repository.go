package mocks

import "backend/internal/models"

// MockSwapiRepository is a mock implementation of the SwapiRepository interface.
type MockSwapiRepository struct {
	PeopleData []models.Person
}

// NewMockSwapiRepository creates a new mock repository.
func NewMockSwapiRepository() *MockSwapiRepository {
	return &MockSwapiRepository{}
}

// GetPeople mocks the GetPeople function in the repository.
func (m *MockSwapiRepository) GetPeople() ([]models.Person, error) {
	return m.PeopleData, nil
}

func (m *MockSwapiRepository) GetPlanets() ([]models.Planet, error) {
	return nil, nil // Not used in this test
}
