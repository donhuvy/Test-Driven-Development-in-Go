// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	db "github.com/PacktPublishing/Test-Driven-Development-in-Go/chapter07/db"
	mock "github.com/stretchr/testify/mock"
)

// BookOperationsService is an autogenerated mock type for the BookOperationsService type
type BookOperationsService struct {
	mock.Mock
}

// ListByUser provides a mock function with given fields: userID
func (_m *BookOperationsService) ListByUser(userID string) ([]db.Book, error) {
	ret := _m.Called(userID)

	var r0 []db.Book
	if rf, ok := ret.Get(0).(func(string) []db.Book); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBookOperationsService interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookOperationsService creates a new instance of BookOperationsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookOperationsService(t mockConstructorTestingTNewBookOperationsService) *BookOperationsService {
	mock := &BookOperationsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
