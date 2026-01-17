package parser_test

import (
	"testing"
	"vm/parser"

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

		It("should parse fibonacci program and count instruction types", func() {
			p := parser.NewParser("testdata/fibonacci.vm")

			counts := make(map[parser.CommandType]int)
			for cmd, err := range p.Commands {
				Expect(err).NotTo(HaveOccurred())
				counts[cmd.Type]++
			}

			Expect(counts[parser.Function]).To(Equal(1))
			Expect(counts[parser.Push]).To(Equal(7))
			Expect(counts[parser.Arithmetic]).To(Equal(4))
			Expect(counts[parser.IfGoto]).To(Equal(1))
			Expect(counts[parser.Goto]).To(Equal(1))
			Expect(counts[parser.Label]).To(Equal(2))
			Expect(counts[parser.Return]).To(Equal(2))
			Expect(counts[parser.Call]).To(Equal(2))

			total := 0
			for _, count := range counts {
				total += count
			}
			Expect(total).To(Equal(20))
		})
	})
})
