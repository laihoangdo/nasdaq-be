// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"
	models "nasdaqvfs/internal/models"

	mock "github.com/stretchr/testify/mock"

	utils "nasdaqvfs/pkg/utils"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *Repository) Create(ctx context.Context, _a1 models.User) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.User) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByUsername provides a mock function with given fields: ctx, username
func (_m *Repository) GetByUsername(ctx context.Context, username string) (models.User, error) {
	ret := _m.Called(ctx, username)

	var r0 models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (models.User, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) models.User); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: ctx, id
func (_m *Repository) GetUserByID(ctx context.Context, id int64) (models.User, error) {
	ret := _m.Called(ctx, id)

	var r0 models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (models.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) models.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: ctx, pq
func (_m *Repository) GetUsers(ctx context.Context, pq *utils.PaginationQuery) ([]models.User, error) {
	ret := _m.Called(ctx, pq)

	var r0 []models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *utils.PaginationQuery) ([]models.User, error)); ok {
		return rf(ctx, pq)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *utils.PaginationQuery) []models.User); ok {
		r0 = rf(ctx, pq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *utils.PaginationQuery) error); ok {
		r1 = rf(ctx, pq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserByID provides a mock function with given fields: ctx, id, _a2
func (_m *Repository) UpdateUserByID(ctx context.Context, id int64, _a2 models.User) error {
	ret := _m.Called(ctx, id, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, models.User) error); ok {
		r0 = rf(ctx, id, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
