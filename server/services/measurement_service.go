package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	. "github.com/altnum/sensorapp/databases/database_instances"
	"github.com/altnum/sensorapp/models"
	"github.com/gorilla/mux"
)

type MeasurementService struct {
	Influxdb IDB
	Pgdb     IDB
}

func (m *MeasurementService) GetAll(writer http.ResponseWriter, request *http.Request) {
	measures, err := m.Influxdb.GetInfluxInstance().GetAllMeasurements(context.Background())
	if err != nil {
		sendErrBadRequest(writer, err, 400)
	}

	err = writeResponse(writer, measures)
	if err != nil {
		sendErrBadRequest(writer, err, 400)
	}
}

func (m *MeasurementService) GetSensorAverage(writer http.ResponseWriter, request *http.Request) {
	vars := make(map[string]string)
	vars["deviceid"] = request.URL.Query().Get("deviceid")
	vars["sensorid"] = request.URL.Query().Get("sensorid")
	vars["startTime"] = request.URL.Query().Get("startTime")
	vars["endTime"] = request.URL.Query().Get("endTime")
	averageMeasure, err := m.Influxdb.GetInfluxInstance().GetSensorAverage(context.Background(), vars)
	if err != nil {
		sendErrBadRequest(writer, err, 400)
	}

	err = writeResponse(writer, averageMeasure)
	if err != nil {
		sendErrBadRequest(writer, err, 400)
	}
}

func (m *MeasurementService) GetOne(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	measures, err := m.Influxdb.GetInfluxInstance().GetMeasurementById(context.Background(), vars["sensorid"])
	if err != nil {
		sendErrBadRequest(writer, err, 400)
	}

	err = writeResponse(writer, measures)
	if err != nil {
		sendErrBadRequest(writer, err, 400)
	}
}

func (m *MeasurementService) SensorCorrelation(writer http.ResponseWriter, request *http.Request) {
	vars := make(map[string]string)
	vars["deviceid1"] = request.URL.Query().Get("deviceid1")
	vars["deviceid2"] = request.URL.Query().Get("deviceid2")
	vars["sensorid1"] = request.URL.Query().Get("sensorid1")
	vars["sensorid2"] = request.URL.Query().Get("sensorid2")
	vars["startTime"] = request.URL.Query().Get("startTime")
	vars["endTime"] = request.URL.Query().Get("endTime")
	correlationCoefficient, err := m.Influxdb.GetInfluxInstance().GetPearsonsCoefficient(context.Background(), vars)
	if err != nil {
		sendErrBadRequest(writer, err, 400)
	}

	err = writeResponse(writer, correlationCoefficient)
	if err != nil {
		sendErrBadRequest(writer, err, 400)
	}
}

func (m *MeasurementService) Create(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	vars := make(map[string]string)
	vars["measuredat"] = request.Form.Get("measuredat")
	vars["sensorid"] = request.Form.Get("sensorid")
	vars["value"] = request.Form.Get("value")
	vars["deviceid"] = request.Form.Get("deviceid")

	device, err := m.checkDeviceExistance(vars["deviceid"])
	if err != nil {
		sendErrBadRequest(writer, err, 400)

		return
	}

	sensorExist := m.checkSensorExistance(vars["sensorid"], device.Sensors)
	if !sensorExist {
		err := errors.New("No such sensor with the specified id in the given device.")
		sendErrBadRequest(writer, err, 400)

		return
	}

	err = m.Influxdb.GetInfluxInstance().CreateMeasurement(vars)
	if err != nil {
		err := errors.New("Failed creating the measurement.")
		sendErrBadRequest(writer, err, 400)
	}
}

func (m *MeasurementService) checkSensorExistance(sensorid string, devicesensors []*models.Sensors) bool {
	for _, devicesensor := range devicesensors {
		if fmt.Sprint(devicesensor.Id) == sensorid {
			return true
		}
	}

	return false
}

func (m *MeasurementService) checkDeviceExistance(deviceid string) (models.Device, error) {
	device, err := m.Pgdb.GetPostgreInstance().GetDevice(deviceid)
	if err != nil {
		return device, err
	}

	return device, nil
}

func (m *MeasurementService) Update(writer http.ResponseWriter, request *http.Request) {
}

func (m *MeasurementService) Delete(writer http.ResponseWriter, request *http.Request) {
}
