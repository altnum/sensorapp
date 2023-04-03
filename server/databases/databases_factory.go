package databases

import (
	"context"

	. "github.com/altnum/sensorapp/databases/database_instances"
)

var dbs []IDB

type IDatabaseFactory interface {
	RetrieveDatabases() []IDB
	ConnectDatabases(context.Context, []IDB) error
}

type DatabaseFactory struct{}

func init() {
	dbs = append(dbs, &PostgreDB{}, &InfluxDB{})
}

func (d *DatabaseFactory) RetrieveDatabases() []IDB {
	return dbs
}

func (d *DatabaseFactory) ConnectDatabases(context context.Context, dbs []IDB) error {
	for _, db := range dbs {
		err := db.Open(context)
		if err != nil {
			return err
		}
	}

	return nil
}
