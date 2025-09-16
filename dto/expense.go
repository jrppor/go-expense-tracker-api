package dto

type ExpenseCreateRequest struct {
	CategoryID    *uint   `json:"category_id" binding:"required"`
	Amount        float64 `json:"amount" binding:"required"`
	Description   string  `json:"description,omitempty"`
	PaymentMethod string  `json:"payment_method,omitempty"`
	Tags          string  `json:"tags,omitempty"`
	// keep as string to support simple "YYYY-MM-DD" input via JSON
	Date string `json:"date" binding:"required"`
}

type ExpenseUpdateRequest struct {
	CategoryID    *uint   `json:"category_id,omitempty"`
	Amount        float64 `json:"amount"`
	Description   string  `json:"description"`
	PaymentMethod string  `json:"payment_method"`
	Tags          string  `json:"tags"`
}

type ExpenseResponse struct {
	CategoryID    *uint   `json:"category_id,omitempty"`
	Amount        float64 `json:"amount"`
	Description   string  `json:"description"`
	PaymentMethod string  `json:"payment_method"`
	Tags          string  `json:"tags"`
}
