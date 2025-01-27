// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	calculator "github.com/donhuvy/Test-Driven-Development-in-Go/chapter03/calculator"

	mock "github.com/stretchr/testify/mock"
)

// OperationProcessor is an autogenerated mock type for the OperationProcessor type
type OperationProcessor struct {
	mock.Mock
}

// ProcessOperation provides a mock function with given fields: operation
func (_m *OperationProcessor) ProcessOperation(operation calculator.Operation) (*string, error) {
	ret := _m.Called(operation)

	var r0 *string
	if rf, ok := ret.Get(0).(func(calculator.Operation) *string); ok {
		r0 = rf(operation)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(calculator.Operation) error); ok {
		r1 = rf(operation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewOperationProcessor interface {
	mock.TestingT
	Cleanup(func())
}

// NewOperationProcessor creates a new instance of OperationProcessor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOperationProcessor(t mockConstructorTestingTNewOperationProcessor) *OperationProcessor {
	mock := &OperationProcessor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
