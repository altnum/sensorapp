package sensors

import (
	"errors"

	"github.com/pbnjay/memory"
	systeminfo "github.wdf.sap.corp/I554249/sensor/system_info"
)

type MemUsedSensor struct {
	BaseSensor
	Used uint64
}

func (m *MemUsedSensor) StartMeasurement(unitF string) (systeminfo.Measurement, error) {
	infoLogger.Println("Measuring the used memory")
	var err error

	m.Used, err = m.setUsedBytesOfMem()
	if err != nil {
		return systeminfo.Measurement{}, err
	}

	measurement.SetTimeStamp()
	measurement.Value = float64(m.Used)
	measurement.SensorId = m.Id
	measurement.DeviceId = m.DeviceId

	return measurement, nil
}

func (m *MemUsedSensor) setUsedBytesOfMem() (uint64, error) {
	memInBytes := memory.TotalMemory() - memory.FreeMemory()
	if memInBytes <= 0 {
		return 0, errors.New("invalid used memory value")
	}

	if m.Unit == "GigaBytes" {
		return memInBytes / 1000000000, nil
	}
	if m.Unit == "MegaBytes" {
		return memInBytes / 1000000, nil
	}
	if m.Unit == "KiloBytes" {
		return memInBytes / 1000, nil
	}

	return memInBytes, nil
}
