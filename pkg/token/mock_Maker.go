// Code generated by mockery v2.28.1. DO NOT EDIT.

package token

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockMaker is an autogenerated mock type for the Maker type
type MockMaker struct {
	mock.Mock
}

// CreateToken provides a mock function with given fields: payload
func (_m *MockMaker) CreateToken(payload Payload) (string, time.Time, error) {
	ret := _m.Called(payload)

	var r0 string
	var r1 time.Time
	var r2 error
	if rf, ok := ret.Get(0).(func(Payload) (string, time.Time, error)); ok {
		return rf(payload)
	}
	if rf, ok := ret.Get(0).(func(Payload) string); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(Payload) time.Time); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Get(1).(time.Time)
	}

	if rf, ok := ret.Get(2).(func(Payload) error); ok {
		r2 = rf(payload)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// VerifyToken provides a mock function with given fields: token
func (_m *MockMaker) VerifyToken(token string) (Payload, error) {
	ret := _m.Called(token)

	var r0 Payload
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (Payload, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) Payload); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(Payload)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockMaker interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockMaker creates a new instance of MockMaker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockMaker(t mockConstructorTestingTNewMockMaker) *MockMaker {
	mock := &MockMaker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
