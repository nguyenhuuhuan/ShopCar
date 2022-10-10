// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	models "Improve/src/models"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// RoleRepository is an autogenerated mock type for the RoleRepository type
type RoleRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, role
func (_m *RoleRepository) Create(ctx context.Context, role *models.Role) error {
	ret := _m.Called(ctx, role)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Role) error); ok {
		r0 = rf(ctx, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *RoleRepository) GetByID(ctx context.Context, id int64) (*models.Role, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Role
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.Role); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsExistRoleAndCode provides a mock function with given fields: ctx, roleName, code
func (_m *RoleRepository) IsExistRoleAndCode(ctx context.Context, roleName string, code string) bool {
	ret := _m.Called(ctx, roleName, code)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, roleName, code)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewRoleRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRoleRepository creates a new instance of RoleRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRoleRepository(t mockConstructorTestingTNewRoleRepository) *RoleRepository {
	mock := &RoleRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
