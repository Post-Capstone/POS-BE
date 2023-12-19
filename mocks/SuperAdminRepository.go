// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "qbills/models/domain"

	mock "github.com/stretchr/testify/mock"
)

// SuperAdminRepository is an autogenerated mock type for the SuperAdminRepository type
type SuperAdminRepository struct {
	mock.Mock
}

// FindAll provides a mock function with given fields:
func (_m *SuperAdminRepository) FindAll() ([]domain.SuperAdmin, error) {
	ret := _m.Called()

	var r0 []domain.SuperAdmin
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.SuperAdmin, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.SuperAdmin); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.SuperAdmin)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: id
func (_m *SuperAdminRepository) FindById(id int) (*domain.SuperAdmin, error) {
	ret := _m.Called(id)

	var r0 *domain.SuperAdmin
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*domain.SuperAdmin, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *domain.SuperAdmin); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.SuperAdmin)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUsername provides a mock function with given fields: username
func (_m *SuperAdminRepository) FindByUsername(username string) (*domain.SuperAdmin, error) {
	ret := _m.Called(username)

	var r0 *domain.SuperAdmin
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.SuperAdmin, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.SuperAdmin); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.SuperAdmin)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSuperAdminRepository creates a new instance of SuperAdminRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSuperAdminRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *SuperAdminRepository {
	mock := &SuperAdminRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
