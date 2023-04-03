// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	database_instances "github.wdf.sap.corp/I554249/sensor/databases/database_instances"

	models "github.wdf.sap.corp/I554249/sensor/models"
)

// IPostgreDB is an autogenerated mock type for the IPostgreDB type
type IPostgreDB struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *IPostgreDB) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateDevice provides a mock function with given fields: _a0
func (_m *IPostgreDB) CreateDevice(_a0 map[string]string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateSensor provides a mock function with given fields: _a0
func (_m *IPostgreDB) CreateSensor(_a0 map[string]string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDevice provides a mock function with given fields: _a0
func (_m *IPostgreDB) DeleteDevice(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSensor provides a mock function with given fields: _a0
func (_m *IPostgreDB) DeleteSensor(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllDevices provides a mock function with given fields:
func (_m *IPostgreDB) GetAllDevices() ([]models.Device, error) {
	ret := _m.Called()

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func() []models.Device); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllSensors provides a mock function with given fields:
func (_m *IPostgreDB) GetAllSensors() ([]models.Sensors, error) {
	ret := _m.Called()

	var r0 []models.Sensors
	if rf, ok := ret.Get(0).(func() []models.Sensors); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Sensors)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevice provides a mock function with given fields: _a0
func (_m *IPostgreDB) GetDevice(_a0 string) (models.Device, error) {
	ret := _m.Called(_a0)

	var r0 models.Device
	if rf, ok := ret.Get(0).(func(string) models.Device); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.Device)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInfluxInstance provides a mock function with given fields:
func (_m *IPostgreDB) GetInfluxInstance() database_instances.IInfluxDB {
	ret := _m.Called()

	var r0 database_instances.IInfluxDB
	if rf, ok := ret.Get(0).(func() database_instances.IInfluxDB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database_instances.IInfluxDB)
		}
	}

	return r0
}

// GetPostgreInstance provides a mock function with given fields:
func (_m *IPostgreDB) GetPostgreInstance() database_instances.IPostgreDB {
	ret := _m.Called()

	var r0 database_instances.IPostgreDB
	if rf, ok := ret.Get(0).(func() database_instances.IPostgreDB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database_instances.IPostgreDB)
		}
	}

	return r0
}

// GetSensor provides a mock function with given fields: _a0
func (_m *IPostgreDB) GetSensor(_a0 string) (models.Sensors, error) {
	ret := _m.Called(_a0)

	var r0 models.Sensors
	if rf, ok := ret.Get(0).(func(string) models.Sensors); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.Sensors)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Open provides a mock function with given fields: _a0
func (_m *IPostgreDB) Open(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDevice provides a mock function with given fields: _a0
func (_m *IPostgreDB) UpdateDevice(_a0 map[string]string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSensor provides a mock function with given fields: _a0
func (_m *IPostgreDB) UpdateSensor(_a0 map[string]string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
