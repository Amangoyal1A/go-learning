package controllers

import (
	"fmt"
	"net/http"
	"worker-app/config"
	"worker-app/models"

	"github.com/gin-gonic/gin"
)

type JobRequest struct {
	Task string `json:"task"`
}

func CreateJob(c *gin.Context) {
	var req JobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	jobs := make([]models.Job, 0, 100000)

	for i := 0; i < 100000; i++ {
		jobs = append(jobs, models.Job{
			Task:   fmt.Sprintf("Task-%d", i+1),
			Status: "pending",
		})
	}
	
	// Use CreateInBatches for better performance
	if err := config.DB.CreateInBatches(jobs, 1000).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create jobs"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "100000 jobs inserted"})

}
