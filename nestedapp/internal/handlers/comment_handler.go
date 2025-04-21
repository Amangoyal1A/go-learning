package handlers

import (
	"net/http"
	"nested-app/internal/model"
	"nested-app/internal/service"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	service service.CommentService
}

func NewCommentHandler(service service.CommentService) *CommentHandler {
	return &CommentHandler{service}
}

func (h *CommentHandler) Create(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.service.Create(&comment)
	c.JSON(http.StatusCreated, comment)
}