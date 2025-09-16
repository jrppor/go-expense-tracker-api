package repositories

import (
	"context"
	"jrppor/go-expense-tracker-api/models"

	"gorm.io/gorm"
)

type ExpenseRepository interface {
	List(ctx context.Context) ([]models.Expense, error)
	Create(ctx context.Context, e *models.Expense) error
	GetByID(ctx context.Context, id uint) (*models.Expense, error)
	Update(ctx context.Context, e *models.Expense) error
	Delete(ctx context.Context, id uint) error
}

type expenseRepository struct{ db *gorm.DB }

func NewExpenseRepository(db *gorm.DB) ExpenseRepository { return &expenseRepository{db: db} }

func (r *expenseRepository) List(ctx context.Context) ([]models.Expense, error) {
	var list []models.Expense
	if err := r.db.WithContext(ctx).Order("date DESC, id DESC").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *expenseRepository) Create(ctx context.Context, e *models.Expense) error {
	return r.db.WithContext(ctx).Create(e).Error
}

func (r *expenseRepository) GetByID(ctx context.Context, id uint) (*models.Expense, error) {
	var e models.Expense
	if err := r.db.WithContext(ctx).First(&e, id).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *expenseRepository) Update(ctx context.Context, e *models.Expense) error {
	return r.db.WithContext(ctx).Save(e).Error
}

func (r *expenseRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Expense{}, id).Error
}
