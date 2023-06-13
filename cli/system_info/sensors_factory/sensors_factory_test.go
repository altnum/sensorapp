package sensorsfactory

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

var _ = Describe("SensorsFactory functionality", func() {
	factory := SensorFactory{}

	Context("SensorsFactory()", func() {
		command := cobra.Command{}
		It("given error when retrieving sensor_group command, then return error", func() {
			getFlagValue = func(cmd *cobra.Command, flag string) (string, error) {
				return "", errors.New("")
			}

			sensors, err := factory.SensorFactory(&command)
			Expect(err).ToNot(BeNil())
			Expect(sensors).To(BeNil())
		})
		It("given incorrect input when using sensor_group command, then return error", func() {
			getFlagValue = func(cmd *cobra.Command, flag string) (string, error) {
				return "C", nil
			}

			sensors, err := factory.SensorFactory(&command)
			Expect(err).ToNot(BeNil())
			Expect(sensors).To(BeNil())
		})
		It("given correct input when using sensor_group command, then return slice of sensors", func() {
			yamlFile = "testData/test.yaml"
			getFlagValue = func(cmd *cobra.Command, flag string) (string, error) {
				return "CPU_TEMP,MEMORY_USAGE,CPU_USAGE", nil
			}

			sensors, err := factory.SensorFactory(&command)
			Expect(err).To(BeNil())
			Expect(sensors).ToNot(BeEmpty())
		})
	})

	Context("getFlagValueString()", func() {
		It("given error when retrieving the command value, then return error", func() {
			command := cobra.Command{}
			_, err := getFlagValueString(&command, "sensor_group")

			Expect(err).ToNot(BeNil())
		})
	})
})
