package services

import (
	"context"
	"errors"
	"jrppor/go-expense-tracker-api/dto"
	"jrppor/go-expense-tracker-api/models"
	"jrppor/go-expense-tracker-api/repositories"
	"time"
)

type ExpenseService interface {
	List(ctx context.Context) ([]models.Expense, error)
	Create(ctx context.Context, req dto.ExpenseCreateRequest) (*models.Expense, error)
	Update(ctx context.Context, id uint, req dto.ExpenseUpdateRequest) (*models.Expense, error)
	GetByID(ctx context.Context, id uint) (*models.Expense, error)
	Delete(ctx context.Context, id uint) error
}

type expenseService struct {
	repo repositories.ExpenseRepository
}

func NewExpenseService(repo repositories.ExpenseRepository) ExpenseService {
	return &expenseService{repo: repo}
}

func (s *expenseService) List(ctx context.Context) ([]models.Expense, error) {
	return s.repo.List(ctx)
}

func (s *expenseService) Create(ctx context.Context, req dto.ExpenseCreateRequest) (*models.Expense, error) {
	// Parse date from string "YYYY-MM-DD"
	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, errors.New("invalid date format, expected YYYY-MM-DD")
	}
	// Validate: require category_id
	if req.CategoryID == nil || *req.CategoryID == 0 {
		return nil, errors.New("category_id is required")
	}
	e := &models.Expense{
		CategoryID:    req.CategoryID,
		Amount:        req.Amount,
		Description:   req.Description,
		PaymentMethod: req.PaymentMethod,
		Tags:          req.Tags,
		Date:          parsedDate,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	if err := s.repo.Create(ctx, e); err != nil {
		return nil, err
	}
	return e, nil
}

func (s *expenseService) Update(ctx context.Context, id uint, req dto.ExpenseUpdateRequest) (*models.Expense, error) {
	e, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if req.CategoryID != nil {
		e.CategoryID = req.CategoryID
	}
	if req.Amount != 0 {
		e.Amount = req.Amount
	}
	if req.Description != "" {
		e.Description = req.Description
	}
	if req.PaymentMethod != "" {
		e.PaymentMethod = req.PaymentMethod
	}
	if req.Tags != "" {
		e.Tags = req.Tags
	}
	e.UpdatedAt = time.Now()
	if err := s.repo.Update(ctx, e); err != nil {
		return nil, err
	}
	return e, nil
}

func (s *expenseService) GetByID(ctx context.Context, id uint) (*models.Expense, error) {
	if id == 0 {
		return nil, errors.New("invalid id")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *expenseService) Delete(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid id")
	}
	return s.repo.Delete(ctx, id)
}
