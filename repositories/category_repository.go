package repositories

import (
	"context"
	"jrppor/go-expense-tracker-api/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	List(ctx context.Context) ([]models.Category, error)
	Create(ctx context.Context, e *models.Category) error
	GetByID(ctx context.Context, id uint) (*models.Category, error)
	Update(ctx context.Context, e *models.Category) error
	Delete(ctx context.Context, id uint) error
}

type categoryRepository struct{ db *gorm.DB }

func NewCategoryRepository(db *gorm.DB) CategoryRepository { return &categoryRepository{db: db} }

func (r *categoryRepository) List(ctx context.Context) ([]models.Category, error) {
	var list []models.Category

	if err := r.db.WithContext(ctx).Order("sort_order ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *categoryRepository) Create(ctx context.Context, e *models.Category) error {
	return r.db.WithContext(ctx).Create(e).Error
}

func (r *categoryRepository) GetByID(ctx context.Context, id uint) (*models.Category, error) {
	var e models.Category
	if err := r.db.WithContext(ctx).First(&e, id).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *categoryRepository) Update(ctx context.Context, e *models.Category) error {
	return r.db.WithContext(ctx).Save(e).Error
}

func (r *categoryRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Category{}, id).Error
}
