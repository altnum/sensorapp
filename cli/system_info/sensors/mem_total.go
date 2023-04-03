package sensors

import (
	"errors"

	"github.com/pbnjay/memory"
	systeminfo "github.wdf.sap.corp/I554249/sensor/system_info"
)

type MemTotalSensor struct {
	BaseSensor
	Total uint64
}

func (m *MemTotalSensor) StartMeasurement(unitF string) (systeminfo.Measurement, error) {
	infoLogger.Println("Measuring the total memory")
	var err error

	m.Total, err = m.setTotalBytesOfMem()
	if err != nil {
		return systeminfo.Measurement{}, err
	}

	measurement.SetTimeStamp()
	measurement.Value = float64(m.Total)
	measurement.SensorId = m.Id
	measurement.DeviceId = m.DeviceId

	return measurement, nil
}

func (m *MemTotalSensor) setTotalBytesOfMem() (uint64, error) {
	memInBytes := memory.TotalMemory()
	if memInBytes <= 0 {
		return 0, errors.New("invalid total memory value ")
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
