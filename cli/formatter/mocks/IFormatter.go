// Code generated by mockery v2.9.3. DO NOT EDIT.

package mocks

import (
	cobra "github.com/spf13/cobra"

	mock "github.com/stretchr/testify/mock"

	systeminfo "github.com/altnum/sensorapp/system_info"
)

// IFormatter is an autogenerated mock type for the IFormatter type
type IFormatter struct {
	mock.Mock
}

// FormatInfoOutput provides a mock function with given fields: _a0, _a1
func (_m *IFormatter) FormatInfoOutput(_a0 string, _a1 systeminfo.Measurement) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, systeminfo.Measurement) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, systeminfo.Measurement) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FormatOutput provides a mock function with given fields: _a0, _a1
func (_m *IFormatter) FormatOutput(_a0 *cobra.Command, _a1 systeminfo.Measurement) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(*cobra.Command, systeminfo.Measurement) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cobra.Command, systeminfo.Measurement) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
