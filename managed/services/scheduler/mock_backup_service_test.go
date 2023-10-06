// Code generated by mockery v2.35.1. DO NOT EDIT.

package scheduler

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	backup "github.com/percona/pmm/managed/services/backup"
)

// mockBackupService is an autogenerated mock type for the backupService type
type mockBackupService struct {
	mock.Mock
}

// PerformBackup provides a mock function with given fields: ctx, params
func (_m *mockBackupService) PerformBackup(ctx context.Context, params backup.PerformBackupParams) (string, error) {
	ret := _m.Called(ctx, params)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, backup.PerformBackupParams) (string, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, backup.PerformBackupParams) string); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, backup.PerformBackupParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newMockBackupService creates a new instance of mockBackupService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockBackupService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockBackupService {
	mock := &mockBackupService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
