package models

type Sensors struct {
	Id          int    `json:"id" db:"id"`
	Device_id   int    `json:"deviceid" db:"deviceid"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Unit        string `json:"unit" db:"unit"`
}
