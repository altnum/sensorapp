package database_instances

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/altnum/sensorapp/models"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testing different databases", func() {
	context := context.Background()
	Context("PostgreSQL", func() {
		Context("Open", func() {
			var pgdb IDB
			BeforeEach(func() {
				pgdb = &PostgreDB{}
				postgreConnect = sqlx.Connect
			})
			It("given error when trying to establish connection with postgreSQL, return error", func() {
				postgreConnect = func(driverName string, dataSourceName string) (*sqlx.DB, error) {
					return nil, errors.New("")
				}

				err := pgdb.Open(context)

				Expect(err).ToNot(BeNil())
			})
			It("given no error when trying to establish connection with postgreSQL return nil", func() {
				err := pgdb.Open(context)

				Expect(err).To(BeNil())
			})
		})

		Context("GetAllDevices", func() {
			pgdb := &PostgreDB{}
			pgdb.Open(context)
			It("given no error when querying all devices return slice of them", func() {
				devices, err := pgdb.GetAllDevices()

				Expect(err).To(BeNil())
				Expect(devices).ToNot(BeEmpty())
			})
		})

		Context("FindSensors", func() {
			pgdb := &PostgreDB{}
			pgdb.Open(context)
			It("given no error when querying all sensors of a device attach slice of them", func() {
				device := &models.Device{Id: 1}
				err := pgdb.FindSensors(device)

				Expect(err).To(BeNil())
				Expect(device.Sensors).ToNot(BeEmpty())
			})
			It("given invalid device when querying all sensors of a device return empty slice", func() {
				device := &models.Device{Id: 0}

				err := pgdb.FindSensors(device)

				Expect(err).To(BeNil())
				Expect(device.Sensors).To(BeEmpty())
			})
		})

		Context("GetDevice", func() {
			pgdb := &PostgreDB{}
			pgdb.Open(context)
			It("given no error when querying a device by id return it", func() {
				device, err := pgdb.GetDevice("1")

				Expect(err).To(BeNil())
				Expect(device.Id).To(BeNumerically("==", 1))
			})
			It("given invalid id when querying a device by id return err", func() {
				device, err := pgdb.GetDevice("-1")

				Expect(err).ToNot(BeNil())
				Expect(device.Id).To(BeNumerically("==", 0))
			})
		})

		Context("CreateDevice", func() {
			pgdb := &PostgreDB{}
			vars := make(map[string]string)
			BeforeEach(func() {
				pgdb.Open(context)
				vars["name"] = ""
				vars["description"] = ""
			})
			It("given no established connection when inserting into the database, then return error", func() {
				pgdb.Close()
				err := pgdb.CreateDevice(vars)

				Expect(err).ToNot(BeNil())
			})
			It("given no error when inserting into the database, then return nil", func() {
				err := pgdb.CreateDevice(vars)

				Expect(err).To(BeNil())
			})
		})

		Context("UpdateDevice", func() {
			pgdb := &PostgreDB{}
			vars := make(map[string]string)
			BeforeEach(func() {
				pgdb.Open(context)
				vars["name"] = ""
				vars["description"] = ""
				vars["id"] = "2"
			})
			It("given no established connection while updating the device, then return error", func() {
				pgdb.Close()
				err := pgdb.UpdateDevice(vars)

				Expect(err).ToNot(BeNil())
			})
			It("given no error while updating the device, then return nil", func() {
				err := pgdb.UpdateDevice(vars)

				Expect(err).To(BeNil())
			})
		})

		Context("DeleteDevice", func() {
			pgdb := &PostgreDB{}
			BeforeEach(func() {
				pgdb.Open(context)
			})
			It("given no established connection while deletion, then return error", func() {
				pgdb.Close()
				err := pgdb.DeleteDevice("1")

				Expect(err).ToNot(BeNil())
			})
			It("given no error while deletion, then return nil", func() {
				err := pgdb.DeleteDevice("-1")

				Expect(err).To(BeNil())
			})
		})

		Context("GetAllSensors", func() {
			pgdb := &PostgreDB{}
			pgdb.Open(context)
			It("given no error when querying all sensors return slice of them", func() {
				sensors, err := pgdb.GetAllSensors()

				Expect(err).To(BeNil())
				Expect(sensors).ToNot(BeEmpty())
			})
		})

		Context("GetSensor", func() {
			pgdb := &PostgreDB{}
			pgdb.Open(context)
			It("given no error when querying a sensor by id return it", func() {
				device, err := pgdb.GetDevice("1")

				Expect(err).To(BeNil())
				Expect(device.Id).To(BeNumerically("==", 1))
			})
			It("given invalid id when querying a device by id return err", func() {
				device, err := pgdb.GetDevice("-1")

				Expect(err).ToNot(BeNil())
				Expect(device.Id).To(BeNumerically("==", 0))
			})
		})

		Context("CreateSensor", func() {
			pgdb := &PostgreDB{}
			vars := make(map[string]string)
			BeforeEach(func() {
				pgdb.Open(context)
				vars["deviceid"] = "1"
				vars["name"] = ""
				vars["description"] = ""
				vars["unit"] = ""
			})
			It("given no established connection when inserting into the database, then return error", func() {
				pgdb.Close()
				err := pgdb.CreateSensor(vars)

				Expect(err).ToNot(BeNil())
			})
			It("given invalid deviceid when inserting into the database, then return err", func() {
				vars["deviceid"] = "-1"
				err := pgdb.CreateSensor(vars)

				Expect(err).ToNot(BeNil())
			})
			It("given no error when inserting into the database, then return nil", func() {
				err := pgdb.CreateSensor(vars)

				Expect(err).To(BeNil())
			})
		})

		Context("UpdateSensor", func() {
			pgdb := &PostgreDB{}
			vars := make(map[string]string)
			BeforeEach(func() {
				pgdb.Open(context)
				vars["deviceid"] = "1"
				vars["name"] = ""
				vars["description"] = ""
				vars["unit"] = ""
				vars["sensorid"] = "9"
			})
			It("given no established connection while updating the device, then return error", func() {
				pgdb.Close()
				err := pgdb.UpdateSensor(vars)

				Expect(err).ToNot(BeNil())
			})
			It("given invalid deviceid when updating the database, then return err", func() {
				err := pgdb.UpdateSensor(vars)

				Expect(err).ToNot(BeNil())
			})

		})

		Context("DeleteSensor", func() {
			pgdb := &PostgreDB{}
			BeforeEach(func() {
				pgdb.Open(context)
			})
			It("given no established connection while deletion, then return error", func() {
				pgdb.Close()
				err := pgdb.DeleteSensor("9")

				Expect(err).ToNot(BeNil())
			})
			It("given no error while deletion, then return nil", func() {
				err := pgdb.DeleteSensor("-1")

				Expect(err).To(BeNil())
			})
		})
	})

	Context("InfluxDB", func() {
		influxdb := &InfluxDB{}
		influxdb.Open(context)

		Context("GetAllMeasurements", func() {
			time := time.Now()
			timeNow := fmt.Sprint(time.Format("2006-01-02 15:04:05"))
			vars := map[string]string{"measuredat": timeNow, "sensorid": "1", "value": "", "deviceid": "1"}
			influxdb.CreateMeasurement(vars)
			AfterEach(func() {
				parseInt = strconv.ParseInt
			})
			It("given no error when querying all measurements then return slice of them", func() {
				measurements, err := influxdb.GetAllMeasurements(context)

				Expect(err).To(BeNil())
				Expect(measurements).ToNot(BeEmpty())
			})
			It("given error while quering then return error", func() {
				parseInt = func(s string, base, bitSize int) (i int64, err error) {
					return 0, errors.New("")
				}
			})
		})

		Context("GetSensorAverage", func() {
			timeMeasure := time.Now()
			timeNow := fmt.Sprint(timeMeasure.Format("2006-01-02 15:04:05"))
			vars := map[string]string{"measuredat": timeNow, "sensorid": "1", "value": "", "deviceid": "1"}
			influxdb.CreateMeasurement(vars)
			timeStart := timeMeasure.Add(-1 * time.Minute)
			timeEnd := timeMeasure.Add(1 * time.Minute)
			vars["startTime"] = fmt.Sprint(timeStart.Format("2006-01-02 15:04:05"))
			vars["endTime"] = fmt.Sprint(timeEnd.Format("2006-01-02 15:04:05"))

			It("given no error when getting the average, return the result", func() {
				averageM, err := influxdb.GetSensorAverage(context, vars)

				Expect(err).To(BeNil())
				Expect(averageM).ToNot(BeNil())
			})
		})

		Context("GetPearsonsCoefficient", func() {
			It("given data with positive correlation when calculating the coefficient then return number between 0 and 1", func() {
				slice1 := []float64{15.5, 13.6, 13.5, 13.0, 13.3, 12.4, 11.1, 13.1, 16.1, 16.4, 13.4, 13.2, 14.3, 16.1}
				slice2 := []float64{0.450, 0.420, 0.440, 0.395, 0.395, 0.370, 0.390, 0.400, 0.445, 0.470, 0.390, 0.400, 0.420, 0.450}

				covariance, err := influxdb.Covariance(slice1, slice2)
				result := influxdb.CalculatePearsonsCoefficient(covariance, slice1, slice2)

				Expect(err).To(BeNil())
				Expect(result).To(BeNumerically(">", 0))
				Expect(result).To(BeNumerically("<=", 1))

			})
			It("given data with negative correlation when calculating the coefficient then return number between -1 and 0", func() {
				slice1 := []float64{15.5, 13.6, 13.5, 13.0, 13.3, 12.4, 11.1, 13.1, 16.1, 16.4, 13.4, 13.2, 14.3, 16.1}
				slice2 := []float64{-0.450, -0.420, -0.440, -0.395, -0.395, -0.370, -0.390, -0.400, -0.445, -0.470, -0.390, -0.400, -0.420, -0.450}

				covariance, err := influxdb.Covariance(slice1, slice2)
				result := influxdb.CalculatePearsonsCoefficient(covariance, slice1, slice2)

				Expect(err).To(BeNil())
				Expect(result).To(BeNumerically("<", 0))
				Expect(result).To(BeNumerically(">=", -1))
			})
			It("given data with positive correlation when calculating the coefficient then return number between 0 and 1", func() {
				slice1 := []float64{15.5, 15.5, 15.5}
				slice2 := []float64{15.5, 15.5, 15.5}

				covariance, err := influxdb.Covariance(slice1, slice2)
				result := influxdb.CalculatePearsonsCoefficient(covariance, slice1, slice2)

				Expect(err).To(BeNil())
				Expect(result).To(BeNumerically("==", 0))
			})
		})

		Context("GetMeasurementById", func() {
			It("given invalid id when querying measurement by id then return empty slice", func() {
				measures, err := influxdb.GetMeasurementById(context, "-3")

				Expect(err.Error()).To(BeEquivalentTo("No such measurements with the specified criteria."))
				Expect(measures).To(BeEmpty())
			})
			It("given correct id when querying measurement by id then return nil", func() {
				measures, err := influxdb.GetMeasurementById(context, "1")

				Expect(err).To(BeNil())
				Expect(measures).ToNot(BeEmpty())
			})
		})

		Context("GetPearsonsCoefficient", func() {
			timeMeasure1 := time.Now()
			timeNow1 := fmt.Sprint(timeMeasure1.Format("2006-01-02 15:04:05"))
			vars := map[string]string{"measuredat": timeNow1, "sensorid": "1", "value": "50", "deviceid": "1"}
			influxdb.CreateMeasurement(vars)
			timeMeasure2 := time.Now()
			timeNow2 := fmt.Sprint(timeMeasure2.Format("2006-01-02 15:04:05"))
			vars = map[string]string{"measuredat": timeNow2, "sensorid": "2", "value": "50", "deviceid": "1"}
			influxdb.CreateMeasurement(vars)
			timeStart := timeMeasure1.Add(-1 * time.Minute)
			timeEnd := timeMeasure2.Add(1 * time.Minute)
			vars = make(map[string]string)
			vars["startTime"] = fmt.Sprint(timeStart.Format("2006-01-02 15:04:05"))
			vars["endTime"] = fmt.Sprint(timeEnd.Format("2006-01-02 15:04:05"))
			vars["deviceid1"] = "1"
			vars["deviceid2"] = "1"
			vars["sensorid1"] = "1"
			vars["sensorid2"] = "2"
			It("given two data sets with identical values, return neutral correlation", func() {
				coefficient, err := influxdb.GetPearsonsCoefficient(context, vars)

				Expect(err).To(BeNil())
				Expect(coefficient).To(BeNumerically("==", 0))
			})
		})

		Context("GetMeasurementsQueryInSpecificTimeFrame", func() {
			timeMeasure := time.Now()
			timeNow := fmt.Sprint(timeMeasure.Format("2006-01-02 15:04:05"))
			vars := map[string]string{"measuredat": timeNow, "sensorid": "1", "value": "50", "deviceid": "1"}
			influxdb.CreateMeasurement(vars)
			timeStart := timeMeasure.Add(-1 * time.Minute)
			timeEnd := timeMeasure.Add(1 * time.Minute)
			It("given valid timeframes when searching between them, then return the right measurements", func() {
				query, err := influxdb.GetMeasurementsQueryInSpecificTimeFrame(fmt.Sprint(timeStart.Format("2006-01-02 15:04:05")), fmt.Sprint(timeEnd.Format("2006-01-02 15:04:05")), "1", "1")
				measures, err := influxdb.GetMeasuresFromQuery(context, query)

				Expect(err).To(BeNil())
				Expect(measures).ToNot(BeEmpty())

			})
		})

		Context("GetMeasuresFromQuery", func() {
			It("given invalid query API when querying then return error", func() {
				influxdb.QueryApi = influxdb.DB.QueryAPI("")

				query := `from(bucket:"sensors-bucket")|> range(start: -24h)`
				measurements, err := influxdb.GetMeasuresFromQuery(context, query)

				Expect(err).ToNot(BeNil())
				Expect(measurements).To(BeEmpty())
			})

		})

		Context("CreateMeasurement", func() {
			It("given incorrect time while creating the measurement, then return error", func() {
				vars := map[string]string{"measuredat": "", "sensorid": "", "value": "", "deviceid": ""}
				err := influxdb.CreateMeasurement(vars)

				Expect(err).ToNot(BeNil())
			})
		})

		Context("GetDB", func() {
			It("return DB instance", func() {
				i := influxdb.GetDB()

				Expect(i).ToNot(BeNil())
			})
		})
	})
})
