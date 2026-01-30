package http

import (
	"api-go/src/domain/entities"
	logic "api-go/src/domain/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase *logic.UserUsecase
}

func NewUserHandler(userUsecase *logic.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

// POST /users
func (uh *UserHandler) Create(c *gin.Context) {
	var input entities.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	if err := uh.userUsecase.Create(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

// GET /users/:id
func (uh *UserHandler) FindByID(c *gin.Context) {
	idParam := c.Param("id")
	idUint, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	user, err := uh.userUsecase.FindByID(uint(idUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GET /users
func (uh *UserHandler) FindAll(c *gin.Context) {
	users, err := uh.userUsecase.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
