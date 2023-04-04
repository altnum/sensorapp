//go:build windows
// +build windows

package sensors

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"syscall"

	"github.com/elastic/go-sysinfo"
	. "github.com/heucuva/go-win32"

	systeminfo "github.com/altnum/sensorapp/system_info"
	"golang.org/x/sys/windows"
)

type CpuTempSensor struct {
	BaseSensor
	Temperature float64
}

func (c *CpuTempSensor) StartMeasurement(unitF string) (systeminfo.Measurement, error) {
	infoLogger.Println("Starting temperature measurements")
	runWMICProcessElevatedMethod = c.runWMICProcessElevated

	err := c.SetTemperatureUnit(unitF)
	if err != nil {
		return measurement, err
	}

	err = c.SetTemperatureValue()
	if err != nil {
		warningLogger.Println(err)
		return measurement, err
	}

	measurement.SetTimeStamp()
	measurement.Value = c.Temperature
	measurement.SensorId = c.Id
	measurement.DeviceId = c.DeviceId

	return measurement, nil
}

// Setting the CPU temperature in Celsius degrees
func (c *CpuTempSensor) SetTemperatureValue() error {
	infoLogger.Println("Setting temperature value")
	var platformTypeStr string

	hostInfo, err := sysinfo.Host()
	if err != nil {
		return err
	}

	platformTypeStr = hostInfo.Info().OS.Type
	if platformTypeStr == "" {
		return errors.New("cannot get the OS type")
	}

	if platformTypeStr != "windows" {
		err = c.getCPUTemperature()
		if err != nil {
			warningLogger.Println("Running an external wmic process.")

			err := runWMICProcessElevatedMethod()
			if err != nil {
				return err
			}

			return nil
		}

		return nil
	}

	return runWMICProcessElevatedMethod()
}

func (c *CpuTempSensor) SetTemperatureUnit(unitF string) error {
	infoLogger.Println("Setting the temperature unit")

	unit := strings.ToLower(unitF)

	if unit != "f" && unit != "c" {
		return errors.New("invalid input for 'unit' command")
	}

	c.Unit = strings.ToUpper(unit)

	return nil
}

func (c *CpuTempSensor) getCPUTemperature() error {
	infoLogger.Println("getting the CPU temperature")

	sensors, err := systemSensorsInfo()
	if err != nil {
		c.Temperature = 0
		return err
	}

	for _, s := range sensors {
		if s.SensorKey == "coretemp_core0_input" {
			if s.Temperature <= 32 {
				return errors.New("invalid measured value for temperature")
			}
			if c.Unit == "F" {
				temperatureLong := (s.Temperature * 1.8) + 32
				c.Temperature, err = parseFloat(fmt.Sprintf("%.2f", temperatureLong), 64)
				if err != nil {
					return err
				}
			}
			c.Temperature = s.Temperature
			break
		}
	}

	if c.Temperature == 0 {
		return errors.New("invalid CPU temperature")
	}

	return nil
}

// Running the wmic process externally with administrator rights and returning the CPU temp in Celsius.
func (c *CpuTempSensor) runWMICProcessElevated() error {
	infoLogger.Println("Running WMIC process")

	verb := "runas"
	cwd, err := Getwd()
	if err != nil {
		return err
	}

	outputFileDest := cwd + "\\tempWMICout.txt"

	if err != nil {
		return err
	}

	verbPtr, err := syscall.UTF16PtrFromString(verb)
	if err != nil {
		return err
	}
	exePtr, err := syscall.UTF16PtrFromString("cmd.exe")
	if err != nil {
		return err
	}
	cwdPtr, err := syscall.UTF16PtrFromString(cwd)
	if err != nil {
		return err
	}
	argPtr, err := syscall.UTF16PtrFromString("/c wmic /namespace:\\\\root\\wmi PATH MSAcpi_ThermalZoneTemperature get CurrentTemperature > " + outputFileDest)
	if err != nil {
		return err
	}

	handle, err := syscall.GetCurrentProcess()

	err = windows.ShellExecute(windows.Handle(handle), verbPtr, exePtr, argPtr, cwdPtr, 1)
	if err != nil {
		return err
	}

	err = WaitForSingleObject(windows.Handle(handle), syscall.INFINITE)
	if err != nil {
		return err
	}

	c.readTempFromFile(outputFileDest)

	return nil
}

func (c *CpuTempSensor) readTempFromFile(fileDest string) error {
	outputFileInfo, err := ioutil.ReadFile(fileDest)
	if err != nil {
		return err
	}

	b := []byte(outputFileInfo)
	myString := string(b)

	tokens := strings.Split(myString, "\n")

	token := strings.Split(tokens[1], "\x00")

	err = c.setTempValueFromString(token)
	if err != nil {
		return err
	}

	return nil
}

func (c *CpuTempSensor) setTempValueFromString(token []string) error {
	infoLogger.Println("Setting the temperature value derived from the system")

	var strToConv string

	for _, ch := range token {
		char := []rune(ch)
		if len(char) == 0 {
			continue
		}
		if int(char[0]) >= 48 && int(char[0]) <= 57 {
			strToConv += ch
		}
	}

	tempOut, err := strconv.ParseInt(strToConv, 0, 64)
	if err != nil {
		return err
	}

	temperatureLong := ((float64(tempOut) / 10) - 273.15)

	if c.Unit == "F" {
		temperatureLong = (temperatureLong * 1.8) + 32
		c.Temperature, err = parseFloat(fmt.Sprintf("%.2f", temperatureLong), 64)
		if err != nil {
			return err
		}

		return nil
	}

	strTemp := fmt.Sprintf("%.2f", temperatureLong)

	formatedTemp, err := parseFloat(strTemp, 64)
	if err != nil {
		return err
	}

	c.Temperature = formatedTemp

	return nil
}
