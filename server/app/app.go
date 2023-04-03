package app

import (
	"context"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.wdf.sap.corp/I554249/sensor/databases"
	"github.wdf.sap.corp/I554249/sensor/databases/database_instances"
	. "github.wdf.sap.corp/I554249/sensor/databases/database_instances"
	"github.wdf.sap.corp/I554249/sensor/routes"
	"github.wdf.sap.corp/I554249/sensor/services"
)

var startServer = http.ListenAndServe

func Run() error {
	context := context.Background()
	router := mux.NewRouter()
	dbFactory := &databases.DatabaseFactory{}
	dbs, err := connectToDBs(context, dbFactory)
	if err != nil {
		return err
	}

	serviceFactory := &services.ServiceFactory{}

	initializeRoutes(router, dbs, serviceFactory)

	err = startServer(":8080", router)
	if err != nil {
		return err
	}

	err = closeDBs(dbs)
	if err != nil {
		return err
	}

	return nil
}

func initializeRoutes(router *mux.Router, dbs []IDB, serviceFactory services.IServiceFactory) {
	routeLoader := &routes.RouteLoader{}
	for _, db := range dbs {
		if reflect.TypeOf(db) == reflect.TypeOf(&database_instances.InfluxDB{}) {
			serviceMeasurement := serviceFactory.CreateService("measurement", dbs)
			routeLoader.LoadRoutes(router, serviceMeasurement)
		}
		if reflect.TypeOf(db) == reflect.TypeOf(&database_instances.PostgreDB{}) {
			serviceDevice := serviceFactory.CreateService("device", dbs)
			routeLoader.LoadRoutes(router, serviceDevice)
			serviceSensor := serviceFactory.CreateService("sensor", dbs)
			routeLoader.LoadRoutes(router, serviceSensor)
		}
	}
}

func connectToDBs(context context.Context, dbFactory databases.IDatabaseFactory) ([]IDB, error) {
	dbs := dbFactory.RetrieveDatabases()

	err := dbFactory.ConnectDatabases(context, dbs)
	if err != nil {
		return nil, err
	}

	return dbs, nil
}

func closeDBs(dbs []IDB) error {
	for _, db := range dbs {
		err := db.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
