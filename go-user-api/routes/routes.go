package routes

import (
	"go-user-api/controllers"
	"go-user-api/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	apiService := services.NewApiService()
	apiController := controllers.NewApiController(apiService)

	router.POST("/users", userController.CreateUser)
	router.GET("/users", userController.GetUsers)
	router.GET("/panic", userController.CausePanic)
	router.GET("/api", apiController.GetApi)
}
