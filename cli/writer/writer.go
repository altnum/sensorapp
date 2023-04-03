package writer

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/altnum/sensorapp/logger"
	systeminfo "github.com/altnum/sensorapp/system_info"
)

var printToConsole = fmt.Println
var sendUrlForm = http.PostForm
var errorLogger = logger.GetLogger().Error

type IWriter interface {
	StartWriting(string, systeminfo.Measurement, string, string) error
}

type Writer struct{}

func (p *Writer) StartWriting(serverURL string, measurement systeminfo.Measurement, output string, destFile string) error {

	var waitGr sync.WaitGroup

	if destFile != "" {
		file, err := os.OpenFile(destFile+".csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return err
		}

		fileWriter := csv.NewWriter(file)
		fileWriter.Comma = '|'

		go func() {
			waitGr.Add(1)
			defer waitGr.Done()

			defer file.Close()
			defer fileWriter.Flush()

			err := p.writeToCsvFile(measurement, fileWriter)
			if err != nil {
				errorLogger.Println(err)
			}
		}()
	}

	go func() {
		waitGr.Add(1)
		defer waitGr.Done()

		err := p.writeToConsole(output)
		if err != nil {
			errorLogger.Fatalln(err)
		}
	}()

	if serverURL != "" {
		waitGr.Add(1)
		go func() {
			defer waitGr.Done()

			err := p.writeToDb(serverURL, measurement)
			if err != nil {
				errorLogger.Println(err)
			}
		}()
	}

	waitGr.Wait()

	return nil
}

func (p *Writer) writeToCsvFile(measurement systeminfo.Measurement, writer *csv.Writer) error {
	toWrite := []string{measurement.MeasuredAt, measurement.DeviceId, measurement.SensorId}
	err := writer.Write(toWrite)
	if err != nil {
		return err
	}

	return nil
}

func (p *Writer) writeToConsole(output string) error {
	_, err := printToConsole(output)
	if err != nil {
		return err
	}

	return nil
}

func (p *Writer) writeToDb(serverURL string, measurement systeminfo.Measurement) error {
	measurementData := url.Values{
		"measuredat": {measurement.MeasuredAt},
		"deviceid":   {measurement.DeviceId},
		"sensorid":   {measurement.SensorId},
		"value":      {fmt.Sprint(measurement.Value)},
	}

	_, err := sendUrlForm(serverURL, measurementData)
	if err != nil {
		return err
	}

	return nil
}
