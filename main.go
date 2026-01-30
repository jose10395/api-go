package main

import (
	handlers "api-go/src/adapters/http/handlers"
	"api-go/src/adapters/persistence/repositories"
	logic "api-go/src/domain/usecases"
	config "api-go/src/infraestructure"
	"api-go/src/infraestructure/database"
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
	userRepository := repositories.NewUserRepositoryImpl(db)
	emergencyRepository := repositories.NewEmergencyRepositoryImpl(db)

	// servicios
	userService := logic.NewUserUsecase(userRepository)
	emergencyService := logic.NewEmergencyUsecase(emergencyRepository)

	// handlers
	userHandler := handlers.NewUserHandler(userService)
	emergencyHandler := handlers.NewEmergencyHandler(emergencyService)

	routes.SetupRoutes(r, userHandler, emergencyHandler)

	r.Run(":8080")
}
