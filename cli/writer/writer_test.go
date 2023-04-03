package writer

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	systeminfo "github.com/altnum/sensorapp/system_info"
)

var _ = Describe("Testing writer package", func() {
	Context("writeToConsole", func() {
		var writerMock Writer
		BeforeEach(func() {
			writerMock = Writer{}
		})
		It("given no errors while printing, then execute normally", func() {
			printToConsole = func(i ...interface{}) (n int, err error) {
				return 0, nil
			}
			writerMock := Writer{}
			err := writerMock.writeToConsole("")

			Expect(err).To(BeNil())
		})
		It("given error while printing, then return error", func() {
			printToConsole = func(i ...interface{}) (n int, err error) {
				return 0, errors.New("")
			}

			err := writerMock.writeToConsole("")

			Expect(err).ToNot(BeNil())
		})
	})

	Context("writeToCsvFile", func() {
		It("given no error while printing, then print correctly to file", func() {
			mockMeasurement := systeminfo.Measurement{}
			mockMeasurement.SetTimeStamp()
			mockMeasurement.DeviceId = "1"
			mockMeasurement.SensorId = "1"

			file, _ := os.Create("test.csv")
			fileWriter := csv.NewWriter(file)

			writerMock := Writer{}

			err := writerMock.writeToCsvFile(mockMeasurement, fileWriter)

			fileWriter.Flush()

			fileToCheck, _ := os.ReadFile("test.csv")

			strToCheck := string(fileToCheck)
			Expect(err).To(BeNil())
			Expect(strToCheck).To(BeEquivalentTo(mockMeasurement.MeasuredAt + "," + mockMeasurement.DeviceId + "," + mockMeasurement.SensorId + "\n"))
		})
	})

	Context("StartParallelWriting", func() {
		It("given no errors, then return equivalent results to the console and the output file", func() {
			var testConsole string
			timeMeasure1 := time.Now()
			timeNow1 := fmt.Sprint(timeMeasure1.Format("2006-01-02 15:04:05"))

			var resultData url.Values
			sendUrlForm = func(url string, data url.Values) (*http.Response, error) {
				resultData = data
				return nil, nil
			}

			printToConsole = func(a ...interface{}) (n int, err error) {
				testConsole = fmt.Sprintf("%v", a)
				testConsole = strings.Replace(testConsole, "[", "", -1)
				testConsole = strings.Replace(testConsole, "]", "", -1)
				return 0, nil
			}

			writerMock := Writer{}

			measurement := systeminfo.Measurement{MeasuredAt: timeNow1, Value: 50, SensorId: "-1", DeviceId: "-1"}

			os.Create("test.csv")

			writerMock.StartWriting("test", measurement, measurement.MeasuredAt, "test")

			fileToCheck, _ := os.ReadFile("test.csv")

			tokens := string(fileToCheck)
			strsToCheck := strings.Split(tokens, "|")

			Expect(strsToCheck[0]).To(BeEquivalentTo(testConsole))
			Expect(resultData.Get("measuredat")).To(BeIdenticalTo(timeNow1))
		})
	})
})
