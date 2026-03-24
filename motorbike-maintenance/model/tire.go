package model

import "time"

type Tire struct {
	ID                  uint      `json:"id" gorm:"primaryKey"`
	BikeID              uint      `json:"bike_id" gorm:"not null;index"`
	Make                string    `json:"make" gorm:"not null"`
	Model               string    `json:"model" gorm:"not null"`
	Size                string    `json:"size" gorm:"not null"`
	MinPressure         int       `json:"min_pressure" gorm:"not null"`
	RecommendedPressure int       `json:"recommended_pressure" gorm:"not null"`
	MaxPressure         int       `json:"max_pressure" gorm:"not null"`
	MaxKilometers       int       `json:"max_kilometers" gorm:"default:0"`
	LastChange          time.Time `json:"last_change" gorm:"type:date"`
	Notes               string    `json:"notes" gorm:"type:text"`
	CreatedAt           time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt           time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	Bike Bike `json:"bike" gorm:"foreignKey:BikeID"`
}
