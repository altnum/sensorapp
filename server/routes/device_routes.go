package routes

import (
	"github.com/altnum/sensorapp/services"
	"github.com/gorilla/mux"
)

func DeviceRoutesInit(router *mux.Router, service services.IService) {
	router.HandleFunc("/device", service.GetAll).Methods("GET")
	router.HandleFunc("/device/{id}", service.GetOne).Methods("GET")
	router.HandleFunc("/device", service.Create).Methods("POST")
	router.HandleFunc("/device", service.Update).Methods("PUT")
	router.HandleFunc("/device/{id}", service.Delete).Methods("DELETE")
}
