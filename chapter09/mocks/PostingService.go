// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	db "github.com/donhuvy/Test-Driven-Development-in-Go/chapter09/db"
	mock "github.com/stretchr/testify/mock"
)

// PostingService is an autogenerated mock type for the PostingService type
type PostingService struct {
	mock.Mock
}

// NewOrder provides a mock function with given fields: b
func (_m *PostingService) NewOrder(b db.Book) error {
	ret := _m.Called(b)

	var r0 error
	if rf, ok := ret.Get(0).(func(db.Book) error); ok {
		r0 = rf(b)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPostingService interface {
	mock.TestingT
	Cleanup(func())
}

// NewPostingService creates a new instance of PostingService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPostingService(t mockConstructorTestingTNewPostingService) *PostingService {
	mock := &PostingService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
