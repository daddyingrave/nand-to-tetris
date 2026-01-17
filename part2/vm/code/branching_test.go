package code_test

import (
	"os"
	"vm/code"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Branching", func() {
	Describe("Label", func() {
		It("should generate correct assembly for label", func() {
			result := code.Label("file", "func", "lbl")

			expected, err := os.ReadFile("testdata/branching/label.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})

	Describe("GoTo", func() {
		It("should generate correct assembly for goto", func() {
			result := code.GoTo("file", "func", "lbl")

			expected, err := os.ReadFile("testdata/branching/goto.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})

	Describe("IfGoTo", func() {
		It("should generate correct assembly for if-goto", func() {
			result := code.IfGoTo("file", "func", "lbl")

			expected, err := os.ReadFile("testdata/branching/if-goto.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})
})
