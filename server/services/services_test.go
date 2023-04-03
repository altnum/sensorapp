package services

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/altnum/sensorapp/databases/database_instances/mocks"
	"github.com/altnum/sensorapp/models"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testing different services functionality", func() {
	context := context.Background()
	Context("DevicesServices", func() {
		Context("GetAll", func() {
			var pgdb *mocks.IPostgreDB
			var router *mux.Router
			BeforeEach(func() {
				pgdb = &mocks.IPostgreDB{}
				pgdb.On("GetPostgreInstance").Return(pgdb)
				deviceService := &DeviceService{Pgdb: pgdb}
				router = mux.NewRouter()
				router.HandleFunc("/device", deviceService.GetAll).Methods("GET")
			})
			AfterEach(func() {
				writeResponse = OutputResponse
			})
			It("given no error while querying all the devices, return them", func() {
				var result interface{}
				writeResponse = func(writer http.ResponseWriter, data interface{}) error {
					result = data
					return nil
				}
				devices := []models.Device{*&models.Device{}}
				pgdb.On("GetAllDevices").Return(devices, nil)
				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "http://localhost:8080/device", nil)

				router.ServeHTTP(respRec, req)

				Expect(result).ToNot(BeEmpty())
				Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
			})
			It("given error while querying all the devices, return StatusBadRequest", func() {
				devices := []models.Device{}
				pgdb.On("GetAllDevices").Return(devices, errors.New(""))

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "http://localhost:8080/device", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusBadRequest))
			})
		})

		Context("GetOne", func() {
			var pgdb *mocks.IPostgreDB
			var router *mux.Router
			BeforeEach(func() {
				pgdb = &mocks.IPostgreDB{}
				pgdb.On("GetPostgreInstance").Return(pgdb)
				deviceService := &DeviceService{Pgdb: pgdb}
				router = mux.NewRouter()
				router.HandleFunc("/device/{id}", deviceService.GetOne).Methods("GET")
			})
			AfterEach(func() {
				writeResponse = OutputResponse
			})
			It("given no error while querying one of the devices, return StatusOK", func() {
				var result interface{}
				device := models.Device{}
				writeResponse = func(writer http.ResponseWriter, data interface{}) error {
					result = data
					return nil
				}
				pgdb.On("GetDevice", "1").Return(device, nil)

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "http://localhost:8080/device/1", nil)
				router.ServeHTTP(respRec, req)

				Expect(result).ToNot(BeNil())
				Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
			})
			It("given error while querying the device, return the StatusBadRequest", func() {
				device := models.Device{}
				writeResponse = func(writer http.ResponseWriter, data interface{}) error {
					return errors.New("")
				}
				pgdb.On("GetDevice", "1").Return(device, errors.New(""))

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "http://localhost:8080/device/1", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusBadRequest))
			})
		})

		Context("Create", func() {
			var pgdb *mocks.IPostgreDB
			var vars map[string]string
			var deviceService *DeviceService
			var router *mux.Router
			BeforeEach(func() {
				pgdb = &mocks.IPostgreDB{}
				vars = make(map[string]string)
				vars["description"] = ""
				vars["name"] = ""
				deviceService = &DeviceService{Pgdb: pgdb}
				router = mux.NewRouter()
				router.HandleFunc("/device", deviceService.Create).Methods("POST")
				pgdb.On("GetPostgreInstance").Return(pgdb)
			})
			AfterEach(func() {
				writeResponse = OutputResponse
			})
			It("given no error while creating one device, return statusOK", func() {
				pgdb.On("CreateDevice", vars).Return(nil)

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "http://localhost:8080/device", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
			})
			It("given error while creating the device, return StatusBadRequest", func() {
				pgdb.On("CreateDevice", vars).Return(errors.New(""))

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "http://localhost:8080/device", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusBadRequest))
			})
		})

		Context("Update", func() {
			var pgdb *mocks.IPostgreDB
			var vars map[string]string
			var deviceService *DeviceService
			var router *mux.Router
			BeforeEach(func() {
				pgdb = &mocks.IPostgreDB{}
				vars = make(map[string]string)
				vars["description"] = ""
				vars["name"] = ""
				deviceService = &DeviceService{Pgdb: pgdb}
				router = mux.NewRouter()
				router.HandleFunc("/device", deviceService.Update).Methods("PUT")
				pgdb.On("GetPostgreInstance").Return(pgdb)
			})
			AfterEach(func() {
				writeResponse = OutputResponse
			})
			It("given no error while updating one device, return statusOK", func() {
				vars["id"] = "1"
				pgdb.On("UpdateDevice", vars).Return(nil)
				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("PUT", "http://localhost:8080/device?id=1", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
			})
			It("given error while updating the device, return StatusBadRequest", func() {
				vars["id"] = ""
				pgdb.On("UpdateDevice", vars).Return(errors.New(""))

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("PUT", "http://localhost:8080/device", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusBadRequest))
			})
		})

		Context("Delete", func() {
			var pgdb *mocks.IPostgreDB
			var vars map[string]string
			var deviceService *DeviceService
			var router *mux.Router
			BeforeEach(func() {
				pgdb = &mocks.IPostgreDB{}
				vars = make(map[string]string)
				vars["description"] = ""
				vars["name"] = ""
				deviceService = &DeviceService{Pgdb: pgdb}
				router = mux.NewRouter()
				router.HandleFunc("/device/{id}", deviceService.Delete).Methods("DELETE")
				pgdb.On("GetPostgreInstance").Return(pgdb)
			})
			It("given no error while deleting one device, return statusOK", func() {
				pgdb.On("DeleteDevice", "2").Return(nil)
				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("DELETE", "http://localhost:8080/device/2", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
			})
			It("given error while deleting the device, return StatusBadRequest", func() {
				pgdb.On("DeleteDevice", "2").Return(errors.New(""))

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("DELETE", "http://localhost:8080/device/2", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusBadRequest))
			})
		})
	})

	Context("SensorsServices", func() {
		Context("GetAll", func() {
			var pgdb *mocks.IPostgreDB
			var router *mux.Router
			BeforeEach(func() {
				pgdb = &mocks.IPostgreDB{}
				pgdb.On("GetPostgreInstance").Return(pgdb)
				deviceService := &SensorService{Pgdb: pgdb}
				router = mux.NewRouter()
				router.HandleFunc("/sensor", deviceService.GetAll).Methods("GET")
			})
			AfterEach(func() {
				writeResponse = OutputResponse
			})
			It("given no error while querying all the sensors, return them", func() {
				var result interface{}
				sensors := []models.Sensors{*&models.Sensors{}}
				writeResponse = func(writer http.ResponseWriter, data interface{}) error {
					result = data
					return nil
				}

				pgdb.On("GetAllSensors").Return(sensors, nil)

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "http://localhost:8080/sensor", nil)

				router.ServeHTTP(respRec, req)

				Expect(result).ToNot(BeEmpty())
				Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
			})
			It("given error while querying all the sensors, return StatusBadRequest", func() {
				sensors := []models.Sensors{}

				pgdb.On("GetAllSensors").Return(sensors, errors.New(""))

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "http://localhost:8080/sensor", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusBadRequest))
			})
		})

		Context("GetOne", func() {
			var pgdb *mocks.IPostgreDB
			var router *mux.Router
			BeforeEach(func() {
				pgdb = &mocks.IPostgreDB{}
				pgdb.On("GetPostgreInstance").Return(pgdb)
				sensorService := &SensorService{Pgdb: pgdb}
				router = mux.NewRouter()
				router.HandleFunc("/sensor/{id}", sensorService.GetOne).Methods("GET")
			})
			AfterEach(func() {
				writeResponse = OutputResponse
			})
			It("given no error while querying one of the sensors, return StatusOK", func() {
				var result interface{}
				sensor := models.Sensors{}
				writeResponse = func(writer http.ResponseWriter, data interface{}) error {
					result = data
					return nil
				}

				pgdb.On("GetSensor", "1").Return(sensor, nil)

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "http://localhost:8080/sensor/1", nil)

				router.ServeHTTP(respRec, req)

				Expect(result).ToNot(BeNil())
				Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
			})
			It("given error while querying the sensor, return the StatusBadRequest", func() {
				sensor := models.Sensors{}
				pgdb.On("GetSensor", "1").Return(sensor, errors.New(""))
				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "http://localhost:8080/sensor/1", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusBadRequest))
			})
		})

		Context("Create", func() {
			var pgdb *mocks.IPostgreDB
			var vars map[string]string
			var sensorService *SensorService
			var router *mux.Router
			BeforeEach(func() {
				pgdb = &mocks.IPostgreDB{}
				vars = make(map[string]string)
				vars["description"] = ""
				vars["name"] = ""
				vars["deviceid"] = "1"
				vars["unit"] = ""
				sensorService = &SensorService{Pgdb: pgdb}
				router = mux.NewRouter()
				router.HandleFunc("/sensor", sensorService.Create).Methods("POST")
				pgdb.On("GetPostgreInstance").Return(pgdb)
			})
			AfterEach(func() {
				writeResponse = OutputResponse
			})
			It("given no error while creating one sensor, return statusOK", func() {
				pgdb.On("CreateSensor", vars).Return(nil)

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "http://localhost:8080/sensor?deviceid=1", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
			})
			It("given error while creating the sensor, return StatusBadRequest", func() {
				pgdb.On("CreateSensor", vars).Return(errors.New(""))

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "http://localhost:8080/sensor?deviceid=1", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusBadRequest))
			})
		})

		Context("Update", func() {
			var pgdb *mocks.IPostgreDB
			var vars map[string]string
			var sensorService *SensorService
			var router *mux.Router
			BeforeEach(func() {
				pgdb = &mocks.IPostgreDB{}
				vars = make(map[string]string)
				vars["description"] = ""
				vars["name"] = ""
				vars["deviceid"] = ""
				vars["unit"] = ""
				sensorService = &SensorService{Pgdb: pgdb}
				router = mux.NewRouter()
				router.HandleFunc("/sensor", sensorService.Update).Methods("PUT")
				pgdb.On("GetPostgreInstance").Return(pgdb)
			})
			AfterEach(func() {
				writeResponse = OutputResponse
			})
			It("given no error while updating one sensor, return statusOK", func() {
				vars["id"] = "1"
				pgdb.On("UpdateSensor", vars).Return(nil)
				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("PUT", "http://localhost:8080/sensor?id=1", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
			})
			It("given error while updating the sensor, return StatusBadRequest", func() {
				vars["id"] = ""
				pgdb.On("UpdateSensor", vars).Return(errors.New(""))

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("PUT", "http://localhost:8080/sensor", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusBadRequest))
			})
		})

		Context("Delete", func() {
			var pgdb *mocks.IPostgreDB
			var sensorService *SensorService
			var router *mux.Router
			BeforeEach(func() {
				pgdb = &mocks.IPostgreDB{}
				sensorService = &SensorService{Pgdb: pgdb}
				router = mux.NewRouter()
				router.HandleFunc("/sensor/{id}", sensorService.Delete).Methods("DELETE")
				pgdb.On("GetPostgreInstance").Return(pgdb)
			})
			It("given no error while deleting one sensor, return statusOK", func() {
				pgdb.On("DeleteSensor", "2").Return(nil)
				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("DELETE", "http://localhost:8080/sensor/2", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
			})
			It("given error while deleting the sensor, return StatusBadRequest", func() {
				pgdb.On("DeleteSensor", "2").Return(errors.New(""))

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("DELETE", "http://localhost:8080/sensor/2", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusBadRequest))
			})
		})
	})

	Context("MeasurementsServices", func() {
		Context("GetAll", func() {
			var pgdb *mocks.IPostgreDB
			var influxdb *mocks.IInfluxDB
			var router *mux.Router
			BeforeEach(func() {
				pgdb = &mocks.IPostgreDB{}
				influxdb = &mocks.IInfluxDB{}
				pgdb.On("GetPostgreInstance").Return(pgdb)
				influxdb.On("GetInfluxInstance").Return(influxdb)
				measurementService := &MeasurementService{Pgdb: pgdb, Influxdb: influxdb}
				router = mux.NewRouter()
				router.HandleFunc("/measurement", measurementService.GetAll).Methods("GET")
			})
			AfterEach(func() {
				writeResponse = OutputResponse
			})
			It("given no error while querying all the measurements, return them", func() {
				var result interface{}
				measurement := &models.Measurements{}
				measurements := []*models.Measurements{measurement}
				writeResponse = func(writer http.ResponseWriter, data interface{}) error {
					result = data
					return nil
				}
				influxdb.On("GetAllMeasurements", context).Return(measurements, nil)

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "http://localhost:8080/measurement", nil)

				router.ServeHTTP(respRec, req)

				Expect(result).ToNot(BeEmpty())
				Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
			})
			It("given error while querying all the measurements, return StatusBadRequest", func() {
				measurements := []*models.Measurements{}
				influxdb.On("GetAllMeasurements", context).Return(measurements, errors.New(""))

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "http://localhost:8080/measurement", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusBadRequest))
			})
		})

		Context("GetSensorAverage", func() {
			var pgdb *mocks.IPostgreDB
			var influxdb *mocks.IInfluxDB
			var vars map[string]string
			var measurementService *MeasurementService
			var router *mux.Router
			BeforeEach(func() {
				pgdb = &mocks.IPostgreDB{}
				influxdb = &mocks.IInfluxDB{}
				pgdb.On("GetPostgreInstance").Return(pgdb)
				influxdb.On("GetInfluxInstance").Return(influxdb)
				measurementService = &MeasurementService{Pgdb: pgdb, Influxdb: influxdb}
				router = mux.NewRouter()
				router.HandleFunc("/sensorAverageValue", measurementService.GetSensorAverage).Methods("GET")

				vars = make(map[string]string)
				vars["deviceid"] = "1"
				vars["sensorid"] = "1"
				vars["startTime"] = "1"
				vars["endTime"] = "1"
			})
			AfterEach(func() {
				writeResponse = OutputResponse
			})
			It("given no error while querying the average measurement of specific sensor, return it", func() {
				averageM := models.AverageMeasurement{Average: "100"}
				influxdb.On("GetSensorAverage", context, vars).Return(averageM, nil)
				var result interface{}
				writeResponse = func(writer http.ResponseWriter, data interface{}) error {
					result = data
					return nil
				}

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "http://localhost:8080/sensorAverageValue?deviceid=1&sensorid=1&startTime=1&endTime=1", nil)

				router.ServeHTTP(respRec, req)

				Expect(result).ToNot(BeNil())
				Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
			})
			It("given error while querying the average measurement of specific sensor, return StatusBadRequest", func() {
				averageM := models.AverageMeasurement{}
				influxdb.On("GetSensorAverage", context, vars).Return(averageM, errors.New(""))

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "http://localhost:8080/sensorAverageValue?deviceid=1&sensorid=1&startTime=1&endTime=1", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusBadRequest))
			})
		})

		Context("SensorCorrelation", func() {
			influxdb := &mocks.IInfluxDB{}
			measurementService := &MeasurementService{Influxdb: influxdb}
			influxdb.On("GetInfluxInstance").Return(influxdb)
			router := mux.NewRouter()
			vars := make(map[string]string)
			vars["deviceid1"] = "1"
			vars["sensorid1"] = "1"
			vars["deviceid2"] = "1"
			vars["sensorid2"] = "1"
			vars["startTime"] = "1"
			vars["endTime"] = "1"
			router.HandleFunc("/sensorsCorrelationCoefficient", measurementService.SensorCorrelation).Methods("GET")
			It("given no error while retrieving the Pearsons's coeffiecient then return result", func() {
				influxdb.On("GetPearsonsCoefficient", context, vars).Return(float64(0), nil)

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "http://localhost:8080/sensorsCorrelationCoefficient?deviceid1=1&sensorid1=1&deviceid2=1&sensorid2=1&startTime=1&endTime=1", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
			})
		})

		Context("GetOne", func() {
			var influxdb *mocks.IInfluxDB
			var router *mux.Router
			BeforeEach(func() {
				influxdb = &mocks.IInfluxDB{}
				influxdb.On("GetInfluxInstance").Return(influxdb)
				measurementService := &MeasurementService{Influxdb: influxdb}
				router = mux.NewRouter()
				router.HandleFunc("/measurement/{id}", measurementService.GetOne).Methods("GET")
			})
			AfterEach(func() {
				writeResponse = OutputResponse
			})

			It("given no error while querying one of the measurements, return StatusOK", func() {
				var result interface{}
				measurement := &models.Measurements{}
				measurements := []*models.Measurements{measurement}
				influxdb.On("GetMeasurementById", context, "").Return(measurements, nil)
				writeResponse = func(writer http.ResponseWriter, data interface{}) error {
					result = data
					return nil
				}

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "http://localhost:8080/measurement/1", nil)

				router.ServeHTTP(respRec, req)

				Expect(result).ToNot(BeNil())
				Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
			})
			It("given error while querying the measurement, return the StatusBadRequest", func() {
				measurements := []*models.Measurements{}
				influxdb.On("GetMeasurementById", context, "").Return(measurements, errors.New(""))
				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "http://localhost:8080/measurement/1", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusBadRequest))
			})
		})

		Context("Create", func() {
			var pgdb *mocks.IPostgreDB
			var influxdb *mocks.IInfluxDB
			var router *mux.Router
			var vars map[string]string
			BeforeEach(func() {
				pgdb = &mocks.IPostgreDB{}
				influxdb = &mocks.IInfluxDB{}
				vars = make(map[string]string)
				vars["measuredat"] = "2006-01-02 15:04:05"
				vars["sensorid"] = "1"
				vars["value"] = ""
				vars["deviceid"] = "1"
				pgdb.On("GetPostgreInstance").Return(pgdb)
				influxdb.On("GetInfluxInstance").Return(influxdb)
				measurementService := &MeasurementService{Pgdb: pgdb, Influxdb: influxdb}
				router = mux.NewRouter()
				router.HandleFunc("/measurement", measurementService.Create).Methods("POST")
			})
			It("given no error while creating one measurement, return statusOK", func() {
				device := models.Device{}
				sensor := models.Sensors{Id: 1}
				device.Sensors = []*models.Sensors{&sensor}
				pgdb.On("GetDevice", "1").Return(device, nil)
				influxdb.On("CreateMeasurement", vars).Return(nil)

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "http://localhost:8080/measurement?sensorid=1&deviceid=1&measuredat=2006-01-02%2015:04:05", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
			})
			It("given error while creating the measurement, return StatusBadRequest", func() {
				device := models.Device{}
				pgdb.On("GetDevice", "1").Return(device, errors.New(""))

				respRec := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "http://localhost:8080/measurement?sensorid=1&deviceid=1&measuredat=2006-01-02%2015:04:05", nil)

				router.ServeHTTP(respRec, req)

				Expect(respRec.Code).To(BeNumerically("==", http.StatusBadRequest))
			})
		})
	})
})
