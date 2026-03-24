package model

import "time"

type Chain struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	BikeID        uint      `json:"bike_id" gorm:"not null;index"`
	LastCleanDate time.Time `json:"last_clean_date" gorm:"type:date"`
	LastCleanKm   int       `json:"last_clean_km" gorm:"default:0"`
	Notes         string    `json:"notes" gorm:"type:text"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	Bike Bike `json:"bike" gorm:"foreignKey:BikeID"`
}
