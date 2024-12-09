package controllers

import (
	"backend/internal/services"
	"backend/internal/sorting"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PeopleController handles requests related to people
type PeopleController struct {
	Service services.PeopleServiceInterface
}

// NewPeopleController creates a new instance of PeopleController
func NewPeopleController(service services.PeopleServiceInterface) *PeopleController {
	return &PeopleController{Service: service}
}

// GetPeople handles fetching people with optional sorting
func (c *PeopleController) GetPeople(ctx *gin.Context) {
	sortParam := ctx.DefaultQuery("sort", "")  // Default to empty string if not provided
	orderParam := ctx.DefaultQuery("order", "asc") // Default to "asc" if not provided

	// Use the SorterFactory to create the appropriate sorter based on the "sort" query parameter
	sorter, err := sorting.SorterFactory(sortParam)
	if err != nil {
		// Return a 400 Bad Request if an invalid sort field is provided
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch and sort the people using the PeopleService. If sorter is nil, no sorting is applied
	people, err := c.Service.GetPeople(sorter, orderParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the sorted people (or unsorted if no sorter was applied) as a JSON response
	ctx.JSON(http.StatusOK, people)
}