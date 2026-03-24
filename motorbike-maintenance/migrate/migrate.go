package main

import (
	"mototrack/initializers"
	"mototrack/model"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {

	initializers.DB.AutoMigrate(
		&model.User{},
		&model.Bike{},
		&model.Garage{},
		&model.Chain{},
		&model.Tire{},
		&model.Gear{},
		&model.Invoice{},
		&model.Brake{},
		&model.Motor{},
	)
}
