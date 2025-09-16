package services

import (
	"context"
	"errors"
	"jrppor/go-expense-tracker-api/dto"
	"jrppor/go-expense-tracker-api/models"
	"jrppor/go-expense-tracker-api/repositories"
	"time"
)

type CategoryService interface {
	List(ctx context.Context) ([]models.Category, error)
	Create(ctx context.Context, req dto.CategoryCreateRequest) (*models.Category, error)
	Update(ctx context.Context, id uint, req dto.CategoryUpdateRequest) (*models.Category, error)
	GetByID(ctx context.Context, id uint) (*models.Category, error)
	Delete(ctx context.Context, id uint) error
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) List(ctx context.Context) ([]models.Category, error) {
	return s.repo.List(ctx)
}

func (s *categoryService) Create(ctx context.Context, req dto.CategoryCreateRequest) (*models.Category, error) {

	e := &models.Category{
		Name:             req.Name,
		Slug:             req.Slug,
		Description:      req.Description,
		Color:            req.Color,
		Icon:             req.Icon,
		BudgetLimitCents: int64(req.BudgetLimitAmount),
		IsActive:         *req.IsActive,
	}

	if err := s.repo.Create(ctx, e); err != nil {
		return nil, err
	}

	return e, nil
}

func (s *categoryService) Update(ctx context.Context, id uint, req dto.CategoryUpdateRequest) (*models.Category, error) {
	e, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	e.UpdatedAt = time.Now()
	if err := s.repo.Update(ctx, e); err != nil {
		return nil, err
	}
	return e, nil
}

func (s *categoryService) GetByID(ctx context.Context, id uint) (*models.Category, error) {
	if id == 0 {
		return nil, errors.New("invalid id")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *categoryService) Delete(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid id")
	}
	return s.repo.Delete(ctx, id)
}
