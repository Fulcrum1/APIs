package model

import "time"

type Gear struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`
	Amount    int       `json:"amount" gorm:"not null"`
	Date      time.Time `json:"date" gorm:"type:date;not null"`
	Notes     string    `json:"notes" gorm:"type:text"`
	Type      string    `json:"type" gorm:"not null"` // "boots", "pants", etc.
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	User User `json:"user" gorm:"foreignKey:UserID"`
}
