package database_instances

import "context"

const GETALLDEVICES string = "getAllDevices"
const GETDEVICE string = "getDevice"
const CREATEDEVICE string = "createDevice"
const UPDATEDEVICE string = "updateDevice"
const DELETEDEVICE string = "deleteDevice"

const GETALLSENSORS string = "getAllSensors"
const GETSENSOR string = "getSensor"
const CREATESENSOR string = "createSensor"
const UPDATESENSOR string = "updateSensor"
const DELETESENSOR string = "deleteSensor"

const GETALLMEASUREMENTS string = "getAllMeasurements"
const GETSENSORAVERAGE string = "getSensorAverage"
const GETMEASUREMENTBYID string = "getMeasurementById"
const CREATEMEASUREMENT string = "createMeasurement"
const GETPEARSONSCOEFFICIENT string = "getPearsonsCorrelation"

type IDB interface {
	Open(context.Context) error
	Close() error
	GetPostgreInstance() IPostgreDB
	GetInfluxInstance() IInfluxDB
}
