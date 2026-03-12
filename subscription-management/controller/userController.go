package controller

import (
	"net/http"
	"subscription-management/initializers"
	"subscription-management/model"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []model.User
	initializers.DB.Select("email", "name").Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
