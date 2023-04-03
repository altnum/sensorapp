package formatter

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
	systeminfo "github.wdf.sap.corp/I554249/sensor/system_info"
)

var _ = Describe("Testing formatter package", func() {
	Context("FormatOutput()", func() {
		measurement := systeminfo.Measurement{MeasuredAt: "", Value: 2}

		format := Formatter{}
		It("given error when called getFlagValue, returns error", func() {
			cmd := &cobra.Command{}
			getFlagValue = func(cmd *cobra.Command, flag string) (string, error) {
				return "", errors.New("err")
			}

			outputStr, err := format.FormatOutput(cmd, measurement)

			Expect(err).ToNot(BeNil())
			Expect(outputStr).To(BeIdenticalTo(""))
		})
		It("given invalid flag value when called getFlagValue, then return error", func() {
			cmd := &cobra.Command{}
			getFlagValue = func(cmd *cobra.Command, flag string) (string, error) {
				return "test", nil
			}

			outputStr, err := format.FormatOutput(cmd, measurement)

			Expect(err).ToNot(BeNil())
			Expect(outputStr).To(BeIdenticalTo(""))
		})
		It("given yaml flag value when called getFlagValue, then prints the object in yaml format", func() {
			cmd := &cobra.Command{}
			getFlagValue = func(cmd *cobra.Command, flag string) (string, error) {
				return "yaml", nil
			}

			outputStr, err := format.FormatOutput(cmd, measurement)

			Expect(err).To(BeNil())
			Expect(outputStr).To(BeIdenticalTo("measuredAt: \"\"\nvalue: 2\nsensorId: \"\"\ndeviceId: \"\"\n\n"))
		})
		It("given json flag value when called getFlagValue, then prints the object in json format", func() {
			cmd := &cobra.Command{}
			getFlagValue = func(cmd *cobra.Command, flag string) (string, error) {
				return "json", nil
			}

			outputStr, err := format.FormatOutput(cmd, measurement)

			Expect(err).To(BeNil())
			Expect(outputStr).To(BeIdenticalTo("{\"measuredAt\":\"\",\"value\":2,\"sensorId\":\"\",\"deviceId\":\"\"}\n"))
		})
		It("given valid flag value and error from jsonMarshal, then return error", func() {
			cmd := &cobra.Command{}
			getFlagValue = func(cmd *cobra.Command, flag string) (string, error) {
				return "json", nil
			}

			jsonMarshal = func(v interface{}) ([]byte, error) {
				return []byte(""), errors.New("err")
			}

			outputStr, err := format.FormatOutput(cmd, measurement)

			Expect(err).ToNot(BeNil())
			Expect(outputStr).To(BeIdenticalTo(""))
		})
		It("given valid flag value and error from yamlMarshal, then return error", func() {
			cmd := &cobra.Command{}
			getFlagValue = func(cmd *cobra.Command, flag string) (string, error) {
				return "yaml", nil
			}

			yamlMarshal = func(v interface{}) ([]byte, error) {
				return []byte(""), errors.New("err")
			}

			outputStr, err := format.FormatOutput(cmd, measurement)

			Expect(err).ToNot(BeNil())
			Expect(outputStr).To(BeIdenticalTo(""))
		})
		It("given empty flag value and error from marshalling, then return error", func() {
			cmd := &cobra.Command{}
			getFlagValue = func(cmd *cobra.Command, flag string) (string, error) {
				return "", nil
			}

			jsonMarshal = func(v interface{}) ([]byte, error) {
				return []byte(""), errors.New("err")
			}

			outputStr, err := format.FormatOutput(cmd, measurement)

			Expect(err).ToNot(BeNil())
			Expect(outputStr).To(BeIdenticalTo(""))
		})
	})

	Context("FormatInfoOutput()", func() {
		measurement := systeminfo.Measurement{MeasuredAt: "", Value: 2}
		format := Formatter{}
		It("given error when called getFlagValue, returns error", func() {

			outputStr, err := format.FormatInfoOutput("", measurement)

			Expect(err).ToNot(BeNil())
			Expect(outputStr).To(BeIdenticalTo(""))
		})
	})
})
