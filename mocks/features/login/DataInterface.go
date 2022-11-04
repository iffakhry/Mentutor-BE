// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	login "be12/mentutor/features/login"

	mock "github.com/stretchr/testify/mock"
)

// DataInterface is an autogenerated mock type for the DataInterface type
type DataInterface struct {
	mock.Mock
}

// Login provides a mock function with given fields: input
func (_m *DataInterface) Login(input login.Core) (login.Core, error) {
	ret := _m.Called(input)

	var r0 login.Core
	if rf, ok := ret.Get(0).(func(login.Core) login.Core); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(login.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(login.Core) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDataInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewDataInterface creates a new instance of DataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDataInterface(t mockConstructorTestingTNewDataInterface) *DataInterface {
	mock := &DataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
