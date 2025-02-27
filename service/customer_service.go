package service

import (
	"context"
	"github.com/aronipurwanto/go-restful-api/model/web"
)

type CustomerService interface {
	Create(ctx context.Context, request web.CustomerCreateRequest) (web.CustomerResponse, error)
	Update(ctx context.Context, request web.CustomerUpdateRequest) (web.CustomerResponse, error)
	Delete(ctx context.Context, categoryId uint64) error
	FindById(ctx context.Context, categoryId uint64) (web.CustomerResponse, error)
	FindAll(ctx context.Context) ([]web.CustomerResponse, error)
}
