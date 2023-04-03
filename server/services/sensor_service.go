package services

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	. "github.wdf.sap.corp/I554249/sensor/databases/database_instances"
)

type SensorService struct {
	Pgdb IDB
}

func (s *SensorService) GetAll(writer http.ResponseWriter, request *http.Request) {
	sensors, err := s.Pgdb.GetPostgreInstance().GetAllSensors()
	if err != nil {
		sendErrBadRequest(writer, err, 400)
	}

	err = writeResponse(writer, sensors)
	if err != nil {
		sendErrBadRequest(writer, err, 400)
	}
}

func (s *SensorService) GetOne(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	sensor, err := s.Pgdb.GetPostgreInstance().GetSensor(vars["id"])
	if err != nil {
		sendErrBadRequest(writer, err, 400)
	}

	err = writeResponse(writer, sensor)
	if err != nil {
		sendErrBadRequest(writer, err, 400)
	}
}

func (s *SensorService) Create(writer http.ResponseWriter, request *http.Request) {
	vars := make(map[string]string)
	vars["deviceid"] = request.URL.Query().Get("deviceid")
	vars["name"] = request.URL.Query().Get("name")
	vars["description"] = request.URL.Query().Get("description")
	vars["unit"] = request.URL.Query().Get("unit")
	err := s.Pgdb.GetPostgreInstance().CreateSensor(vars)
	if err != nil {
		err := errors.New("Failed creating the sensor.")
		sendErrBadRequest(writer, err, 400)
	}
}

func (s *SensorService) Update(writer http.ResponseWriter, request *http.Request) {
	vars := make(map[string]string)
	vars["deviceid"] = request.URL.Query().Get("deviceid")
	vars["name"] = request.URL.Query().Get("name")
	vars["description"] = request.URL.Query().Get("description")
	vars["unit"] = request.URL.Query().Get("unit")
	vars["id"] = request.URL.Query().Get("id")
	err := s.Pgdb.GetPostgreInstance().UpdateSensor(vars)
	if err != nil {
		err := errors.New("Failed updating the sensor.")
		sendErrBadRequest(writer, err, 400)
	}
}

func (s *SensorService) Delete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	err := s.Pgdb.GetPostgreInstance().DeleteSensor(vars["id"])
	if err != nil {
		err := errors.New("Failed deleting the sensor.")
		sendErrBadRequest(writer, err, 400)
	}
}

func (s *SensorService) GetSensorAverage(writer http.ResponseWriter, request *http.Request) {
}

func (s *SensorService) SensorCorrelation(writer http.ResponseWriter, request *http.Request) {
}
