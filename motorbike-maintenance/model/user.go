package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Name      string    `json:"name" gorm:"not null"`
	Password  string    `json:"-" gorm:"not null"` // Hidden in JSON
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	Bikes    []Bike    `json:"bikes" gorm:"foreignKey:UserID"`
	Garages  []Garage  `json:"garages" gorm:"foreignKey:UserID"`
	Gears    []Gear    `json:"gears" gorm:"foreignKey:UserID"`
	Invoices []Invoice `json:"invoices" gorm:"foreignKey:UserID"`
}
