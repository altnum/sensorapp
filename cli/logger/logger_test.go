package logger

import (
	"errors"
	"io/fs"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testing formatter package", func() {
	Context("GetLogger()", func() {
		It("get the instances of the logger", func() {
			l := GetLogger()

			Expect(l.Info).ToNot(BeNil())
			Expect(l.Warning).ToNot(BeNil())
			Expect(l.Error).ToNot(BeNil())
		})
		It("", func() {
			openFile = func(s string, i int, fm fs.FileMode) (*os.File, error) {
				return nil, errors.New("")
			}
			cwd, _ := os.Getwd()
			file, err := findLogsFile(cwd)

			Expect(err).ToNot(BeNil())
			Expect(file).To(BeNil())
		})
	})
})
