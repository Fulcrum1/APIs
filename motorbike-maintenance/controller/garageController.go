package controller

import (
	"fmt"
	"mototrack/initializers"
	"mototrack/model"

	"github.com/gin-gonic/gin"
)

func GetGarage(c *gin.Context) {
	userId := c.GetUint("userID")
	fmt.Println(userId)

	var garage model.Garage
	result := initializers.DB.Where("user_id = ?", userId).First(&garage)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Garage not found"})
		return
	}

	c.JSON(200, gin.H{"data": garage})
}

func CreateGarage(c *gin.Context) {
	userId := c.GetUint("userID")
	fmt.Println(userId)

	var garage model.Garage
	if err := c.ShouldBindJSON(&garage); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	garage.UserID = userId
	result := initializers.DB.Create(&garage)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create garage"})
		return
	}

	c.JSON(201, gin.H{"data": garage})
}
