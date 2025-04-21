package controllers

import (
	"net/http"
	"myapp/models"

	"github.com/gin-gonic/gin"
)

// CreateUser handles POST /api/user request
func CreateUser(c *gin.Context) {
	var user models.User
	// Bind JSON body to user struct; ShouldBindJSON returns an error if binding fails.
	if err := c.ShouldBindJSON(&user); err != nil {
		// Abort the request and return 400 Bad Request with error message.
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// In real scenario, you would save the user to the database.

	// Return success response in JSON format.
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

// GetUser handles GET /api/user/:id request
func GetUser(c *gin.Context) {
	// Retrieve parameter from URL path.
	id := c.Param("id")

	// Check if middleware has set a "userID" key in context.
	if contextUser, exists := c.Get("userID"); exists {
		c.JSON(http.StatusOK, gin.H{
			"message":     "User found",
			"userIdFromPath": id,
			"contextUser": contextUser,
		})
		return
	}

	// Default response if no context key exists.
	c.JSON(http.StatusOK, gin.H{
		"message": "User found",
		"id":      id,
	})
}
