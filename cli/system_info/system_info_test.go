package systeminfo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Measurements", func() {
	Context("SetTimeStamp()", func() {
		It("given the local time, then set it to the measurement istance", func() {
			measurement := Measurement{}

			measurement.SetTimeStamp()

			Expect(measurement.MeasuredAt).ToNot(BeIdenticalTo(""))
		})
	})
})
