package executiontimer_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestExecutionTimer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ExecutionTimer Suite")
}
