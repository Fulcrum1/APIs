package model

import "time"

type Invoice struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`
	BikeID    uint      `json:"bike_id" gorm:"not null;index"`
	Name      string    `json:"name" gorm:"not null"`
	Amount    int       `json:"amount" gorm:"not null"`
	Date      time.Time `json:"date" gorm:"type:date;not null"`
	Notes     string    `json:"notes" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	User User `json:"user" gorm:"foreignKey:UserID"`
	Bike Bike `json:"bike" gorm:"foreignKey:BikeID"`
}
