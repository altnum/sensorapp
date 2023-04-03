package sensors

import (
	"errors"
	"runtime"

	systeminfo "github.wdf.sap.corp/I554249/sensor/system_info"
)

var getCoreNum = runtime.NumCPU

type CpuCoresCounter struct {
	BaseSensor
	NumberOfCores int
}

func (c *CpuCoresCounter) StartMeasurement(unitF string) (systeminfo.Measurement, error) {
	infoLogger.Println("Starting cores counter measurements")

	c.NumberOfCores = getCoreNum()
	if c.NumberOfCores <= 0 {
		return measurement, errors.New("invalid number of cores")
	}

	measurement.SetTimeStamp()
	measurement.Value = float64(c.NumberOfCores)
	measurement.SensorId = c.Id
	measurement.DeviceId = c.DeviceId

	return measurement, nil
}
