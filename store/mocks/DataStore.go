// Copyright 2021 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/mendersoftware/iot-manager/model"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// DataStore is an autogenerated mock type for the DataStore type
type DataStore struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *DataStore) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateIntegration provides a mock function with given fields: _a0, _a1
func (_m *DataStore) CreateIntegration(_a0 context.Context, _a1 model.Integration) (*model.Integration, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.Integration
	if rf, ok := ret.Get(0).(func(context.Context, model.Integration) *model.Integration); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Integration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Integration) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDevice provides a mock function with given fields: ctx, deviceID
func (_m *DataStore) DeleteDevice(ctx context.Context, deviceID string) error {
	ret := _m.Called(ctx, deviceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, deviceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DoDevicesExistByIntegrationID provides a mock function with given fields: _a0, _a1
func (_m *DataStore) DoDevicesExistByIntegrationID(_a0 context.Context, _a1 uuid.UUID) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevice provides a mock function with given fields: ctx, deviceID
func (_m *DataStore) GetDevice(ctx context.Context, deviceID string) (*model.Device, error) {
	ret := _m.Called(ctx, deviceID)

	var r0 *model.Device
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Device); ok {
		r0 = rf(ctx, deviceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, deviceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceByIntegrationID provides a mock function with given fields: ctx, deviceID, integrationID
func (_m *DataStore) GetDeviceByIntegrationID(ctx context.Context, deviceID string, integrationID uuid.UUID) (*model.Device, error) {
	ret := _m.Called(ctx, deviceID, integrationID)

	var r0 *model.Device
	if rf, ok := ret.Get(0).(func(context.Context, string, uuid.UUID) *model.Device); ok {
		r0 = rf(ctx, deviceID, integrationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, uuid.UUID) error); ok {
		r1 = rf(ctx, deviceID, integrationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIntegrationById provides a mock function with given fields: _a0, _a1
func (_m *DataStore) GetIntegrationById(_a0 context.Context, _a1 uuid.UUID) (*model.Integration, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.Integration
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *model.Integration); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Integration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIntegrations provides a mock function with given fields: _a0, _a1
func (_m *DataStore) GetIntegrations(_a0 context.Context, _a1 model.IntegrationFilter) ([]model.Integration, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []model.Integration
	if rf, ok := ret.Get(0).(func(context.Context, model.IntegrationFilter) []model.Integration); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Integration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.IntegrationFilter) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ping provides a mock function with given fields: ctx
func (_m *DataStore) Ping(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveIntegration provides a mock function with given fields: _a0, _a1
func (_m *DataStore) RemoveIntegration(_a0 context.Context, _a1 uuid.UUID) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetIntegrationCredentials provides a mock function with given fields: _a0, _a1, _a2
func (_m *DataStore) SetIntegrationCredentials(_a0 context.Context, _a1 uuid.UUID, _a2 model.Credentials) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, model.Credentials) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertDeviceIntegrations provides a mock function with given fields: ctx, deviceID, integrationIDs
func (_m *DataStore) UpsertDeviceIntegrations(ctx context.Context, deviceID string, integrationIDs []uuid.UUID) (*model.Device, error) {
	ret := _m.Called(ctx, deviceID, integrationIDs)

	var r0 *model.Device
	if rf, ok := ret.Get(0).(func(context.Context, string, []uuid.UUID) *model.Device); ok {
		r0 = rf(ctx, deviceID, integrationIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []uuid.UUID) error); ok {
		r1 = rf(ctx, deviceID, integrationIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
