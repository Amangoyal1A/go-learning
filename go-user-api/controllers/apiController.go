package controllers

import (
	"go-user-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiController struct {
	apiService services.ApiService
}

func NewApiController(service services.ApiService) *ApiController {
	return &ApiController{apiService: service}
}

func (ctrl *ApiController) GetApi(c *gin.Context) {

	user, err := ctrl.apiService.GetApi()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
