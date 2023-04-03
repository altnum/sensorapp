package sensors

import (
	"fmt"
	"time"

	systeminfo "github.com/altnum/sensorapp/system_info"
	"github.com/shirou/gopsutil/cpu"
)

type CpuUsageSensor struct {
	BaseSensor
	UsedPercent float64
}

func (c *CpuUsageSensor) StartMeasurement(unitF string) (systeminfo.Measurement, error) {
	infoLogger.Println("Starting CPU usage measurements")

	err := c.SetUsagePercentage()
	if err != nil {
		return measurement, err
	}

	measurement.SetTimeStamp()
	measurement.Value = float64(c.UsedPercent)
	measurement.SensorId = c.Id
	measurement.DeviceId = c.DeviceId

	return measurement, nil
}

func (c *CpuUsageSensor) SetUsagePercentage() error {
	infoLogger.Println("Calculating usage percentage for the current processes")

	_, err := cpu.Percent(time.Second, false)
	if err != nil {
		return err
	}

	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return err
	}

	percentage := fmt.Sprintf("%.0f", percent[0])

	c.UsedPercent, err = parseFloat(percentage, 64)
	if err != nil {
		return err
	}

	return nil
}
