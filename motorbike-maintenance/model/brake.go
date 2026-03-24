package model

import "time"

type Brake struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	BikeID     uint      `json:"bike_id" gorm:"not null;index"`
	Usage      string    `json:"usage" gorm:"not null"`
	LastChange time.Time `json:"last_change" gorm:"type:date"`
	NextChange time.Time `json:"next_change" gorm:"type:date"`
	Notes      string    `json:"notes" gorm:"type:text"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	Bike Bike `json:"bike" gorm:"foreignKey:BikeID"`
}
