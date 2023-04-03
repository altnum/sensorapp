package executiontimer

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

var _ = Describe("Testing the timer struct responsible for the execution", func() {
	timer := Timer{}
	Context("SetTotalDuration()", func() {
		It("given error when getting info from flag, then return error", func() {
			getFlagValueString = func(s string) (string, error) {
				return "test", errors.New("err")
			}

			cmd := cobra.Command{}
			err := timer.SetTotalDuration(&cmd)

			Expect(err).ToNot(BeNil())
		})
		It("given error when parsing, then return error", func() {
			getFlagValueString = func(s string) (string, error) {
				return "0", nil
			}

			cmd := cobra.Command{}
			err := timer.SetTotalDuration(&cmd)

			Expect(err).ToNot(BeNil())
		})
	})

	Context("SetDeltaDuration()", func() {
		It("given error when getting info from flag, then return error", func() {
			getFlagValueString = func(s string) (string, error) {
				return "test", errors.New("err")
			}
			cmd := cobra.Command{}
			err := timer.SetDeltaDuration(&cmd)

			Expect(err).ToNot(BeNil())
		})
		It("given error when parsing, then return error", func() {
			getFlagValueString = func(s string) (string, error) {
				return "test", nil
			}

			cmd := cobra.Command{}
			err := timer.SetDeltaDuration(&cmd)

			Expect(err).ToNot(BeNil())
		})
		It("given error when delta_duration input is bigger than the input of total_duration,"+
			"then return error", func() {
			getFlagValueString = func(s string) (string, error) {
				if s == "delta_duration" {
					return "3", nil
				}
				if s == "total_duraiton" {
					return "2", nil
				}

				return "", nil
			}

			cmd := cobra.Command{}
			err := timer.SetDeltaDuration(&cmd)

			Expect(err).ToNot(BeNil())
		})
	})

	Context("SetUpTimer", func() {
		It("given errors when parsing, then return error", func() {
			parseInt = func(s string, base, bitSize int) (i int64, err error) {
				return -1, errors.New("err")
			}

			cmd := cobra.Command{}
			err := timer.SetUpTimer(&cmd)

			Expect(err).ToNot(BeNil())
		})
	})

	Context("Getters", func() {
		It("when called GetTotalDuration then return total_duration", func() {
			getFlagValueString = func(s string) (string, error) {
				return "1", nil
			}

			tD := timer.GetTotalDuration()
			Expect(tD).To(BeNumerically("==", 1))
		})
		It("when called GetDeltaDuration then return delta_duration", func() {
			getFlagValueString = func(s string) (string, error) {
				return "0", nil
			}

			tD := timer.GetDeltaDuration()
			Expect(tD).To(BeNumerically("==", 0))
		})
	})
})
