package code_test

import (
	"os"
	"project7-vm1/code"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PushPop", func() {
	Describe("StackPush", func() {
		It("should generate correct assembly for push constant", func() {
			result, err := code.StackPush("constant", 10, "test.asm")
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/push_pop/push_constant_10.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})

		It("should generate correct assembly for push local", func() {
			result, err := code.StackPush("local", 5, "test.asm")
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/push_pop/push_local_5.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})

		It("should generate correct assembly for push static", func() {
			result, err := code.StackPush("static", 3, "test.asm")
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/push_pop/push_static_3.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})

	Describe("StackPop", func() {
		It("should generate correct assembly for pop local", func() {
			result, err := code.StackPop("local", 2, "test.asm")
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/push_pop/pop_local_2.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})

		It("should generate correct assembly for pop argument", func() {
			result, err := code.StackPop("argument", 1, "test.asm")
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/push_pop/pop_argument_1.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})

		It("should generate correct assembly for pop static", func() {
			result, err := code.StackPop("static", 7, "test.asm")
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/push_pop/pop_static_7.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})
})
