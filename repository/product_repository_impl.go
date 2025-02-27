package repository

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

// Save product
func (repository *ProductRepositoryImpl) Save(ctx context.Context, product domain.Product) (domain.Product, error) {
	if err := repository.db.WithContext(ctx).Create(&product).Error; err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

// Update product
func (repository *ProductRepositoryImpl) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	if err := repository.db.WithContext(ctx).Save(&product).Error; err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

// Delete product
func (repository *ProductRepositoryImpl) Delete(ctx context.Context, product domain.Product) error {
	if err := repository.db.WithContext(ctx).Delete(&product).Error; err != nil {
		return err
	}
	return nil
}

// FindById - Get product by ID
func (repository *ProductRepositoryImpl) FindById(ctx context.Context, productId uint64) (domain.Product, error) {
	var product domain.Product
	err := repository.db.WithContext(ctx).First(&product, productId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return product, errors.New("product is not found")
	}
	return product, err
}

// FindAll - Get all products
func (repository *ProductRepositoryImpl) FindAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	err := repository.db.WithContext(ctx).Find(&products).Error
	return products, err
}
