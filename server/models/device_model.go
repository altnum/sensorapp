package models

type Device struct {
	Id          int        `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Description string     `json:"description" db:"description"`
	Sensors     []*Sensors `json:"sensors"`
}
