package http

import (
	"api-go/src/domain/entities"
	logic "api-go/src/domain/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmergenciaHandler struct {
	emergenciaService *logic.EmergenciaService
}

func NewEmergenciaHandler(emergenciaService *logic.EmergenciaService) *EmergenciaHandler {
	return &EmergenciaHandler{
		emergenciaService: emergenciaService,
	}
}

func (h *EmergenciaHandler) RegisterRoutes(r *gin.Engine) {
	emergencias := r.Group("/emergencias")
	{
		emergencias.POST("", h.Create)
		emergencias.GET("", h.FindAll)
		emergencias.GET("/:id", h.FindByID)
	}
}

// POST /emergencias
func (uh *EmergenciaHandler) Create(c *gin.Context) {
	var input entities.Emergencia

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	if err := uh.emergenciaService.Create(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

// GET /emergencias/:id
func (uh *EmergenciaHandler) FindByID(c *gin.Context) {
	idParam := c.Param("id")
	idUint, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	user, err := uh.emergenciaService.FindByID(uint(idUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GET /emergencias
func (uh *EmergenciaHandler) FindAll(c *gin.Context) {
	emergencias, err := uh.emergenciaService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, emergencias)
}
