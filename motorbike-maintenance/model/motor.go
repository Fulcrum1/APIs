package model

import "time"

type Motor struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	BikeID    uint      `json:"bike_id" gorm:"not null;index"`
	Oil       string    `json:"oil" gorm:"not null"`
	OilFilter string    `json:"oil_filter" gorm:"not null"`
	ChangeOil time.Time `json:"change_oil" gorm:"type:date"`
	NextOil   time.Time `json:"next_oil" gorm:"type:date"`
	Notes     string    `json:"notes" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	Bike Bike `json:"bike" gorm:"foreignKey:BikeID"`
}
