package mocks

import (
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/sorting"
)

// MockPeopleService is a mock implementation of the PeopleService
type MockPeopleService struct {
	PeopleData []models.Person
	Err        error
}

// Ensure MockPeopleService implements the PeopleService interface
var _ services.PeopleServiceInterface = (*MockPeopleService)(nil)

// NewMockPeopleService creates a new mock PeopleService
func NewMockPeopleService() *MockPeopleService {
	return &MockPeopleService{}
}

// GetPeople mocks the GetPeople method in the PeopleService
func (m *MockPeopleService) GetPeople(sorter sorting.Sorter, order string) ([]models.Person, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.PeopleData, nil
}