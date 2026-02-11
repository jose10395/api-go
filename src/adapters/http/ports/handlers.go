package http

import "github.com/gin-gonic/gin"

// UserHandlerPort define los métodos HTTP que exponemos para usuarios.
type UserHandlerPort interface {
	Create(c *gin.Context)
	FindAll(c *gin.Context)
	FindByID(c *gin.Context)
}

// EmergencyHandlerPort define los métodos HTTP que exponemos para emergencias.
type EmergencyHandlerPort interface {
	Create(c *gin.Context)
	FindAll(c *gin.Context)
	FindByID(c *gin.Context)
}
