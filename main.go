package main

import (
	config "api-go/src"
	"api-go/src/data/database"
	"api-go/src/data/repositories"
	logic "api-go/src/domain/services"
	httpInterfaces "api-go/src/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("Hola Joven, su hua de API está lista 🚀")

	cfg := config.LoadConfigFromEnv()

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	r := gin.Default()

	userRepository := repositories.NewUserGormRepository(db)
	userService := logic.NewUserService(userRepository)
	userHandler := httpInterfaces.NewUserHandler(userService)
	userHandler.RegisterRoutes(r)

	emergenciaRepository := repositories.NewEmergenciaGormRepository(db)
	emergenciaService := logic.NewEmergenciaService(emergenciaRepository)
	emergenciaHandler := httpInterfaces.NewEmergenciaHandler(emergenciaService)
	emergenciaHandler.RegisterRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Joven, su API creo que ya sirve... 🚀",
		})
	})

	r.Run(":8080")
}
