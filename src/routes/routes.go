package routes

import (
	httpPorts "api-go/src/adapters/http/ports"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	userHandler httpPorts.UserHandlerPort,
	emergencyHandler httpPorts.EmergencyHandlerPort,
) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Joven, su API ya sirve... 🚀",
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
