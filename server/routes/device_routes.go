package routes

import (
	"github.com/gorilla/mux"
	"github.wdf.sap.corp/I554249/sensor/services"
)

func DeviceRoutesInit(router *mux.Router, service services.IService) {
	router.HandleFunc("/device", service.GetAll).Methods("GET")
	router.HandleFunc("/device/{id}", service.GetOne).Methods("GET")
	router.HandleFunc("/device", service.Create).Methods("POST")
	router.HandleFunc("/device", service.Update).Methods("PUT")
	router.HandleFunc("/device/{id}", service.Delete).Methods("DELETE")
}
