package routes

import (
	"context"
	"net/http"
	"net/http/httptest"

	. "github.com/altnum/sensorapp/databases/database_instances"
	"github.com/altnum/sensorapp/services"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testing routes addresses", func() {
	Context("Devices routes", func() {
		context := context.Background()
		PgHost = "localhost"
		pgdb := &PostgreDB{}
		pgdb.Open(context)
		router := mux.NewRouter()
		deviceService := &services.DeviceService{Pgdb: pgdb}
		DeviceRoutesInit(router, deviceService)

		It("GetAll", func() {
			respRec := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "http://localhost:8080/device", nil)

			router.ServeHTTP(respRec, req)

			Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
		})
		It("GetOne", func() {
			respRec := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "http://localhost:8080/device/1", nil)

			router.ServeHTTP(respRec, req)

			Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
		})
		It("Create", func() {
			respRec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "http://localhost:8080/device", nil)

			router.ServeHTTP(respRec, req)

			Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
		})
		It("Update", func() {
			respRec := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "http://localhost:8080/device?id=2", nil)

			router.ServeHTTP(respRec, req)

			Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
		})
		It("Delete", func() {
			respRec := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "http://localhost:8080/device/2", nil)

			router.ServeHTTP(respRec, req)

			Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
		})
		It("Given wrong route, when navigating, then return error", func() {
			respRec := httptest.NewRecorder()
			req := httptest.NewRequest("PATCH", "http://localhost:8080/devices", nil)

			router.ServeHTTP(respRec, req)

			Expect(respRec.Code).To(BeNumerically("==", http.StatusNotFound))
		})
	})

	Context("Sensors routes", func() {
		context := context.Background()
		pgdb := &PostgreDB{}
		pgdb.Open(context)
		router := mux.NewRouter()
		sensorsService := &services.SensorService{Pgdb: pgdb}
		SensorRoutesInit(router, sensorsService)

		It("GetAll", func() {
			respRec := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "http://localhost:8080/sensor", nil)

			router.ServeHTTP(respRec, req)

			Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
		})
		It("GetOne", func() {
			respRec := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "http://localhost:8080/sensor/1", nil)

			router.ServeHTTP(respRec, req)

			Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
		})
		It("Create", func() {
			respRec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "http://localhost:8080/sensor?deviceid=1", nil)

			router.ServeHTTP(respRec, req)

			Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
		})
		It("Update", func() {
			respRec := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "http://localhost:8080/sensor?id=1&deviceid=1", nil)

			router.ServeHTTP(respRec, req)

			Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
		})
		It("Delete", func() {
			respRec := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "http://localhost:8080/sensor/12", nil)

			router.ServeHTTP(respRec, req)

			Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
		})
		It("Given wrong route, when navigating, then return error", func() {
			respRec := httptest.NewRecorder()
			req := httptest.NewRequest("PATCH", "http://localhost:8080/sensors", nil)

			router.ServeHTTP(respRec, req)

			Expect(respRec.Code).To(BeNumerically("==", http.StatusNotFound))
		})
	})

	Context("Measurements routes", func() {
		context := context.Background()
		PgHost = "localhost"
		InHost = "localhost"
		pgdb := &PostgreDB{}
		influxdb := &InfluxDB{}
		pgdb.Open(context)
		influxdb.Open(context)
		router := mux.NewRouter()
		measurementsService := &services.MeasurementService{Pgdb: pgdb, Influxdb: influxdb}
		MeasurementRoutesInit(router, measurementsService)

		It("GetAll", func() {
			respRec := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "http://localhost:8080/measurement", nil)

			router.ServeHTTP(respRec, req)

			Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
		})
		It("GetOne", func() {
			respRec := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "http://localhost:8080/measurement/1", nil)

			router.ServeHTTP(respRec, req)

			Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
		})
		It("Create", func() {
			respRec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "http://localhost:8080/measurement?measuredat=2006-01-02%2015:04:05&sensorid=1&deviceid=1&value=0", nil)

			router.ServeHTTP(respRec, req)

			Expect(respRec.Code).To(BeNumerically("==", http.StatusOK))
		})
		It("Given wrong route, when navigating, then return error", func() {
			respRec := httptest.NewRecorder()
			req := httptest.NewRequest("PATCH", "http://localhost:8080/measurements", nil)

			router.ServeHTTP(respRec, req)

			Expect(respRec.Code).To(BeNumerically("==", http.StatusNotFound))
		})
	})

	Context("LoadRoutes", func() {
		router := mux.Router{}
		pgdb := &PostgreDB{}
		pgdb.Open(context.Background())
		service := &services.DeviceService{Pgdb: pgdb}
		loader := RouteLoader{}
		It("given specific service, when initializing the routes, then initialize just its routes", func() {
			loader.LoadRoutes(&router, service)

			respRec1 := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "http://localhost:8080/device", nil)

			router.ServeHTTP(respRec1, req)

			respRec2 := httptest.NewRecorder()
			req2 := httptest.NewRequest("GET", "http://localhost:8080/measurement", nil)

			router.ServeHTTP(respRec2, req2)

			Expect(respRec1.Code).To(BeNumerically("==", http.StatusOK))
			Expect(respRec2.Code).To(BeNumerically("==", http.StatusNotFound))
		})
	})
})
