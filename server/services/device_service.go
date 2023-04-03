package services

import (
	"errors"
	"net/http"

	. "github.com/altnum/sensorapp/databases/database_instances"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var writeResponse = OutputResponse

type DeviceService struct {
	Pgdb IDB
}

func (d *DeviceService) GetAll(writer http.ResponseWriter, request *http.Request) {
	devices, err := d.Pgdb.GetPostgreInstance().GetAllDevices()
	if err != nil {
		sendErrBadRequest(writer, err, 400)
	}

	err = writeResponse(writer, devices)
	if err != nil {
		sendErrBadRequest(writer, err, 400)
	}
}

func (d *DeviceService) GetOne(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	device, err := d.Pgdb.GetPostgreInstance().GetDevice(vars["id"])
	if err != nil {
		sendErrBadRequest(writer, err, 400)
	}

	err = writeResponse(writer, device)
	if err != nil {
		sendErrBadRequest(writer, err, 400)
	}
}

func (d *DeviceService) Create(writer http.ResponseWriter, request *http.Request) {
	vars := make(map[string]string)
	vars["name"] = request.URL.Query().Get("name")
	vars["description"] = request.URL.Query().Get("description")
	err := d.Pgdb.GetPostgreInstance().CreateDevice(vars)
	if err != nil {
		err := errors.New("Failed creating the device.")
		sendErrBadRequest(writer, err, 400)
	}
}

func (d *DeviceService) Update(writer http.ResponseWriter, request *http.Request) {
	vars := make(map[string]string)
	vars["name"] = request.URL.Query().Get("name")
	vars["description"] = request.URL.Query().Get("description")
	vars["id"] = request.URL.Query().Get("id")
	err := d.Pgdb.GetPostgreInstance().UpdateDevice(vars)
	if err != nil {
		err := errors.New("Failed updating the device.")
		sendErrBadRequest(writer, err, 400)
	}
}

func (d *DeviceService) Delete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	err := d.Pgdb.GetPostgreInstance().DeleteDevice(vars["id"])
	if err != nil {
		err := errors.New("Failed deleting the device.")
		sendErrBadRequest(writer, err, 400)
	}
}

func (d *DeviceService) GetSensorAverage(writer http.ResponseWriter, request *http.Request) {

}

func (d *DeviceService) SensorCorrelation(writer http.ResponseWriter, request *http.Request) {

}
