package sensors_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSystemSensors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SystemSensors Suite")
}
