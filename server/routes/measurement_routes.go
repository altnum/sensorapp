package routes

import (
	"github.com/gorilla/mux"
	"github.wdf.sap.corp/I554249/sensor/services"
)

func MeasurementRoutesInit(router *mux.Router, service services.IService) {
	router.HandleFunc("/measurement", service.GetAll).Methods("GET")
	router.HandleFunc("/sensorAverageValue", service.GetSensorAverage).Methods("GET")
	router.HandleFunc("/sensorsCorrelationCoefficient", service.SensorCorrelation).Methods("GET")
	router.HandleFunc("/measurement/{sensorid}", service.GetOne).Methods("GET")
	router.HandleFunc("/measurement", service.Create).Methods("POST")
}
