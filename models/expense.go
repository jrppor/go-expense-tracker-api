package models

import (
	"time"

	"gorm.io/gorm"
)

// Expense represents an expense record stored in the database.
type Expense struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	CategoryID    *uint          `gorm:"index" json:"category_id,omitempty"`
	Amount        float64        `gorm:"not null" json:"amount"`
	Description   string         `gorm:"type:text" json:"description"`
	PaymentMethod string         `gorm:"type:varchar(50)" json:"payment_method"`
	Tags          string         `gorm:"type:varchar(255)" json:"tags"`
	Date          time.Time      `gorm:"not null" json:"date"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Expense) TableName() string { return "expenses" }
