package models

type Measurements struct {
	MeasuredAt string `json:"measuredat" db:"measuredat"`
	Device_id  int64  `json:"deviceid" db:"deviceid"`
	Sensor_id  int64  `json:"sensorid" db:"sensorid"`
	Value      string `json:"value" db:"value"`
}

type AverageMeasurement struct {
	Device_id int64  `json:"deviceid"`
	Sensor_id int64  `json:"sensorid"`
	Average   string `json:"average"`
}

type SensorsCorrelation struct {
	Device_id1  int     `json:"deviceid1"`
	Sensor_id1  int     `json:"sensorid1"`
	Device_id2  int     `json:"deviceid2"`
	Sensor_id2  int     `json:"sensorid2"`
	Correlation float64 `json:"correlation"`
}
