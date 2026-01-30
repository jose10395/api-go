package routes

import (
	handlers "api-go/src/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	userHandler *handlers.UserHandler,
	emergencyHandler *handlers.EmergencyHandler,
) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Joven, su API creo que ya sirve... 🚀",
		})
	})

	users := r.Group("/users")
	{
		users.POST("", userHandler.Create)
		users.GET("", userHandler.FindAll)
		users.GET("/:id", userHandler.FindByID)
	}

	emergencies := r.Group("/emergencies")
	{
		emergencies.POST("", emergencyHandler.Create)
		emergencies.GET("", emergencyHandler.FindAll)
		emergencies.GET("/:id", emergencyHandler.FindByID)
	}
}
