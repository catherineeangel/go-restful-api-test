package repository

import (
	"context"
	"github.com/aronipurwanto/go-restful-api/model/domain"
)

type CustomerRepository interface {
	Save(ctx context.Context, customer domain.Customer) (domain.Customer, error)
	Update(ctx context.Context, customer domain.Customer) (domain.Customer, error)
	Delete(ctx context.Context, customer domain.Customer) error
	FindById(ctx context.Context, customerId uint64) (domain.Customer, error)
	FindAll(ctx context.Context) ([]domain.Customer, error)
}
