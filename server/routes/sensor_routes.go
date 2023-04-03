package routes

import (
	"github.com/altnum/sensorapp/services"
	"github.com/gorilla/mux"
)

func SensorRoutesInit(router *mux.Router, service services.IService) {
	router.HandleFunc("/sensor", service.GetAll).Methods("GET")
	router.HandleFunc("/sensor/{id}", service.GetOne).Methods("GET")
	router.HandleFunc("/sensor", service.Create).Methods("POST")
	router.HandleFunc("/sensor", service.Update).Methods("PUT")
	router.HandleFunc("/sensor/{id}", service.Delete).Methods("DELETE")
}
