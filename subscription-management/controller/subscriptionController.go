package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"subscription-management/initializers"
	"subscription-management/model"
)

func GetSubscriptions(c *gin.Context) {
	userID := c.GetUint("userID")

	type SubscriptionResponse struct {
		ID          uint    `json:"id"`
		Name        string  `json:"name"`
		Amount      float64 `json:"amount"`
		Cycle       string  `json:"cycle"`
		RenewalDate string  `json:"renewal_date"`
		Status      string  `json:"status"`
		Notes       string  `json:"notes"`
	}

	// Dans le handler :
	var subscriptions []model.Subscription
	initializers.DB.Where("user_id = ?", userID).Find(&subscriptions)

	var result []SubscriptionResponse
	for _, s := range subscriptions {
		result = append(result, SubscriptionResponse{
			ID:          s.ID,
			Name:        s.Name,
			Amount:      s.Amount,
			Cycle:       s.Cycle,
			RenewalDate: s.RenewalDate.Format("01-02-2006"),
			Status:      s.Status,
			Notes:       s.Notes,
		})
	}

	c.JSON(200, result)
}

func CreateSubscription(c *gin.Context) {
	userID := c.GetUint("userID")

	var subscription model.Subscription
	if err := c.ShouldBindJSON(&subscription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subscription.UserID = userID

	initializers.DB.Create(&subscription)
	c.JSON(http.StatusOK, subscription)
}

func GetSubscription(c *gin.Context) {
	var subscription model.Subscription
	initializers.DB.First(&subscription, c.Param("id"))
	c.JSON(http.StatusOK, subscription)
}

func GetSubscriptionTotal(c *gin.Context) {
	var total int64
	initializers.DB.Model(&model.Subscription{}).Count(&total)
	c.JSON(http.StatusOK, gin.H{"total": total})
}

func UpdateSubscription(c *gin.Context) {
	var subscription model.Subscription
	if err := c.ShouldBindJSON(&subscription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	initializers.DB.Save(&subscription)
	c.JSON(http.StatusOK, subscription)
}

func DeleteSubscription(c *gin.Context) {
	var subscription model.Subscription
	initializers.DB.Delete(&subscription, c.Param("id"))
	c.JSON(http.StatusOK, subscription)
}
