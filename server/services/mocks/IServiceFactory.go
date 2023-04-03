// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	database_instances "github.wdf.sap.corp/I554249/sensor/databases/database_instances"

	services "github.wdf.sap.corp/I554249/sensor/services"
)

// IServiceFactory is an autogenerated mock type for the IServiceFactory type
type IServiceFactory struct {
	mock.Mock
}

// CreateService provides a mock function with given fields: _a0, _a1
func (_m *IServiceFactory) CreateService(_a0 string, _a1 []database_instances.IDB) services.IService {
	ret := _m.Called(_a0, _a1)

	var r0 services.IService
	if rf, ok := ret.Get(0).(func(string, []database_instances.IDB) services.IService); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(services.IService)
		}
	}

	return r0
}
