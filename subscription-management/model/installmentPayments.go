package model

import "time"

type InstallmentPayment struct {
	ID            uint `gorm:"primaryKey" json:"id"`
	PaymentPlanID uint `gorm:"not null;index" json:"payment_plan_id"`
	InstallmentNo int  `gorm:"not null" json:"installment_no"`

	Amount     float64 `gorm:"type:decimal(10,2);not null" json:"amount"`
	PaidAmount float64 `gorm:"type:decimal(10,2)" json:"paid_amount"`

	DueDate time.Time  `gorm:"not null" json:"due_date"`
	PaidAt  *time.Time `json:"paid_at,omitempty"`

	Status string `gorm:"type:varchar(20);not null" json:"status"` // pending, paid, overdue

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
