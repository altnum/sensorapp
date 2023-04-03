package writer_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestParallelWriter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ParallelWriter Suite")
}
