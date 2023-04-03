package routes

import (
	"github.com/altnum/sensorapp/services"
	"github.com/gorilla/mux"
)

func MeasurementRoutesInit(router *mux.Router, service services.IService) {
	router.HandleFunc("/measurement", service.GetAll).Methods("GET")
	router.HandleFunc("/sensorAverageValue", service.GetSensorAverage).Methods("GET")
	router.HandleFunc("/sensorsCorrelationCoefficient", service.SensorCorrelation).Methods("GET")
	router.HandleFunc("/measurement/{sensorid}", service.GetOne).Methods("GET")
	router.HandleFunc("/measurement", service.Create).Methods("POST")
}
