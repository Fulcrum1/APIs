package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"subscription-management/initializers"
	"subscription-management/model"
)

func GetInstallmentPayments(c *gin.Context) {
	var installmentPayments []model.InstallmentPayment
	initializers.DB.Find(&installmentPayments)

	c.JSON(http.StatusOK, installmentPayments)
}

func GetInstallmentPayment(c *gin.Context) {
	var installmentPayment model.InstallmentPayment
	initializers.DB.First(&installmentPayment, c.Param("id"))
	c.JSON(http.StatusOK, installmentPayment)
}

func CreateInstallmentPayment(c *gin.Context) {
	var installmentPayment model.InstallmentPayment
	if err := c.ShouldBindJSON(&installmentPayment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	initializers.DB.Create(&installmentPayment)
	c.JSON(http.StatusOK, installmentPayment)
}

func UpdateInstallmentPayment(c *gin.Context) {
	var installmentPayment model.InstallmentPayment
	if err := c.ShouldBindJSON(&installmentPayment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	initializers.DB.Save(&installmentPayment)
	c.JSON(http.StatusOK, installmentPayment)
}

func DeleteInstallmentPayment(c *gin.Context) {
	var installmentPayment model.InstallmentPayment
	initializers.DB.Delete(&installmentPayment, c.Param("id"))
	c.JSON(http.StatusOK, installmentPayment)
}
