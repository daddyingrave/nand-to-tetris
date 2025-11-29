package code_test

import (
	"os"
	"vm/code"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Arithmetic", func() {
	Describe("Add", func() {
		It("should generate correct assembly for add", func() {
			result := code.Add()

			expected, err := os.ReadFile("testdata/arithmetic/add.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})

	Describe("Sub", func() {
		It("should generate correct assembly for sub", func() {
			result := code.Sub()

			expected, err := os.ReadFile("testdata/arithmetic/sub.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})

	Describe("Neg", func() {
		It("should generate correct assembly for neg", func() {
			result := code.Neg()

			expected, err := os.ReadFile("testdata/arithmetic/neg.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})

	Describe("Eq", func() {
		It("should generate correct assembly for eq", func() {
			result := code.Eq(666)

			expected, err := os.ReadFile("testdata/arithmetic/eq.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})

	Describe("Gt", func() {
		It("should generate correct assembly for gt", func() {
			result := code.Gt(666)

			expected, err := os.ReadFile("testdata/arithmetic/gt.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})

	Describe("Lt", func() {
		It("should generate correct assembly for lt", func() {
			result := code.Lt(666)

			expected, err := os.ReadFile("testdata/arithmetic/lt.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})

	Describe("And", func() {
		It("should generate correct assembly for and", func() {
			result := code.And()

			expected, err := os.ReadFile("testdata/arithmetic/and.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})

	Describe("Or", func() {
		It("should generate correct assembly for or", func() {
			result := code.Or()

			expected, err := os.ReadFile("testdata/arithmetic/or.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})

	Describe("Not", func() {
		It("should generate correct assembly for not", func() {
			result := code.Not()

			expected, err := os.ReadFile("testdata/arithmetic/not.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})
})
