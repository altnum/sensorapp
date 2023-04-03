package systeminfo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAppMeasurements(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SystemInfo Suite")
}
