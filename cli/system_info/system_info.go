package systeminfo

import (
	"fmt"
	"time"
)

type IMeasurement interface {
	SetTimeStamp()
}

type Measurement struct {
	MeasuredAt string  `json:"measuredAt" yaml:"measuredAt"`
	Value      float64 `json:"value" yaml:"value"`
	SensorId   string  `json:"sensorId" yaml:"sensorId"`
	DeviceId   string  `json:"deviceId" yaml:"deviceId"`
}

//Sets timestamp as string for the moment when measurements were taken.
func (m *Measurement) SetTimeStamp() {
	time := time.Now()

	m.MeasuredAt = fmt.Sprint(time.Format("2006-01-02 15:04:05"))
}
