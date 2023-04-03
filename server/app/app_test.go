package app

import (
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.wdf.sap.corp/I554249/sensor/databases/database_instances"
	mocksIDB "github.wdf.sap.corp/I554249/sensor/databases/database_instances/mocks"
	"github.wdf.sap.corp/I554249/sensor/databases/mocks"
	"github.wdf.sap.corp/I554249/sensor/services"
	mocksServiceFactory "github.wdf.sap.corp/I554249/sensor/services/mocks"
)

var _ = Describe("Testing initial app methods", func() {
	Context("ConnectToDBs", func() {
		context := context.Background()
		var mockDBFactory mocks.IDatabaseFactory
		BeforeEach(func() {
			mockDBFactory = mocks.IDatabaseFactory{}

		})
		It("given error while connecting to the databases, then return err and empty slice of them", func() {
			sliceDbs := []IDB{&PostgreDB{}, &InfluxDB{}}
			mockDBFactory.On("RetrieveDatabases").Return(sliceDbs)
			mockDBFactory.On("ConnectDatabases", context, sliceDbs).Return(errors.New(""))
			dbs, err := connectToDBs(context, &mockDBFactory)

			Expect(err).ToNot(BeNil())
			Expect(dbs).To(BeEmpty())
		})
		It("given no error while connecting to the databases, then return nil and slice of them", func() {
			sliceDbs := []IDB{&PostgreDB{}, &InfluxDB{}}
			mockDBFactory.On("RetrieveDatabases").Return(sliceDbs)
			mockDBFactory.On("ConnectDatabases", context, sliceDbs).Return(nil)
			dbs, err := connectToDBs(context, &mockDBFactory)

			Expect(err).To(BeNil())
			Expect(dbs).ToNot(BeEmpty())
		})
	})

	Context("Run", func() {
		context := context.Background()
		It("given error when connectToDBs(), then return error", func() {
			sliceDbs := []IDB{&PostgreDB{}, &InfluxDB{}}
			mockDBFactory := mocks.IDatabaseFactory{}
			mockDBFactory.On("RetrieveDatabases").Return(sliceDbs)
			mockDBFactory.On("ConnectDatabases", context, sliceDbs).Return(errors.New(""))

			dbs, err := connectToDBs(context, &mockDBFactory)

			Expect(err).ToNot(BeNil())
			Expect(dbs).To(BeEmpty())
		})
		It("given error when closeDBs() then return error", func() {
			db := &mocksIDB.IDB{}
			dbs := []IDB{db}
			db.On("Close").Return(errors.New(""))

			err := closeDBs(dbs)

			Expect(err).ToNot(BeNil())
		})
		It("given no errors when initializeRoutes(), create all instances successfuly", func() {
			router := mux.NewRouter()
			dbs := []IDB{&PostgreDB{}, &InfluxDB{}}
			serviceFactory := &mocksServiceFactory.IServiceFactory{}
			serviceFactory.On("CreateService", "measurement", dbs).Return(&services.MeasurementService{})
			serviceFactory.On("CreateService", "device", dbs).Return(&services.DeviceService{})
			serviceFactory.On("CreateService", "sensor", dbs).Return(&services.SensorService{})

			initializeRoutes(router, dbs, serviceFactory)
		})
		It("given error when startServer() then return error", func() {
			startServer = func(addr string, handler http.Handler) error {
				return errors.New("")
			}

			err := Run()

			Expect(err).ToNot(BeNil())
		})
	})
})
