package parser_test

import (
	"project7-vm1/parser"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestParser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parser Suite")
}

var _ = Describe("Parser", func() {
	Describe("Commands", func() {
		It("should read program from testdata directory", func() {
			p := parser.NewParser("testdata/simple.vm")

			var commands []*parser.Command
			for cmd, err := range p.Commands {
				Expect(err).NotTo(HaveOccurred())
				commands = append(commands, cmd)
			}

			Expect(commands).To(HaveLen(3))
		})
	})
})
