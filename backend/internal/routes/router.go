package routes

import (
	"backend/internal/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(planetController *controllers.PlanetController, peopleController *controllers.PeopleController) *gin.Engine {
	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Allow the frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true, // Allow cookies or authentication headers
		MaxAge:           12 * time.Hour,
	}))

	// Define routes

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Star Wars API!",
		})
	})

	api := router.Group("/api")
	{
		api.GET("/planets", planetController.GetPlanets)
		api.GET("/people", peopleController.GetPeople)
	}



	return router
}
