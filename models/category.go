package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID uint `gorm:"primaryKey"`

	// data
	Name        string `gorm:"size:64;not null"`
	Slug        string `gorm:"size:80;not null"`
	Description string `gorm:"size:255"`

	Color     string `gorm:"size:7"`  // เช่น #FF9900
	Icon      string `gorm:"size:64"` // เช่น "mdi-food"
	SortOrder int    `gorm:"default:0"`

	ParentID *uint `gorm:"index"`

	// budget: เก็บเป็นหน่วยย่อยเพื่อลดปัญหา floating-point
	BudgetLimitCents int64 `gorm:"default:0"`

	IsActive bool `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// relations (optional)
	// Expenses []Expense `gorm:"foreignKey:CategoryID" json:"-"`
}
