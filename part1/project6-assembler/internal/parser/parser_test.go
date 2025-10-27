package parser_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestParser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parser Suite")
}

var _ = Describe("Parser", func() {
	Describe("ParseLine", func() {
		Context("when parsing valid instructions", func() {
			It("should parse A-instructions", func() {
				line, err := ParseLine("@123")
				Expect(err).NotTo(HaveOccurred())
				Expect(line.Type()).To(Equal(A))
				Expect(line.Symbol()).To(Equal("123"))
			})

			It("should parse C-instruction with dest", func() {
				line, err := ParseLine("D=M-1")
				Expect(err).NotTo(HaveOccurred())
				Expect(line.Type()).To(Equal(C))
				Expect(line.Symbol()).To(Equal(""))
				Expect(line.Comp()).To(Equal("M-1"))
				Expect(line.Dest()).To(Equal("D"))
				Expect(line.Jump()).To(Equal(""))
			})

			It("should parse C-instruction with jump", func() {
				line, err := ParseLine("D;JLT")
				Expect(err).NotTo(HaveOccurred())
				Expect(line.Type()).To(Equal(C))
				Expect(line.Symbol()).To(Equal(""))
				Expect(line.Comp()).To(Equal("D"))
				Expect(line.Dest()).To(Equal(""))
				Expect(line.Jump()).To(Equal("JLT"))
			})

			It("should parse labels", func() {
				line, err := ParseLine("(LABEL)")
				Expect(err).NotTo(HaveOccurred())
				Expect(line.Type()).To(Equal(L))
				Expect(line.Symbol()).To(Equal("LABEL"))
			})

			It("should return error if line contains only comment", func() {
				line, err := ParseLine("// hello darkness my old friend")
				Expect(err).To(Equal(ErrLineNotInstruction))
				Expect(line).To(BeNil())
			})

			It("should return error if line is empty", func() {
				line, err := ParseLine("")
				Expect(err).To(Equal(ErrLineNotInstruction))
				Expect(line).To(BeNil())
			})

			It("should return error if line contains no symbols", func() {
				line, err := ParseLine("     ")
				Expect(err).To(Equal(ErrLineNotInstruction))
				Expect(line).To(BeNil())
			})
		})
	})
})
