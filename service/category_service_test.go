package service

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/repository/mocks"
	"github.com/go-playground/validator/v10"
	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategoryRepository(ctrl)
	mockValidator := validator.New()
	categoryService := NewCategoryService(mockRepo, mockValidator)

	tests := []struct {
		name      string
		input     web.CategoryCreateRequest
		mock      func()
		expect    web.CategoryResponse
		expectErr bool
	}{
		{
			name:  "success",
			input: web.CategoryCreateRequest{Name: "Electronics"},
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Category{Id: 1, Name: "Electronics"}, nil)
			},
			expect:    web.CategoryResponse{Id: 1, Name: "Electronics"},
			expectErr: false,
		},
		{
			name:      "validation error",
			input:     web.CategoryCreateRequest{Name: ""},
			mock:      func() {},
			expect:    web.CategoryResponse{},
			expectErr: true,
		},
		{
			name:  "repository error",
			input: web.CategoryCreateRequest{Name: "Toys"},
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Category{}, errors.New("database error"))
			},
			expect:    web.CategoryResponse{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := categoryService.Create(context.Background(), tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}

func TestDeleteCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategoryRepository(ctrl)
	categoryService := NewCategoryService(mockRepo, validator.New())

	tests := []struct {
		name       string
		categoryId uint64
		mock       func()
		expectErr  bool
	}{
		{
			name:       "success",
			categoryId: 1,
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), uint64(1)).Return(domain.Category{Id: 1, Name: "Electronics"}, nil)
				mockRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectErr: false,
		},
		{
			name:       "not found",
			categoryId: 99,
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), uint64(99)).Return(domain.Category{}, errors.New("not found"))
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := categoryService.Delete(context.Background(), tt.categoryId)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateCategory(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mockCategoryRepo *mocks.MockCategoryRepository)
		input   web.CategoryUpdateRequest
		expects error
	}{
		{
			name: "Success",
			mock: func(mockCategoryRepo *mocks.MockCategoryRepository) {
				mockCategoryRepo.EXPECT().FindById(gomock.Any(), uint64(1)).
					Return(domain.Category{Id: 1, Name: "Old Name"}, nil)
				mockCategoryRepo.EXPECT().Update(gomock.Any(), gomock.Any()).
					Return(domain.Category{Id: 1, Name: "New Name"}, nil)
			},
			input:   web.CategoryUpdateRequest{Id: 1, Name: "New Name"},
			expects: nil,
		},
		{
			name: "Category Not Found",
			mock: func(mockCategoryRepo *mocks.MockCategoryRepository) {
				mockCategoryRepo.EXPECT().FindById(gomock.Any(), uint64(1)).
					Return(domain.Category{}, errors.New("not found"))
			},
			input:   web.CategoryUpdateRequest{Id: 1, Name: "New Name"},
			expects: errors.New("not found"),
		},
		{
			name: "Validation Error - Empty Name",
			mock: func(mockCategoryRepo *mocks.MockCategoryRepository) {
				// Tidak perlu mock FindById karena validasi gagal sebelum ke repository
			},
			input:   web.CategoryUpdateRequest{Id: 1, Name: ""},
			expects: errors.New("CategoryUpdateRequest.Name"),
		},
		{
			name: "Database Error on Update",
			mock: func(mockCategoryRepo *mocks.MockCategoryRepository) {
				mockCategoryRepo.EXPECT().FindById(gomock.Any(), uint64(1)).
					Return(domain.Category{Id: 1, Name: "Old Name"}, nil)
				mockCategoryRepo.EXPECT().Update(gomock.Any(), gomock.Any()).
					Return(domain.Category{}, errors.New("database error"))
			},
			input:   web.CategoryUpdateRequest{Id: 1, Name: "Updated Name"},
			expects: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockCategoryRepo := mocks.NewMockCategoryRepository(ctrl)
			tt.mock(mockCategoryRepo)

			service := NewCategoryService(mockCategoryRepo, validator.New())
			_, err := service.Update(context.Background(), tt.input)

			if tt.expects != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expects.Error()) // Alternatif untuk assert.ErrorContains
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestFindAllCategories(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mockCategoryRepo *mocks.MockCategoryRepository)
		expects []web.CategoryResponse
		err     error
	}{
		{
			name: "Success",
			mock: func(mockCategoryRepo *mocks.MockCategoryRepository) {
				mockCategoryRepo.EXPECT().FindAll(gomock.Any()).Return([]domain.Category{{Id: 1, Name: "Category 1"}}, nil)
			},
			expects: []web.CategoryResponse{{Id: 1, Name: "Category 1"}},
			err:     nil,
		},
		{
			name: "Database Error",
			mock: func(mockCategoryRepo *mocks.MockCategoryRepository) {
				mockCategoryRepo.EXPECT().FindAll(gomock.Any()).Return(nil, errors.New("database error"))
			},
			expects: nil,
			err:     errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockCategoryRepo := mocks.NewMockCategoryRepository(ctrl)
			tt.mock(mockCategoryRepo)

			service := NewCategoryService(mockCategoryRepo, validator.New())
			result, err := service.FindAll(context.Background())
			assert.Equal(t, tt.expects, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestFindByIdCategory(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mockCategoryRepo *mocks.MockCategoryRepository)
		input   uint64
		expects web.CategoryResponse
		err     error
	}{
		{
			name: "Success",
			mock: func(mockCategoryRepo *mocks.MockCategoryRepository) {
				mockCategoryRepo.EXPECT().FindById(gomock.Any(), uint64(1)).Return(domain.Category{Id: 1, Name: "Category 1"}, nil)
			},
			input:   1,
			expects: web.CategoryResponse{Id: 1, Name: "Category 1"},
			err:     nil,
		},
		{
			name: "Not Found",
			mock: func(mockCategoryRepo *mocks.MockCategoryRepository) {
				mockCategoryRepo.EXPECT().FindById(gomock.Any(), uint64(1)).Return(domain.Category{}, errors.New("not found"))
			},
			input:   1,
			expects: web.CategoryResponse{},
			err:     errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockCategoryRepo := mocks.NewMockCategoryRepository(ctrl)
			tt.mock(mockCategoryRepo)

			service := NewCategoryService(mockCategoryRepo, validator.New())
			result, err := service.FindById(context.Background(), tt.input)
			assert.Equal(t, tt.expects, result)
			assert.Equal(t, tt.err, err)
		})
	}
}
