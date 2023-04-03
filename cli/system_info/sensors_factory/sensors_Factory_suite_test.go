package sensorsfactory_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSensorsFactory(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SensorsFactory Suite")
}
