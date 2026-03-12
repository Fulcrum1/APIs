package model

import "time"

type Subscription struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Amount      float64   `json:"amount"`
	Cycle       string    `json:"cycle"`
	RenewalDate time.Time `json:"renewal_date"`
	Status      string    `json:"status"`
	Notes       string    `json:"notes"`

	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
