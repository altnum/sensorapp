package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"github.com/altnum/sensorapp/databases/database_instances"
	. "github.com/altnum/sensorapp/databases/database_instances"
)

const DEVICE string = "device"
const SENSOR string = "sensor"
const MEASUREMENT string = "measurement"

type IService interface {
	GetAll(http.ResponseWriter, *http.Request)
	GetSensorAverage(http.ResponseWriter, *http.Request)
	GetOne(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	SensorCorrelation(http.ResponseWriter, *http.Request)
}

type IServiceFactory interface {
	CreateService(string, []IDB) IService
}

type ServiceFactory struct{}

func (s *ServiceFactory) CreateService(kind string, dbs []IDB) IService {
	var influxdb IDB
	var pgdb IDB
	for _, db := range dbs {
		if reflect.TypeOf(db) == reflect.TypeOf(&database_instances.InfluxDB{}) {
			influxdb = db
		}
		if reflect.TypeOf(db) == reflect.TypeOf(&database_instances.PostgreDB{}) {
			pgdb = db
		}
	}
	switch kind {
	case DEVICE:
		return &DeviceService{Pgdb: pgdb}
	case SENSOR:
		return &SensorService{Pgdb: pgdb}
	case MEASUREMENT:
		return &MeasurementService{Pgdb: pgdb, Influxdb: influxdb}
	}

	return nil
}

func OutputResponse(writer http.ResponseWriter, data interface{}) error {
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		err := errors.New("Failed encoding the data.")
		return err
	}

	return nil
}

func sendErrBadRequest(writer http.ResponseWriter, err error, status int) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)

	writer.Write([]byte(err.Error()))
}
