// Code generated by mockery v1.0.0. DO NOT EDIT.

// This comment works around https://github.com/vektra/mockery/issues/155

package shells

import credentials "github.com/minio/minio-go/pkg/credentials"
import mock "github.com/stretchr/testify/mock"

// mockFakeIAMCredentialsProvider is an autogenerated mock type for the fakeIAMCredentialsProvider type
type mockFakeIAMCredentialsProvider struct {
	mock.Mock
}

// IsExpired provides a mock function with given fields:
func (_m *mockFakeIAMCredentialsProvider) IsExpired() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Retrieve provides a mock function with given fields:
func (_m *mockFakeIAMCredentialsProvider) Retrieve() (credentials.Value, error) {
	ret := _m.Called()

	var r0 credentials.Value
	if rf, ok := ret.Get(0).(func() credentials.Value); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(credentials.Value)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}