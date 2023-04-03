package sensorsfactory

import (
	"errors"
	"io/ioutil"
	"regexp"

	systeminfo "github.com/altnum/sensorapp/system_info"
	"github.com/altnum/sensorapp/system_info/sensors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var getFlagValue = getFlagValueString

const CPU_TEMP_SENSOR string = "cpuTempCelsius"
const CPU_CORES_COUNTER string = "cpuCoresCount"
const CPU_FREQ_SENSOR string = "cpuFrequency"
const CPU_USAGE_SENSOR string = "cpuUsagePercent"
const MEM_TOTAL_SENSOR string = "memoryTotal"
const MEM_USED_SENSOR string = "memoryUsedBytes"
const MEM_AVAILABLE_SENSOR string = "memoryAvailableBytes"
const MEM_USAGE_SENSOR string = "memoryUsedPercent"

const CPU_TEMP string = "CPU_TEMP"
const CPU_USAGE string = "CPU_USAGE"
const MEM_USAGE string = "MEMORY_USAGE"

var configData map[string][]ConfigDevice

var yamlFile string = "model.yaml"

type ConfigDevice struct {
	Id          string         `yaml:"id"`
	Name        string         `yaml:"name"`
	Description string         `yaml:"description"`
	Sensors     []ConfigSensor `yaml:"sensors"`
}

type ConfigSensor struct {
	Id           string                   `yaml:"id"`
	Name         string                   `yaml:"name"`
	Description  string                   `yaml:"description"`
	Unit         string                   `yaml:"unit"`
	SensorGroups []string                 `yaml:"sensorGroups"`
	Measurements []systeminfo.Measurement `yaml:"measurements"`
}

type ISensorFactory interface {
	SensorFactory(*cobra.Command) ([]sensors.SystemSensor, error)
	CreateSensor(ConfigSensor, string) (sensors.SystemSensor, error)
}

type SensorFactory struct{}

func (s *SensorFactory) SensorFactory(cmd *cobra.Command) ([]sensors.SystemSensor, error) {
	configDevices, err := readConfigDevices()
	if err != nil {
		return nil, err
	}

	kinds, err := getFlagValue(cmd, "sensor_group")
	if err != nil {
		return nil, err
	}

	sensors, err := s.checkCommand(kinds, configDevices[0])
	if err != nil {
		return nil, err
	}

	return sensors, nil
}

func readConfigDevices() ([]ConfigDevice, error) {
	file, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return nil, err
	}
	configData = make(map[string][]ConfigDevice)

	err = yaml.Unmarshal([]byte(file), &configData)
	if err != nil {
		return nil, err
	}

	return configData["devices"], nil
}

func (s *SensorFactory) checkCommand(kinds string, configDevice ConfigDevice) ([]sensors.SystemSensor, error) {
	regexpr := regexp.MustCompile(`[A-Za-z_]+`)

	userGroups := regexpr.FindAllString(kinds, -1)

	var sysSensors []sensors.SystemSensor
	var err error

	for _, uGroup := range userGroups {
		if uGroup == CPU_TEMP || uGroup == CPU_USAGE || uGroup == MEM_USAGE {
			sensorsToCreate := s.getSensorsToCreate(configDevice, sysSensors, uGroup)
			if err != nil {
				return nil, err
			}

			for _, confSensor := range sensorsToCreate {
				sensor, err := s.CreateSensor(confSensor, configDevice.Id)
				if err != nil {
					return nil, err
				}

				sysSensors = append(sysSensors, sensor)
			}

			continue
		}

		return nil, errors.New("invalid input for 'sensor_group' command")
	}

	return sysSensors, nil
}

func (s *SensorFactory) getSensorsToCreate(configDevice ConfigDevice, sysSensors []sensors.SystemSensor, userGroup string) []ConfigSensor {
	var sensorsToCreate []ConfigSensor
	for _, confSensor := range configDevice.Sensors {
		for _, group := range confSensor.SensorGroups {
			if group == userGroup {
				if exists := checkIfSensorExists(confSensor.Id, sysSensors); !exists {
					sensorsToCreate = append(sensorsToCreate, confSensor)
				}
			}
		}
	}

	return sensorsToCreate
}

func checkIfSensorExists(sensorId string, sensors []sensors.SystemSensor) bool {
	for _, sensor := range sensors {
		if sensor.GetInstanceId() == sensorId {
			return true
		}
	}

	return false
}

func (s *SensorFactory) CreateSensor(configSensor ConfigSensor, deviceId string) (sensors.SystemSensor, error) {
	b := &sensors.BaseSensor{Id: configSensor.Id, DeviceId: deviceId, Unit: configSensor.Unit}
	switch configSensor.Name {
	case CPU_TEMP_SENSOR:
		c := &sensors.CpuTempSensor{BaseSensor: *b}
		return c, nil
	case CPU_CORES_COUNTER:
		c := &sensors.CpuCoresCounter{BaseSensor: *b}
		return c, nil
	case CPU_FREQ_SENSOR:
		c := &sensors.CpuFrequencySensor{BaseSensor: *b}
		return c, nil
	case CPU_USAGE_SENSOR:
		c := &sensors.CpuUsageSensor{BaseSensor: *b}
		return c, nil
	case MEM_TOTAL_SENSOR:
		m := &sensors.MemTotalSensor{BaseSensor: *b}
		return m, nil
	case MEM_USED_SENSOR:
		m := &sensors.MemUsedSensor{BaseSensor: *b}
		return m, nil
	case MEM_AVAILABLE_SENSOR:
		m := &sensors.MemAvailableSensor{BaseSensor: *b}
		return m, nil
	case MEM_USAGE_SENSOR:
		m := &sensors.MemUsageSensor{BaseSensor: *b}
		return m, nil
	}

	return nil, errors.New("failed creating particular sensor")
}

func getFlagValueString(cmd *cobra.Command, flag string) (string, error) {
	flag, err := cmd.Flags().GetString("sensor_group")
	if err != nil {
		return "", err
	}

	return flag, nil
}
