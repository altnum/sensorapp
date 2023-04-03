package databases

import (
	"context"
	"errors"

	. "github.com/altnum/sensorapp/databases/database_instances"
	mocksDatabase "github.com/altnum/sensorapp/databases/database_instances/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Databases functionality", func() {
	dbFactory := &DatabaseFactory{}
	Context("RetrieveDatabases", func() {
		It("given correct input when creating databases, then return slice of IDB", func() {
			dbs := dbFactory.RetrieveDatabases()

			Expect(dbs).ToNot(BeEmpty())
		})
	})

	Context("ConnectDatabases", func() {
		context := context.Background()
		It("given a database, when trying to connect to it, then return no error", func() {
			var dbs []IDB
			dbs = append(dbs, &PostgreDB{}, &InfluxDB{})
			err := dbFactory.ConnectDatabases(context, dbs)

			Expect(err).To(BeNil())
		})
		It("given error when opening connection to IDB, when trying to connect, then return error", func() {
			pgdb := &mocksDatabase.IDB{}
			pgdb.On("Open", context).Return(errors.New(""))
			dbs := []IDB{pgdb}

			err := dbFactory.ConnectDatabases(context, dbs)

			Expect(err).ToNot(BeNil())
		})
	})
})
