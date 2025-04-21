package handlers

import (
	"net/http"
	"nested-app/internal/models"
	"nested-app/internal/services"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	service services.CommentService
}

func NewCommentHandler(service services.CommentService) *CommentHandler {
	return &CommentHandler{service}
}

func (h *CommentHandler) Create(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.service.Create(&comment)
	c.JSON(http.StatusCreated, comment)
}