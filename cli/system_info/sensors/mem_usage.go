package sensors

import (
	"errors"

	systeminfo "github.com/altnum/sensorapp/system_info"
	"github.com/pbnjay/memory"
)

type MemUsageSensor struct {
	BaseSensor
	UsedPercent float64
}

func (m *MemUsageSensor) StartMeasurement(unitF string) (systeminfo.Measurement, error) {
	infoLogger.Println("Measuring the percentage of memory being used")

	err := m.setBusyPercentageOfMem()
	if err != nil {
		return systeminfo.Measurement{}, err
	}

	measurement.SetTimeStamp()
	measurement.Value = float64(m.UsedPercent)
	measurement.SensorId = m.Id
	measurement.DeviceId = m.DeviceId

	return measurement, nil
}

func (m *MemUsageSensor) setBusyPercentageOfMem() error {
	totalMemory := memory.TotalMemory()
	if totalMemory <= 0 {
		return errors.New("invalid usage memory value")
	}

	t := (float64((totalMemory - memory.FreeMemory())) / float64(totalMemory)) * float64(100)
	m.UsedPercent = float64(t)

	return nil
}
