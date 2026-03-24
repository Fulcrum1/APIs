package controller

import (
	"github.com/gin-gonic/gin"

	"mototrack/initializers"
	"mototrack/model"
)

func GetBike(c *gin.Context) {
	bikeId := c.Param("id")
	var bike model.Bike
	initializers.DB.First(&bike, bikeId)
	if bike.ID == 0 {
		c.JSON(404, gin.H{"error": "Bike not found"})
		return
	}
	c.JSON(200, gin.H{"data": bike})
}

func GetBikes(c *gin.Context) {
	var bikes []model.Bike
	userId := c.GetHeader("user_id")
	initializers.DB.Where("user_id = ?", userId).Find(&bikes)
	c.JSON(200, gin.H{"data": bikes})
}

func CreateBike(c *gin.Context) {
	var bike model.Bike
	c.ShouldBindJSON(&bike)
	initializers.DB.Create(&bike)
	c.JSON(201, gin.H{"data": bike})
}

func UpdateBike(c *gin.Context) {
	var bike model.Bike
	c.ShouldBindJSON(&bike)
	initializers.DB.Save(&bike)
	c.JSON(200, gin.H{"data": bike})
}

func DeleteBike(c *gin.Context) {
	bikeId := c.Param("id")
	var bike model.Bike
	initializers.DB.Delete(&bike, bikeId)
	c.JSON(200, gin.H{"data": "Bike deleted"})
}
