package repository

import (
	"context"
	"github.com/aronipurwanto/go-restful-api/model/domain"
)

type DiscountRepository interface {
	Save(ctx context.Context, discount domain.Discount) (domain.Discount, error)
	Update(ctx context.Context, discount domain.Discount) (domain.Discount, error)
	Delete(ctx context.Context, discount domain.Discount) error
	FindById(ctx context.Context, discountId uint64) (domain.Discount, error)
	FindAll(ctx context.Context) ([]domain.Discount, error)
}
