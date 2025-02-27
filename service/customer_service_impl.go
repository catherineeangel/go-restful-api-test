package service

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/exception"
	"github.com/aronipurwanto/go-restful-api/helper"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	Validate           *validator.Validate
}

func NewCustomerService(customerRepository repository.CustomerRepository, validate *validator.Validate) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: customerRepository,
		Validate:           validate,
	}
}

// Create Customer
func (service *CustomerServiceImpl) Create(ctx context.Context, request web.CustomerCreateRequest) (web.CustomerResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return web.CustomerResponse{}, err
	}

	customer := domain.Customer{Name: request.Name}
	savedCustomer, err := service.CustomerRepository.Save(ctx, customer)
	if err != nil {
		return web.CustomerResponse{}, err
	}

	return helper.ToCustomerResponse(savedCustomer), nil
}

// Update Customer
func (service *CustomerServiceImpl) Update(ctx context.Context, request web.CustomerUpdateRequest) (web.CustomerResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return web.CustomerResponse{}, err
	}

	customer, err := service.CustomerRepository.FindById(ctx, request.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return web.CustomerResponse{}, exception.NewNotFoundError("Customer not found")
	} else if err != nil {
		return web.CustomerResponse{}, err
	}

	customer.Name = request.Name
	updatedCustomer, err := service.CustomerRepository.Update(ctx, customer)
	if err != nil {
		return web.CustomerResponse{}, err
	}

	return helper.ToCustomerResponse(updatedCustomer), nil
}

// Delete Customer
func (service *CustomerServiceImpl) Delete(ctx context.Context, customerId uint64) error {
	customer, err := service.CustomerRepository.FindById(ctx, customerId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return exception.NewNotFoundError("Customer not found")
	} else if err != nil {
		return err
	}

	return service.CustomerRepository.Delete(ctx, customer)
}

// Find Customer By ID
func (service *CustomerServiceImpl) FindById(ctx context.Context, customerId uint64) (web.CustomerResponse, error) {
	customer, err := service.CustomerRepository.FindById(ctx, customerId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return web.CustomerResponse{}, exception.NewNotFoundError("Customer not found")
	} else if err != nil {
		return web.CustomerResponse{}, err
	}

	return helper.ToCustomerResponse(customer), nil
}

// Find All Customer
func (service *CustomerServiceImpl) FindAll(ctx context.Context) ([]web.CustomerResponse, error) {
	customers, err := service.CustomerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return helper.ToCustomerResponses(customers), nil
}
