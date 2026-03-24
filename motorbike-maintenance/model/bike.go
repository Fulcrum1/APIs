// model/bike.go
package model

import "time"

type Bike struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`
	GarageID  uint      `json:"garage_id" gorm:"not null;index"`
	Make      string    `json:"make" gorm:"not null"`
	Model     string    `json:"model" gorm:"not null"`
	Year      int       `json:"year" gorm:"not null"`
	Type      string    `json:"type" gorm:"not null"`
	Kilometer int       `json:"kilometer" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	Garage    Garage    `json:"garage" gorm:"foreignKey:GarageID"`
	Motor     *Motor    `json:"motor" gorm:"foreignKey:BikeID"`
	Brake     *Brake    `json:"brake" gorm:"foreignKey:BikeID"`
	FrontTire *Tire     `json:"front_tire" gorm:"foreignKey:BikeID"`
	BackTire  *Tire     `json:"back_tire" gorm:"foreignKey:BikeID"`
	Chain     *Chain    `json:"chain" gorm:"foreignKey:BikeID"`
	Invoices  []Invoice `json:"invoices" gorm:"foreignKey:BikeID"`
}
