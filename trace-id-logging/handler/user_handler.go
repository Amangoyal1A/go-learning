// handler/user_handler.go
package handler

import (
	"log"
	"net/http"
	"trace-id-logging/middleware"
	"trace-id-logging/models"
	"trace-id-logging/repository"
	"trace-id-logging/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler() *UserHandler {
	repo := repository.NewUserRepository()
	service := service.NewUserService(repo)
	return &UserHandler{userService: service}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	traceID := c.GetString(middleware.RequestIDKey)
	id := c.Param("id")

	log.Printf("[%s] GetUser called with ID = %s", traceID, id)

	user, err := h.userService.GetUser(id)
	if err != nil {
		log.Printf("[%s] Error: %v", traceID, err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	traceID := c.GetString(middleware.RequestIDKey)

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("[%s] Bad request: %v", traceID, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, err := h.userService.CreateUser(user)
	if err != nil {
		log.Printf("[%s] Error creating user: %v", traceID, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not save user"})
		return
	}

	log.Printf("[%s] User created: %+v", traceID, user)
	c.JSON(http.StatusCreated, user)
}
