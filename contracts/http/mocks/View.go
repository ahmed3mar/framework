// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// View is an autogenerated mock type for the View type
type View struct {
	mock.Mock
}

// Exists provides a mock function with given fields: view
func (_m *View) Exists(view string) bool {
	ret := _m.Called(view)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(view)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetShared provides a mock function with given fields:
func (_m *View) GetShared() map[string]interface{} {
	ret := _m.Called()

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func() map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// Share provides a mock function with given fields: key, value
func (_m *View) Share(key string, value interface{}) {
	_m.Called(key, value)
}

// Shared provides a mock function with given fields: key, def
func (_m *View) Shared(key string, def ...interface{}) interface{} {
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, def...)
	ret := _m.Called(_ca...)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, ...interface{}) interface{}); ok {
		r0 = rf(key, def...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

type mockConstructorTestingTNewView interface {
	mock.TestingT
	Cleanup(func())
}

// NewView creates a new instance of View. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewView(t mockConstructorTestingTNewView) *View {
	mock := &View{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
