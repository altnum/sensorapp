package database_instances

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.wdf.sap.corp/I554249/sensor/models"
)

const (
	pgHost           = "localhost"
	pgPort           = "5432"
	pgUser           = "postgres"
	pgPassword       = "1234"
	pgDbname         = "postgres"
	columnsWithAlias = true
)

var postgreConnect = sqlx.Connect

type IPostgreDB interface {
	IDB
	GetAllDevices() ([]models.Device, error)
	GetDevice(string) (models.Device, error)
	CreateDevice(map[string]string) error
	UpdateDevice(map[string]string) error
	DeleteDevice(string) error
	GetAllSensors() ([]models.Sensors, error)
	GetSensor(string) (models.Sensors, error)
	CreateSensor(map[string]string) error
	UpdateSensor(map[string]string) error
	DeleteSensor(string) error
}

type PostgreDB struct {
	DB sqlx.DB
}

func (p *PostgreDB) GetPostgreInstance() IPostgreDB {
	return p
}

func (p *PostgreDB) GetInfluxInstance() IInfluxDB {
	return nil
}

func (p *PostgreDB) Open(context context.Context) error {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pgHost, pgPort, pgUser, pgPassword, pgDbname)
	pgdb, err := postgreConnect("postgres", psqlconn)
	if err != nil {
		return err
	}

	p.DB = *pgdb

	return nil
}

func (p *PostgreDB) Close() error {
	return p.DB.Close()
}

func (p *PostgreDB) GetAllDevices() ([]models.Device, error) {
	query := "SELECT * FROM device"
	rows, err := p.DB.Queryx(query)
	if err != nil {
		return nil, err
	}

	var devices []models.Device

	for rows.Next() {
		var device models.Device
		err = rows.StructScan(&device)
		if err != nil {
			return nil, err
		}

		err = p.FindSensors(&device)
		if err != nil {
			return nil, err
		}

		devices = append(devices, device)

	}

	return devices, nil
}

func (p *PostgreDB) FindSensors(device *models.Device) error {
	sensorsRows, err := p.DB.Queryx("SELECT * FROM sensors WHERE deviceid = " + fmt.Sprint(device.Id))
	if err != nil {
		return err
	}

	var sensors []*models.Sensors

	for sensorsRows.Next() {
		var sensor models.Sensors
		err := sensorsRows.StructScan(&sensor)
		if err != nil {
			return err
		}

		sensors = append(sensors, &sensor)
	}

	device.Sensors = sensors

	return nil
}

func (p *PostgreDB) GetDevice(id string) (models.Device, error) {
	query := "SELECT id, name, description FROM device WHERE id=" + id
	row := p.DB.QueryRow(query)

	device := models.Device{}
	err := row.Scan(&device.Id, &device.Name, &device.Description)
	if err != nil {
		return device, err
	}

	err = p.FindSensors(&device)
	if err != nil {
		return device, err
	}

	return device, nil
}

func (p *PostgreDB) CreateDevice(vars map[string]string) error {
	name := vars["name"]
	description := vars["description"]
	_, err := p.DB.Exec("INSERT INTO device (name, description) VALUES ($1, $2)", name, description)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgreDB) UpdateDevice(vars map[string]string) error {
	name := vars["name"]
	description := vars["description"]
	id := vars["id"]
	_, err := p.DB.Exec("UPDATE device SET name=$1, description=$2 WHERE id=$3", name, description, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgreDB) DeleteDevice(id string) error {
	_, err := p.DB.Exec("DELETE FROM device WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgreDB) GetAllSensors() ([]models.Sensors, error) {
	query := "SELECT * FROM sensors"
	rows, err := p.DB.Queryx(query)
	if err != nil {
		return nil, err
	}

	var sensors []models.Sensors

	for rows.Next() {
		var sensor models.Sensors
		err = rows.StructScan(&sensor)
		if err != nil {
			return nil, err
		}

		sensors = append(sensors, sensor)
	}

	return sensors, nil
}

func (p *PostgreDB) GetSensor(id string) (models.Sensors, error) {
	query := "SELECT id, deviceid, name, description, unit FROM sensors WHERE id=" + id
	row := p.DB.QueryRow(query)

	sensor := models.Sensors{}
	err := row.Scan(&sensor.Id, &sensor.Device_id, &sensor.Name, &sensor.Description, &sensor.Unit)
	if err != nil {
		return sensor, err
	}

	return sensor, nil
}

func (p *PostgreDB) CreateSensor(vars map[string]string) error {
	deviceid := vars["deviceid"]
	name := vars["name"]
	description := vars["description"]
	unit := vars["unit"]
	_, err := p.DB.Exec("INSERT INTO sensors (deviceid, name, description, unit) VALUES ($1, $2, $3, $4)",
		deviceid, name, description, unit)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgreDB) UpdateSensor(vars map[string]string) error {
	deviceid := vars["deviceid"]
	name := vars["name"]
	description := vars["description"]
	unit := vars["unit"]
	id := vars["id"]

	_, err := p.DB.Exec("UPDATE sensors SET deviceid=$1, name=$2, description=$3, unit=$4 WHERE id=$5",
		deviceid, name, description, unit, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgreDB) DeleteSensor(id string) error {
	_, err := p.DB.Exec("DELETE FROM sensors WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
