package routes

import (
	"nested-app/internal/handlers"
	"nested-app/internal/repositories"
	"nested-app/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	userRepo := repositories.NewUserRepo()
	postRepo := repositories.NewPostRepo()
	commentRepo := repositories.NewCommentRepo()

	userService := services.NewUserService(userRepo)
	postService := services.NewPostService(postRepo)
	commentService := services.NewCommentService(commentRepo)

	hUser := handlers.NewUserHandler(userService)
	hPost := handlers.NewPostHandler(postService)
	hComment := handlers.NewCommentHandler(commentService)

	r.POST("/users", hUser.CreateUser)
	r.GET("/users/:id", hUser.GetUser)
	r.DELETE("/users/:id", hUser.DeleteUser)

	r.POST("/posts", hPost.CreatePost)

	r.POST("/comments", hComment.Create)
}
