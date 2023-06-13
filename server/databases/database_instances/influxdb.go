package database_instances

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/altnum/sensorapp/models"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"gonum.org/v1/gonum/stat"
)

const (
	inHost   = "localhost"
	inPort   = "8086"
	myOrg    = "sensor-org"
	myBucket = "sensors-bucket"
)

var parseInt = strconv.ParseInt

type IInfluxDB interface {
	IDB
	GetAllMeasurements(context.Context) ([]*models.Measurements, error)
	GetSensorAverage(context.Context, map[string]string) (models.AverageMeasurement, error)
	GetPearsonsCoefficient(context.Context, map[string]string) (float64, error)
	GetMeasurementById(context.Context, string) ([]*models.Measurements, error)
	CreateMeasurement(map[string]string) error
	GetDB() influxdb2.Client
}

type InfluxDB struct {
	WriteAPI api.WriteAPI
	DB       influxdb2.Client
	QueryApi api.QueryAPI
}

func (i *InfluxDB) GetPostgreInstance() IPostgreDB {
	return nil
}

func (i *InfluxDB) GetInfluxInstance() IInfluxDB {
	return i
}

func (i *InfluxDB) Open(context context.Context) error {
	i.DB = influxdb2.NewClient("http://"+inHost+":"+inPort, "mytoken")
	i.WriteAPI = i.DB.WriteAPI(myOrg, myBucket)
	i.QueryApi = i.DB.QueryAPI(myOrg)
	health, err := i.DB.Health(context)
	if err != nil {
		return err
	}
	if fmt.Sprint(health.Status) != "pass" {
		return errors.New(fmt.Sprint(health.Message))
	}

	return nil
}

func (i *InfluxDB) Close() error {
	i.DB.Close()
	return nil
}

func (i *InfluxDB) GetAllMeasurements(context context.Context) ([]*models.Measurements, error) {
	query := `from(bucket:"` + myBucket + `")|> range(start: -24h)`

	measures, err := i.GetMeasuresFromQuery(context, query)
	if err != nil {
		return nil, err
	}

	return measures, nil
}

func (i *InfluxDB) GetMeasuresFromQuery(context context.Context, query string) ([]*models.Measurements, error) {
	result, err := i.QueryApi.Query(context, query)
	if err != nil {
		return nil, err
	}

	measures := []*models.Measurements{}

	for result.Next() {
		m := models.Measurements{}
		m.Sensor_id, err = parseInt(fmt.Sprint(result.Record().Measurement()), 0, 64)
		if err != nil {
			return nil, err
		}

		timeM := result.Record().Time().Add(3 * time.Hour)
		m.MeasuredAt = fmt.Sprint(timeM.Format("2006-01-02 15:04:05"))
		m.Device_id, err = parseInt(fmt.Sprint(result.Record().ValueByKey("deviceid")), 0, 64)
		if err != nil {
			return nil, err
		}

		if result.Record().Field() == "value" {
			m.Value = fmt.Sprint(result.Record().Value())
		}

		measures = append(measures, &m)
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	if len(measures) == 0 {
		return nil, errors.New("No such measurements with the specified criteria.")
	}

	return measures, nil
}

func (i *InfluxDB) GetSensorAverage(context context.Context, vars map[string]string) (models.AverageMeasurement, error) {
	deviceid := vars["deviceid"]
	sensorid := vars["sensorid"]
	start := vars["startTime"]
	end := vars["endTime"]
	var averageMeasurement models.AverageMeasurement

	query, err := i.GetMeasurementsQueryInSpecificTimeFrame(start, end, deviceid, sensorid)
	if err != nil {
		return averageMeasurement, err
	}

	measures, err := i.GetMeasuresFromQuery(context, query)
	if err != nil {
		return averageMeasurement, err
	}

	if len(measures) <= 0 {
		return averageMeasurement, errors.New("No measurements with the specified criteria.")
	}

	averageMeasurement, err = i.CalculateAverageValue(measures)
	if err != nil {
		return averageMeasurement, err
	}

	return averageMeasurement, nil
}

func (i *InfluxDB) GetMeasurementsQueryInSpecificTimeFrame(start string, end string, deviceid string, sensorid string) (string, error) {
	layout := "2006-01-02 15:04:05"
	timeStart, err := time.Parse(layout, start)
	if err != nil {
		return "", err
	}

	timeStart = timeStart.Add(-3 * time.Hour)
	startSecondsUnix := fmt.Sprint(timeStart.Unix())

	timeEnd, err := time.Parse(layout, end)
	if err != nil {
		return "", err
	}

	timeEnd = timeEnd.Add(-3 * time.Hour)
	endSecondsUnix := fmt.Sprint(timeEnd.Unix())

	if startSecondsUnix >= endSecondsUnix {
		return "", errors.New("Not valid timestamps.")
	}

	query := `from(bucket:"` + myBucket + `")|> range(start: ` + startSecondsUnix + `, stop: ` + endSecondsUnix + `) |> filter(fn: (r) => r._measurement == "` + sensorid + `" and r.deviceid == "` + deviceid + `")`

	return query, nil
}

func (i *InfluxDB) CalculateAverageValue(measures []*models.Measurements) (models.AverageMeasurement, error) {
	var averageMeasurement models.AverageMeasurement
	var sum float64
	for _, measure := range measures {
		value, err := strconv.ParseFloat(measure.Value, 64)
		if err != nil {
			return averageMeasurement, err
		}
		sum += value
	}
	measuresLength, err := strconv.ParseFloat(fmt.Sprint(len(measures)), 64)
	if err != nil {
		return averageMeasurement, err
	}
	averageSum := sum / measuresLength
	averageMeasurement = models.AverageMeasurement{Device_id: measures[0].Device_id, Sensor_id: measures[0].Sensor_id, Average: fmt.Sprint(averageSum)}

	return averageMeasurement, nil
}

func (i *InfluxDB) GetPearsonsCoefficient(context context.Context, args map[string]string) (float64, error) {
	query1, err := i.GetMeasurementsQueryInSpecificTimeFrame(args["startTime"], args["endTime"], args["deviceid1"], args["sensorid1"])
	if err != nil {
		return 0, err
	}

	query2, err := i.GetMeasurementsQueryInSpecificTimeFrame(args["startTime"], args["endTime"], args["deviceid2"], args["sensorid2"])
	if err != nil {
		return 0, err
	}

	measures1, err := i.GetMeasuresFromQuery(context, query1)
	if err != nil {
		return 0, err
	}

	measures2, err := i.GetMeasuresFromQuery(context, query2)
	if err != nil {
		return 0, err
	}

	measures1Values, measures2Values, err := i.GetValuesDataFromMeasurements(measures1, measures2)
	if err != nil {
		return 0, err
	}

	covariance, err := i.Covariance(measures1Values, measures2Values)
	if err != nil {
		return 0, err
	}

	coefficient := i.CalculatePearsonsCoefficient(covariance, measures1Values, measures2Values)
	return coefficient, nil
}

func (i *InfluxDB) GetValuesDataFromMeasurements(data1 []*models.Measurements, data2 []*models.Measurements) ([]float64, []float64, error) {
	minLength := MinLength(data1, data2)

	var measures1Values []float64
	for i := 0; i < minLength; i++ {
		m1Value, err := strconv.ParseFloat(data1[i].Value, 64)
		if err != nil {
			return nil, nil, err
		}

		measures1Values = append(measures1Values, m1Value)
	}

	var measures2Values []float64
	for i := 0; i < minLength; i++ {
		m2Value, err := strconv.ParseFloat(data2[i].Value, 64)
		if err != nil {
			return nil, nil, err
		}

		measures2Values = append(measures2Values, m2Value)
	}

	return measures1Values, measures2Values, nil
}

func MinLength(slice1 []*models.Measurements, slice2 []*models.Measurements) int {
	if len(slice1) < len(slice2) {
		return len(slice1)
	}

	return len(slice2)
}

func (i *InfluxDB) Covariance(data1 []float64, data2 []float64) (float64, error) {
	if len(data1) != len(data2) {
		return 0, errors.New("Data slices with different length")
	}

	var data1Cov []float64
	for x := 0; x < len(data1); x++ {
		data1Cov = append(data1Cov, data1[x]-stat.Mean(data1, nil))
	}

	var data2Cov []float64
	for y := 0; y < len(data1); y++ {
		data2Cov = append(data2Cov, data2[y]-stat.Mean(data2, nil))
	}

	var sumOfCov float64
	for i := 0; i < len(data1); i++ {
		sumOfCov += (data1Cov[i] * data2Cov[i])
	}

	data1Length, err := strconv.ParseFloat(fmt.Sprint(len(data1)), 64)
	if err != nil {
		return 0, err
	}

	return sumOfCov / data1Length, nil
}

func (i *InfluxDB) CalculatePearsonsCoefficient(covariance float64, data1 []float64, data2 []float64) float64 {
	if covariance == 0 {
		return 0
	}

	return covariance / (stat.StdDev(data1, nil) * stat.StdDev(data2, nil))
}

func (i *InfluxDB) GetMeasurementById(context context.Context, sensorid string) ([]*models.Measurements, error) {
	query := `from(bucket:"` + myBucket + `")|> range(start: -24h) |> filter(fn: (r) => r._measurement == "` + sensorid + `")`

	measures, err := i.GetMeasuresFromQuery(context, query)
	if err != nil {
		return nil, err
	}

	return measures, nil
}

func (i *InfluxDB) CreateMeasurement(vars map[string]string) error {
	measuredAt := vars["measuredat"]
	sensorid := vars["sensorid"]
	value := vars["value"]
	deviceid := vars["deviceid"]
	layout := "2006-01-02 15:04:05"
	timeM, err := time.Parse(layout, measuredAt)
	if err != nil {
		return err
	}

	timeM = timeM.Local().Add(-3 * time.Hour)

	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}

	point := influxdb2.NewPointWithMeasurement(sensorid).
		AddField("value", f).
		AddTag("deviceid", deviceid).
		SetTime(timeM)

	i.WriteAPI.WritePoint(point)
	i.WriteAPI.Flush()

	return nil
}

func (i *InfluxDB) GetDB() influxdb2.Client {
	return i.DB
}
