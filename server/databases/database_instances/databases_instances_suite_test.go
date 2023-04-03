package database_instances_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDatabasesInstances(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DatabasesInstances Suite")
}
