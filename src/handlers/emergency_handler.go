package http

import (
	"api-go/src/domain/entities"
	logic "api-go/src/domain/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmergencyHandler struct {
	emergencyService *logic.EmergencyService
}

func NewEmergenciaHandler(emergencyService *logic.EmergencyService) *EmergencyHandler {
	return &EmergencyHandler{
		emergencyService: emergencyService,
	}
}

// POST /emergencias
func (uh *EmergencyHandler) Create(c *gin.Context) {
	var input entities.Emergency

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	if err := uh.emergencyService.Create(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

// GET /emergencias/:id
func (uh *EmergencyHandler) FindByID(c *gin.Context) {
	idParam := c.Param("id")
	idUint, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	user, err := uh.emergencyService.FindByID(uint(idUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GET /emergencias
func (uh *EmergencyHandler) FindAll(c *gin.Context) {
	emergencias, err := uh.emergencyService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, emergencias)
}
