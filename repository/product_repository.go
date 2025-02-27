package repository

import (
	"context"
	"github.com/aronipurwanto/go-restful-api/model/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, product domain.Product) (domain.Product, error)
	Update(ctx context.Context, product domain.Product) (domain.Product, error)
	Delete(ctx context.Context, product domain.Product) error
	FindById(ctx context.Context, productId uint64) (domain.Product, error)
	FindAll(ctx context.Context) ([]domain.Product, error)
}
