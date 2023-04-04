package sensors

import (
	"errors"
	"fmt"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shirou/gopsutil/host"
)

var _ = Describe("Measurements", func() {
	Context("CpuCoresCounter struct", func() {
		sensor := &CpuCoresCounter{}
		It("given valid number of cores when calling StartMeasurement() return no error", func() {
			var err error
			_, err = sensor.StartMeasurement("")

			Expect(err).To(BeNil())
			Expect(sensor.NumberOfCores).To(BeNumerically(">", 0))
		})

		It("given invalid number of cores when calling StartMeasurement() return error", func() {
			getCoreNum = func() int {
				return 0
			}
			var err error
			_, err = sensor.StartMeasurement("")

			Expect(err).ToNot(BeNil())
		})

	})

	Context("CpuFrequencySensor struct", func() {
		sensor := &CpuFrequencySensor{}
		AfterEach(func() {
			parseFloat = strconv.ParseFloat
		})
		It("given no error when parsing then return frequency result", func() {
			_, err := sensor.StartMeasurement("")

			Expect(err).To(BeNil())
			Expect(sensor.Frequency).To(BeNumerically(">", 0))
		})
		It("given error when parsing then return error", func() {
			parseFloat = func(s string, bitSize int) (float64, error) {
				return 0, errors.New("")
			}
			_, err := sensor.StartMeasurement("")

			Expect(err).ToNot(BeNil())
		})
	})

	Context("CpuTempSensor struct", func() {
		var cpuSensor CpuTempSensor
		BeforeEach(func() {
			cpuSensor = CpuTempSensor{}
		})
		AfterEach(func() {
			parseFloat = strconv.ParseFloat
		})

		Context("SetTemperatureValue()", func() {
			It("should return error", func() {
				runWMICProcessElevatedMethod = func() error {
					return errors.New("err")
				}
				err := cpuSensor.SetTemperatureValue()

				Expect(err).ToNot(BeNil())
			})
		})

		Context("SetTemperatureUnit()", func() {
			type tempUnitTest struct {
				input    string
				expected string
			}

			var unitTests = []tempUnitTest{
				{"c", "C"}, {"C", "C"}, {"F", "F"}, {"f", "F"},
			}
			BeforeEach(func() {
				cpuSensor.Temperature = 100
			})
			It("given a specific valid input, when setting the unit, then unit to be set to the corresponding capital letter", func() {
				for _, testInput := range unitTests {
					err := cpuSensor.SetTemperatureUnit(testInput.input)
					Expect(err).To(BeNil())
					Expect(cpuSensor.Unit).To(BeIdenticalTo(testInput.expected))
				}

			})
			It("given an invalid input, when setting the unit, then return error", func() {
				err := cpuSensor.SetTemperatureUnit("test")

				Expect(err).ToNot(BeNil())
			})
		})

		Context("getCPUTemperature()", func() {
			var tempStats []host.TemperatureStat
			systemSensorsInfo = func() ([]host.TemperatureStat, error) {
				return tempStats, nil
			}

			BeforeEach(func() {
				tempStats = []host.TemperatureStat{}
			})
			It("should set temperature, when the OS provides info, in the correct value", func() {
				var infoStat host.TemperatureStat
				infoStat.SensorKey = "coretemp_core0_input"
				infoStat.Temperature = float64(50)

				tempStats = append(tempStats, infoStat)

				err := cpuSensor.getCPUTemperature()

				Expect(err).To(BeNil())
				Expect(cpuSensor.Temperature).To(BeEquivalentTo(float64(50)))
			})

			It("should return error, when the OS doesn't provide info", func() {
				err := cpuSensor.getCPUTemperature()

				Expect(err).ToNot(BeNil())
			})

			It("should return error, when the OS provides incorrect info", func() {
				var infoStat host.TemperatureStat
				infoStat.SensorKey = "coretemp_core0_input"
				infoStat.Temperature = float64(0)

				tempStats = append(tempStats, infoStat)

				err := cpuSensor.getCPUTemperature()

				Expect(err).ToNot(BeNil())
			})
		})

		Context("StartMeasurements()", func() {
			It("given some valid input for unit then CpuSensor fields are always set and no error", func() {
				measurement, err := cpuSensor.StartMeasurement("f")

				Expect(err).To(BeNil())
				Expect(measurement.MeasuredAt).ToNot(BeEquivalentTo(""))
				Expect(measurement.Value).ToNot(BeEquivalentTo(0))
			})
			It("given an error while executing then should return error", func() {
				Getwd = func() (dir string, err error) {
					return "", errors.New("")
				}

				_, err := cpuSensor.StartMeasurement("")

				Expect(err).ToNot(BeNil())
			})
		})

		Context("setTempValueFromString()", func() {
			It("given error when parsing invalid input then return error", func() {
				var strings = []string{""}
				err := cpuSensor.setTempValueFromString(strings)

				Expect(err).ToNot(BeNil())
			})

			It("given Unit=F when calculating the temperature, then convert to Fahrenheit", func() {
				cpuSensor.Unit = "F"
				var strings = []string{"test", "3232"}
				err := cpuSensor.setTempValueFromString(strings)

				checkTemp, _ := strconv.ParseInt(strings[1], 0, 64)
				checkTempLong := ((float64(checkTemp) / 10) - 273.15)
				checkTempLong = (checkTempLong * 1.8) + 32
				checkTempLong, _ = parseFloat(fmt.Sprintf("%.2f", checkTempLong), 64)

				Expect(err).To(BeNil())
				Expect(cpuSensor.Temperature).To(BeNumerically("==", checkTempLong))
			})

			It("given Unit=F when calculating the temperature, then convert to Fahrenheit", func() {
				cpuSensor.Unit = "F"
				var strings = []string{"test", "3232"}

				parseFloat = func(s string, bitSize int) (float64, error) {
					return 0, errors.New("err")
				}

				err := cpuSensor.setTempValueFromString(strings)

				Expect(err).ToNot(BeNil())
			})

			It("given error when parsing valid input then return error", func() {
				var strings = []string{"test", "3232"}
				parseFloat = func(s string, bitSize int) (float64, error) {
					return 0, errors.New("")
				}

				err := cpuSensor.setTempValueFromString(strings)

				Expect(err).ToNot(BeNil())
			})
		})
	})

	Context("CpuUsageSensor struct", func() {
		sensor := CpuUsageSensor{}
		AfterEach(func() {
			parseFloat = strconv.ParseFloat
		})
		It("given no errors when getting the usage then return proper measurement struct", func() {
			measurement, err := sensor.StartMeasurement("")

			Expect(err).To(BeNil())
			Expect(measurement.Value).To(BeNumerically(">", 0))
		})
		It("given errors when parsing then return error", func() {
			parseFloat = func(s string, bitSize int) (float64, error) {
				return 0, errors.New("")
			}
			_, err := sensor.StartMeasurement("")

			Expect(err).ToNot(BeNil())
		})
	})

	Context("MemAvailableSensor struct", func() {
		sensor := MemAvailableSensor{}
		It("return proper measurement struct", func() {
			measurement, err := sensor.StartMeasurement("")

			Expect(err).To(BeNil())
			Expect(measurement.Value).To(BeNumerically(">", 0))
		})
	})

	Context("MemTotalSensor struct", func() {
		sensor := MemTotalSensor{}
		It("given no error when measuring total memory then return proper measurement struct", func() {
			measurement, err := sensor.StartMeasurement("")

			Expect(err).To(BeNil())
			Expect(measurement.Value).To(BeNumerically(">", 0))
		})
	})

	Context("MemUsageSensor struct", func() {
		sensor := MemUsageSensor{}
		It("given no error when measuring usage of memory then return proper measurement struct", func() {
			measurement, err := sensor.StartMeasurement("")

			Expect(err).To(BeNil())
			Expect(measurement.Value).To(BeNumerically(">", 0))
		})
	})

	Context("MemUsedSensor struct", func() {
		sensor := MemUsedSensor{}
		It("given no error when measuring used memory then return proper measurement struct", func() {
			measurement, err := sensor.StartMeasurement("")

			Expect(err).To(BeNil())
			Expect(measurement.Value).To(BeNumerically(">", 0))
		})
	})

})
