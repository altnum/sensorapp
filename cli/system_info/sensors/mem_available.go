package sensors

import (
	"github.com/pbnjay/memory"
	systeminfo "github.wdf.sap.corp/I554249/sensor/system_info"
)

type MemAvailableSensor struct {
	BaseSensor
	Available uint64
}

func (m *MemAvailableSensor) StartMeasurement(unitF string) (systeminfo.Measurement, error) {
	infoLogger.Println("Measuring the available memory")

	m.Available = m.setFreeBytesOfMem()

	measurement.SetTimeStamp()
	measurement.Value = float64(m.Available)
	measurement.SensorId = m.Id
	measurement.DeviceId = m.DeviceId

	return measurement, nil
}

func (m *MemAvailableSensor) setFreeBytesOfMem() uint64 {
	memInBytes := memory.FreeMemory()

	if m.Unit == "GigaBytes" {
		return memInBytes / 1000000000
	}
	if m.Unit == "MegaBytes" {
		return memInBytes / 1000000
	}
	if m.Unit == "KiloBytes" {
		return memInBytes / 1000
	}

	return memInBytes
}
