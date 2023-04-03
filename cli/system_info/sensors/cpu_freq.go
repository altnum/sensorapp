package sensors

import (
	"fmt"
	"time"

	"github.com/dterei/gotsc"

	systeminfo "github.com/altnum/sensorapp/system_info"
)

type CpuFrequencySensor struct {
	BaseSensor
	Frequency float64
}

func (c *CpuFrequencySensor) StartMeasurement(unitF string) (systeminfo.Measurement, error) {
	infoLogger.Println("Starting cores frequency measurements")

	err := c.SetFrequency()
	if err != nil {
		return measurement, err
	}

	measurement.SetTimeStamp()
	measurement.Value = float64(c.Frequency)
	measurement.SensorId = c.Id
	measurement.DeviceId = c.DeviceId

	return measurement, nil
}

func (c *CpuFrequencySensor) SetFrequency() error {
	infoLogger.Println("Setting the frequency")

	var err error
	tsc := gotsc.TSCOverhead()

	start := gotsc.BenchStart()
	time.Sleep(1 * time.Second)
	end := gotsc.BenchEnd()

	cycles := (end - start - tsc)

	var frequencyHz float64

	frequencyHz = float64(cycles)

	if c.Unit == "GHz" {
		frequencyHz /= 1000000000
	}
	if c.Unit == "MHz" {
		frequencyHz /= 1000000
	}

	c.Frequency, err = parseFloat(fmt.Sprintf("%.0f", frequencyHz), 64)
	if err != nil {
		return err
	}

	return nil
}
