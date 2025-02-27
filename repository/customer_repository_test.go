package repository

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomerRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockCustomerRepository(ctrl)
	ctx := context.Background()
	testCustomer := domain.Customer{CustomerID: 1, Name: "Andi Noya", Email: "test@test.com", Phone: "123456", Address: "Test address"}
	updatedCustomer := domain.Customer{CustomerID: 1, Name: "Updated Name", Email: "updated@test.com", Phone: "12345678", Address: "Updated Test address"}
	tests := []struct {
		name      string
		mock      func()
		method    func() (interface{}, error)
		expect    interface{}
		expectErr bool
	}{
		{
			name: "Save Success",
			mock: func() {
				repo.EXPECT().Save(ctx, testCustomer).Return(testCustomer, nil)
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, testCustomer)
			},
			expect:    testCustomer,
			expectErr: false,
		},
		{
			name: "Save Failure",
			mock: func() {
				repo.EXPECT().Save(ctx, gomock.Any()).Return(domain.Customer{}, errors.New("error saving"))
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Customer{Name: "Invalid", Email: "Invalid", Phone: "Invalid", Address: "Invalid"})
			},
			expect:    domain.Customer{},
			expectErr: true,
		},
		{
			name: "Update Success",
			mock: func() {

				repo.EXPECT().Update(ctx, testCustomer).Return(updatedCustomer, nil)
			},
			method: func() (interface{}, error) {

				return repo.Update(ctx, testCustomer)
			},
			expect:    updatedCustomer,
			expectErr: false,
		},
		{
			name: "FindById Success",
			mock: func() {
				repo.EXPECT().FindById(ctx, uint64(1)).Return(testCustomer, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, 1)
			},
			expect:    testCustomer,
			expectErr: false,
		},
		{
			name: "FindById Not Found",
			mock: func() {
				repo.EXPECT().FindById(ctx, uint64(999)).Return(domain.Customer{}, errors.New("not found"))
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, 999)
			},
			expect:    domain.Customer{},
			expectErr: true,
		},
		{
			name: "FindAll Success",
			mock: func() {
				repo.EXPECT().FindAll(ctx).Return([]domain.Customer{testCustomer}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindAll(ctx)
			},
			expect:    []domain.Customer{testCustomer},
			expectErr: false,
		},
		{
			name: "Delete Success",
			mock: func() {
				repo.EXPECT().Delete(ctx, domain.Customer{CustomerID: 1}).Return(nil)
			},
			method: func() (interface{}, error) {
				return nil, repo.Delete(ctx, domain.Customer{CustomerID: 1})
			},
			expect:    nil,
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			result, err := tt.method()

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, result)
			}
		})
	}
}
