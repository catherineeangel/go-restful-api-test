package repository

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"gorm.io/gorm"
)

type DiscountRepositoryImpl struct {
	db *gorm.DB
}

func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	return &DiscountRepositoryImpl{db: db}
}

// Save discount
func (repository *DiscountRepositoryImpl) Save(ctx context.Context, discount domain.Discount) (domain.Discount, error) {
	if err := repository.db.WithContext(ctx).Create(&discount).Error; err != nil {
		return domain.Discount{}, err
	}
	return discount, nil
}

// Update discount
func (repository *DiscountRepositoryImpl) Update(ctx context.Context, discount domain.Discount) (domain.Discount, error) {
	if err := repository.db.WithContext(ctx).Save(&discount).Error; err != nil {
		return domain.Discount{}, err
	}
	return discount, nil
}

// Delete discount
func (repository *DiscountRepositoryImpl) Delete(ctx context.Context, discount domain.Discount) error {
	if err := repository.db.WithContext(ctx).Delete(&discount).Error; err != nil {
		return err
	}
	return nil
}

// FindById - Get discount by ID
func (repository *DiscountRepositoryImpl) FindById(ctx context.Context, discountId uint64) (domain.Discount, error) {
	var discount domain.Discount
	err := repository.db.WithContext(ctx).First(&discount, discountId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return discount, errors.New("discount is not found")
	}
	return discount, err
}

// FindAll - Get all discounts
func (repository *DiscountRepositoryImpl) FindAll(ctx context.Context) ([]domain.Discount, error) {
	var discounts []domain.Discount
	err := repository.db.WithContext(ctx).Find(&discounts).Error
	return discounts, err
}
