// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "qbills/models/domain"

	mock "github.com/stretchr/testify/mock"
)

// PaymentTypeRepository is an autogenerated mock type for the PaymentTypeRepository type
type PaymentTypeRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: paymentType
func (_m *PaymentTypeRepository) Create(paymentType *domain.PaymentType) (*domain.PaymentType, error) {
	ret := _m.Called(paymentType)

	var r0 *domain.PaymentType
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.PaymentType) (*domain.PaymentType, error)); ok {
		return rf(paymentType)
	}
	if rf, ok := ret.Get(0).(func(*domain.PaymentType) *domain.PaymentType); ok {
		r0 = rf(paymentType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.PaymentType)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.PaymentType) error); ok {
		r1 = rf(paymentType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *PaymentTypeRepository) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields:
func (_m *PaymentTypeRepository) FindAll() ([]domain.PaymentType, error) {
	ret := _m.Called()

	var r0 []domain.PaymentType
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.PaymentType, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.PaymentType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.PaymentType)
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
func (_m *PaymentTypeRepository) FindById(id int) (*domain.PaymentType, error) {
	ret := _m.Called(id)

	var r0 *domain.PaymentType
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*domain.PaymentType, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *domain.PaymentType); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.PaymentType)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByName provides a mock function with given fields: name
func (_m *PaymentTypeRepository) FindByName(name string) (*domain.PaymentType, error) {
	ret := _m.Called(name)

	var r0 *domain.PaymentType
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.PaymentType, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.PaymentType); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.PaymentType)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: paymentType, id
func (_m *PaymentTypeRepository) Update(paymentType *domain.PaymentType, id int) (*domain.PaymentType, error) {
	ret := _m.Called(paymentType, id)

	var r0 *domain.PaymentType
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.PaymentType, int) (*domain.PaymentType, error)); ok {
		return rf(paymentType, id)
	}
	if rf, ok := ret.Get(0).(func(*domain.PaymentType, int) *domain.PaymentType); ok {
		r0 = rf(paymentType, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.PaymentType)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.PaymentType, int) error); ok {
		r1 = rf(paymentType, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPaymentTypeRepository creates a new instance of PaymentTypeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPaymentTypeRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PaymentTypeRepository {
	mock := &PaymentTypeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
