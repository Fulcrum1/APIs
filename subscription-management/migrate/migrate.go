package main

import (
	"subscription-management/initializers"
	"subscription-management/model"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {

	initializers.DB.AutoMigrate(
		&model.User{},
		&model.Subscription{},
		&model.InstallmentPayment{},
		&model.Setting{})
}
