package controllers

import (
	"backend/internal/services"
	"backend/internal/sorting"
	"net/http"
	"github.com/gin-gonic/gin"
)

// PlanetController handles requests related to planets.
type PlanetController struct {
	Service *services.PlanetService
}

// NewPlanetController creates a new instance of PlanetController.
func NewPlanetController(service *services.PlanetService) *PlanetController {
	return &PlanetController{Service: service}
}

// GetPlanets handles fetching planets with optional sorting.
func (c *PlanetController) GetPlanets(ctx *gin.Context) {
	sortParam := ctx.DefaultQuery("sort", "")  // Default to empty string if not provided
	orderParam := ctx.DefaultQuery("order", "asc") // Default to "asc" if not provided

	// Use the SorterFactory to create the appropriate sorter based on the "sort" query parameter.
	sorter, err := sorting.SorterFactory(sortParam)
	if err != nil {
		// Return a 400 Bad Request if an invalid sort field is provided
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch and sort the planets using the PlanetService. If sorter is nil, no sorting is applied.
	planets, err := c.Service.GetPlanets(sorter, orderParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the sorted planets (or unsorted if no sorter was applied) as a JSON response.
	ctx.JSON(http.StatusOK, planets)
}
