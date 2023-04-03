package routes

import (
	"github.com/gorilla/mux"
	"github.wdf.sap.corp/I554249/sensor/services"
)

func SensorRoutesInit(router *mux.Router, service services.IService) {
	router.HandleFunc("/sensor", service.GetAll).Methods("GET")
	router.HandleFunc("/sensor/{id}", service.GetOne).Methods("GET")
	router.HandleFunc("/sensor", service.Create).Methods("POST")
	router.HandleFunc("/sensor", service.Update).Methods("PUT")
	router.HandleFunc("/sensor/{id}", service.Delete).Methods("DELETE")
}
