package dto

type CategoryCreateRequest struct {
	Name              string  `json:"name" binding:"required,min=1,max=64"`
	Slug              string  `json:"slug,omitempty"`
	Description       string  `json:"description,omitempty"`
	Color             string  `json:"color,omitempty"` // "#RRGGBB"
	Icon              string  `json:"icon,omitempty"`
	ParentID          *uint   `json:"parent_id,omitempty"`
	BudgetLimitAmount float64 `json:"budget_limit_amount,omitempty"` // รับเป็นบาท/หน่วยหลักจาก client
	Currency          string  `json:"currency,omitempty"`            // ถ้าต้องการรองรับหลายสกุล
	IsActive          *bool   `json:"is_active,omitempty"`
	// UserID optional ถ้ามีระบบผู้ใช้ในชั้น service
}

type CategoryUpdateRequest struct {
	Name              *string  `json:"name,omitempty"`
	Slug              *string  `json:"slug,omitempty"`
	Description       *string  `json:"description,omitempty"`
	Color             *string  `json:"color,omitempty"`
	Icon              *string  `json:"icon,omitempty"`
	ParentID          *uint    `json:"parent_id,omitempty"`
	BudgetLimitAmount *float64 `json:"budget_limit_amount,omitempty"`
	Currency          *string  `json:"currency,omitempty"`
	IsActive          *bool    `json:"is_active,omitempty"`
}

type CategoryResponse struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	Slug               string `json:"slug"`
	Description        string `json:"description,omitempty"`
	Color              string `json:"color,omitempty"`
	Icon               string `json:"icon,omitempty"`
	ParentID           *uint  `json:"parent_id,omitempty"`
	BudgetLimitCents   int64  `json:"budget_limit_cents"`
	BudgetLimitDisplay string `json:"budget_limit_display,omitempty"` // คำนวณแสดงผล เช่น "฿1,000.00"
	IsActive           bool   `json:"is_active"`
}
