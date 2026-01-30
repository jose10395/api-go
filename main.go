package main

import (
	config "api-go/src"
	"api-go/src/data/database"
	"api-go/src/data/repositories"
	logic "api-go/src/domain/services"
	handlers "api-go/src/handlers"
	"api-go/src/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfigFromEnv()

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	r := gin.Default()

	// repositorios
	userRepository := repositories.NewUserGormRepository(db)
	emergencyRepository := repositories.NewEmergenciaGormRepository(db)

	// servicios
	userService := logic.NewUserService(userRepository)
	emergencyService := logic.NewEmergencyService(emergencyRepository)

	// handlers
	userHandler := handlers.NewUserHandler(userService)
	emergencyHandler := handlers.NewEmergenciaHandler(emergencyService)

	routes.SetupRoutes(r, userHandler, emergencyHandler)

	r.Run(":8080")
}
