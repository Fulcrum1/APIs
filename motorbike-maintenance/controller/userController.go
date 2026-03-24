package controller

import (
	"mototrack/initializers"
	"mototrack/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []model.User
	initializers.DB.Select("email", "name").Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
