package handlers

import (
	"net/http"
	"nested-app/internal/model"
	"nested-app/internal/service"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service service.PostService
}

func NewPostHandler(service service.PostService) *PostHandler {
	return &PostHandler{service}
}

func (h *PostHandler) Create(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.service.Create(&post)
	c.JSON(http.StatusCreated, post)
}