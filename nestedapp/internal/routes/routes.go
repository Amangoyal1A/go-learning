package routes

import (
	"nested-app/internal/handler"
	"nested-app/internal/repository"
	"nested-app/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	userRepo := repository.NewUserRepo()
	postRepo := repository.NewPostRepo()
	commentRepo := repository.NewCommentRepo()

	userService := service.NewUserService(userRepo)
	postService := service.NewPostService(postRepo)
	commentService := service.NewCommentService(commentRepo)

	hUser := handler.NewUserHandler(userService)
	hPost := handler.NewPostHandler(postService)
	hComment := handler.NewCommentHandler(commentService)

	r.POST("/users", hUser.Create)
	r.GET("/users/:id", hUser.Get)
	r.DELETE("/users/:id", hUser.Delete)

	r.POST("/posts", hPost.Create)
	r.POST("/comments", hComment.Create)
}