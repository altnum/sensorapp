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
