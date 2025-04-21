package handlers

import (
	"nested-app/internal/models"
	"nested-app/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service services.PostService
}

func NewPostHandler(service services.PostService) *PostHandler {
	return &PostHandler{service}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.service.Create(&post)
	c.JSON(http.StatusCreated, post)
}
