package routes

import (
	"reflect"

	"github.com/gorilla/mux"
	"github.wdf.sap.corp/I554249/sensor/services"
)

type IRouteLoader interface {
	LoadRoutes()
}

type RouteLoader struct{}

func (r *RouteLoader) LoadRoutes(router *mux.Router, service services.IService) {
	if reflect.TypeOf(service) == reflect.TypeOf(&services.DeviceService{}) {
		DeviceRoutesInit(router, service)
	}
	if reflect.TypeOf(service) == reflect.TypeOf(&services.SensorService{}) {
		SensorRoutesInit(router, service)
	}
	if reflect.TypeOf(service) == reflect.TypeOf(&services.MeasurementService{}) {
		MeasurementRoutesInit(router, service)
	}
}
