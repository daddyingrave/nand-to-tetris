package code_test

import (
	"os"
	"vm/code"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PushPop", func() {
	Describe("StackPush", func() {
		It("should generate correct assembly for push constant", func() {
			result, err := code.StackPush("constant", 10, "test")
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/push_pop/push_constant_10.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})

		It("should generate correct assembly for push local", func() {
			result, err := code.StackPush("local", 5, "test")
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/push_pop/push_local_5.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})

		It("should generate correct assembly for push static", func() {
			result, err := code.StackPush("static", 3, "test")
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/push_pop/push_static_3.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})

		It("should generate correct assembly for push pointer 0 (THIS)", func() {
			result, err := code.StackPush("pointer", 0, "test")
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/push_pop/push_pointer_0.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})

		It("should generate correct assembly for push pointer 1 (THAT)", func() {
			result, err := code.StackPush("pointer", 1, "test")
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/push_pop/push_pointer_1.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})

	Describe("StackPop", func() {
		It("should generate correct assembly for pop local", func() {
			result, err := code.StackPop("local", 2, "test")
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/push_pop/pop_local_2.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})

		It("should generate correct assembly for pop argument", func() {
			result, err := code.StackPop("argument", 1, "test")
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/push_pop/pop_argument_1.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})

		It("should generate correct assembly for pop static", func() {
			result, err := code.StackPop("static", 7, "test")
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/push_pop/pop_static_7.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})

		It("should generate correct assembly for pop pointer 0 (THIS)", func() {
			result, err := code.StackPop("pointer", 0, "test")
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/push_pop/pop_pointer_0.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})

		It("should generate correct assembly for pop pointer 1 (THAT)", func() {
			result, err := code.StackPop("pointer", 1, "test")
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/push_pop/pop_pointer_1.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})
})
