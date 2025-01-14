// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	db "github.com/donhuvy/Test-Driven-Development-in-Go/chapter11/db"
	mock "github.com/stretchr/testify/mock"
)

// MagazineOperationsService is an autogenerated mock type for the MagazineOperationsService type
type MagazineOperationsService struct {
	mock.Mock
}

// ListByUser provides a mock function with given fields: userID
func (_m *MagazineOperationsService) ListByUser(userID string) ([]db.Magazine, error) {
	ret := _m.Called(userID)

	var r0 []db.Magazine
	if rf, ok := ret.Get(0).(func(string) []db.Magazine); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Magazine)
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

type mockConstructorTestingTNewMagazineOperationsService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMagazineOperationsService creates a new instance of MagazineOperationsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMagazineOperationsService(t mockConstructorTestingTNewMagazineOperationsService) *MagazineOperationsService {
	mock := &MagazineOperationsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
