package cmd

import (
	"errors"

	timerMocks "github.com/altnum/sensorapp/execution_timer/mocks"
	formatMocks "github.com/altnum/sensorapp/formatter/mocks"
	systeminfo "github.com/altnum/sensorapp/system_info"
	"github.com/altnum/sensorapp/system_info/sensors"
	writerMocks "github.com/altnum/sensorapp/writer/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	sensorsMocks "github.com/altnum/sensorapp/system_info/sensors/mocks"
	sensorFactoryMocks "github.com/altnum/sensorapp/system_info/sensors_factory/mocks"
)

var _ = Describe("Testing the main calls and functionalities", func() {
	Context("createSensors()", func() {
		It("given no errors when creating the sensors, return slice of them", func() {
			mockSensorFactory := &sensorFactoryMocks.ISensorFactory{}
			mockSensor := &sensorsMocks.SystemSensor{}
			mockSensorFactory.On("SensorFactory", rootCmd).Return([]sensors.SystemSensor{mockSensor}, nil)
			sensorFactory = mockSensorFactory

			sensors := createSensors(rootCmd)

			Expect(sensors).ToNot(BeEmpty())
		})
	})

	Context("executeProgram()", func() {
		var (
			mockSystemSensor      sensorsMocks.SystemSensor
			formatter             formatMocks.IFormatter
			collectionMockSensors []sensors.SystemSensor
		)
		mockMeasurement := systeminfo.Measurement{}
		collectionMeasurements := []systeminfo.Measurement{mockMeasurement}

		BeforeEach(func() {
			mockSystemSensor = sensorsMocks.SystemSensor{}
			collectionMockSensors = []sensors.SystemSensor{&mockSystemSensor}
			formatter = formatMocks.IFormatter{}
			delta_duration = 0
			getFlagInfo = func(string) (string, error) {
				return "", nil
			}
			wg.Add(1)
		})
		It("given error, when formatting the output, then returns error", func() {
			mockSystemSensor.On("StartMeasurement", "").Return(mockMeasurement, nil)
			formatter.On("FormatOutput", rootCmd, collectionMeasurements[0]).Return("", errors.New("err"))
			var err error

			exec := func() {
				go func() {
					err = executeProgram(rootCmd, collectionMockSensors, &formatter)
					wg.Done()
				}()
			}

			exec()

			wg.Wait()
			Expect(err).ToNot(BeNil())
		})
		It("given error, when retrieving flag info in getSystemInfo, then return error", func() {
			getFlagInfo = func(string) (string, error) {
				return "", errors.New("err")
			}
			mockSystemSensor.On("StartMeasurement", "").Return(mockMeasurement, nil)
			formatter.On("FormatOutput", rootCmd, collectionMeasurements[0]).Return("{...}", nil)
			var err error

			exec := func() {
				go func() {
					err = executeProgram(rootCmd, collectionMockSensors, &formatter)
					wg.Done()
				}()
			}

			exec()

			wg.Wait()
			Expect(err).ToNot(BeNil())
		})
		It("given errors, when StartMeasurement, then return error ", func() {
			mockSystemSensor.On("StartMeasurement", "").Return(mockMeasurement, errors.New(""))
			formatter.On("FormatOutput", rootCmd, collectionMeasurements[0]).Return("{...}", nil)
			var err error

			exec := func() {
				go func() {
					err = executeProgram(rootCmd, collectionMockSensors, &formatter)
					wg.Done()
				}()
			}
			exec()

			wg.Wait()
			Expect(err).ToNot(BeNil())
		})
		It("given error, when writing the output, then return error", func() {
			getFlagInfo = func(s string) (string, error) {
				return "file", nil
			}
			mockSystemSensor.On("StartMeasurement", "file").Return(mockMeasurement, nil)
			formatter.On("FormatOutput", rootCmd, collectionMeasurements[0]).Return("", nil)

			mockWriter := &writerMocks.IWriter{}
			mockWriter.On("StartWriting", "file", mockMeasurement, "", "file").Return(errors.New(""))
			mockWriter.On("writeToDb", "file", mockMeasurement).Return(nil)
			outWriter = mockWriter

			total_duration = 1

			var err error

			exec := func() {
				go func() {
					err = executeProgram(rootCmd, collectionMockSensors, &formatter)
					wg.Done()
				}()
			}
			exec()

			wg.Wait()
			Expect(err).ToNot(BeNil())
		})
	})

	Context("setUpTimer()", func() {
		It("given no errors when setting up the timer then return no error", func() {

			timerMock := &timerMocks.ITimer{}
			timerMock.On("SetUpTimer", rootCmd).Return(nil)
			timerMock.On("GetTotalDuration").Return(int64(1))
			timerMock.On("GetDeltaDuration").Return(int64(0))
			timer = timerMock

			err := setUpTimer(rootCmd)

			Expect(err).To(BeNil())
		})
		It("given an error when setting up the timer then return it", func() {

			timerMock := &timerMocks.ITimer{}
			timerMock.On("SetUpTimer", rootCmd).Return(errors.New(""))
			timerMock.On("GetTotalDuration").Return(int64(1))
			timerMock.On("GetDeltaDuration").Return(int64(0))
			timer = timerMock

			err := setUpTimer(rootCmd)

			Expect(err).ToNot(BeNil())
		})
	})

	Context("buildOutput()", func() {
		mockMeasurement := systeminfo.Measurement{}
		collectionMeasurements := []systeminfo.Measurement{mockMeasurement}

		It("given no errors, when formatting the output, then returns the given string", func() {
			formatter := &formatMocks.IFormatter{}
			formatter.On("FormatOutput", rootCmd, collectionMeasurements[0]).Return("str", nil)

			outStr, err := buildOutputStr(rootCmd, collectionMeasurements[0], formatter)

			Expect(err).To(BeNil())
			Expect(outStr).To(BeIdenticalTo("str"))
		})
		It("given error, when formatting the output, then returns the given error", func() {
			formatter := &formatMocks.IFormatter{}
			err := errors.New("err")
			formatter.On("FormatOutput", rootCmd, collectionMeasurements[0]).Return("str", err)

			outStr, err := buildOutputStr(rootCmd, collectionMeasurements[0], formatter)

			Expect(err).ToNot(BeNil())
			Expect(outStr).To(BeIdenticalTo(""))
		})
	})

	Context("writeOutput()", func() {
		var mockMeasurement systeminfo.Measurement

		BeforeEach(func() {
			mockMeasurement = systeminfo.Measurement{}
		})
		It("given error when retrieving output_file value then return error", func() {
			getFlagInfo = func(s string) (string, error) {
				return "", errors.New("")
			}

			err := writeOutput(rootCmd, "", mockMeasurement)

			Expect(err).ToNot(BeNil())
		})
		It("given error when writing the output then return error", func() {
			mockWriter := &writerMocks.IWriter{}
			mockWriter.On("StartWriting", "...", mockMeasurement, "", "...").Return(errors.New(""))
			outWriter = mockWriter

			getFlagInfo = func(s string) (string, error) {
				return "...", nil
			}

			err := writeOutput(rootCmd, "", mockMeasurement)

			Expect(err).ToNot(BeNil())
		})
		It("given no error when writing the output then return nil", func() {
			mockWriter := &writerMocks.IWriter{}
			mockWriter.On("StartWriting", "", mockMeasurement, "", "").Return(nil)
			outWriter = mockWriter

			getFlagInfo = func(s string) (string, error) {
				return "", nil
			}

			err := writeOutput(rootCmd, "", mockMeasurement)

			Expect(err).To(BeNil())
		})
		It("given no value of output_file then return write to console", func() {
			getFlagInfo = func(s string) (string, error) {
				return "", nil
			}

			err := writeOutput(rootCmd, "", mockMeasurement)

			Expect(err).To(BeNil())
		})
	})
})
