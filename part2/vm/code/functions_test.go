package code_test

import (
	"os"
	"vm/code"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Functions", func() {
	Describe("Call", func() {
		It("should generate correct assembly for call", func() {
			result, err := code.Call("file", "caller", "callee", 2, 1)
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/functions/call.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})

	Describe("Function", func() {
		It("should generate correct assembly for function", func() {
			result, err := code.Function("file", "fn", 2)
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/functions/function.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})

	Describe("Return", func() {
		It("should generate correct assembly for return", func() {
			result, err := code.Return("file", "fn")
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/functions/return.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})
})
