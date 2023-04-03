package sensors

import (
	"os"
	"strconv"

	"github.com/shirou/gopsutil/host"
	"github.wdf.sap.corp/I554249/sensor/logger"
	systeminfo "github.wdf.sap.corp/I554249/sensor/system_info"
)

var systemSensorsInfo = host.SensorsTemperatures
var runWMICProcessElevatedMethod func() error
var Getwd = os.Getwd
var parseFloat = strconv.ParseFloat
var measurement systeminfo.Measurement
var infoLogger = logger.GetLogger().Info
var warningLogger = logger.GetLogger().Warning

type SystemSensor interface {
	StartMeasurement(string) (systeminfo.Measurement, error)
	GetInstanceId() string
}

type BaseSensor struct {
	Id           string
	DeviceId     string
	Unit         string
	SensorGroups string `default:"CPU_USAGE"`
	Measurements []systeminfo.Measurement
}

func (c *BaseSensor) GetInstanceId() string {
	return c.Id
}
