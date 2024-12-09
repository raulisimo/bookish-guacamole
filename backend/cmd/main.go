package main

import (
	"net/http"
	"time"

	"backend/internal/cache"
	"backend/internal/config"
	"backend/internal/controllers"
	"backend/internal/repositories"
	"backend/internal/routes"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
)

type AppContainer struct {
	Router           *gin.Engine
	PlanetController *controllers.PlanetController
	PeopleController *controllers.PeopleController
	Port             string
}

func NewAppContainer() *AppContainer {
	// Load configuration
	cfg := config.LoadConfig()

	// Create an HTTP client
	httpClient := &http.Client{Timeout: 10 * time.Second}

	// Initialize cache
	cache := cache.NewInMemoryCache()

	// Initialize repositories
	repo := repositories.NewSwapiRepository(cfg.BaseURL, httpClient, cache)

	// Initialize services
	planetService := services.NewPlanetService(repo)
	peopleService := services.NewPeopleService(repo)

	// Initialize controllers
	planetController := controllers.NewPlanetController(planetService)
	peopleController := controllers.NewPeopleController(peopleService)

	// Set up routes
	router := routes.SetupRouter(planetController, peopleController)

	return &AppContainer{
		Router:           router,
		PlanetController: planetController,
		PeopleController: peopleController,
		Port:             cfg.Port,
	}
}

func main() {
	// Create the application container
	container := NewAppContainer()

	// Start the server using the configured port
	if err := container.Router.Run(":" + container.Port); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
