package handler

import (
	"go-app/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	RID  string `json:"request_id"`
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	userName := service.GetUserName(c, id)

	resp := User{
		ID:   id,
		Name: userName,
		RID:  "",
	}

	c.JSON(http.StatusOK, resp)
}
