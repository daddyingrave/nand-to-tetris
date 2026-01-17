package code_test

import (
	"os"
	"vm/code"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bootstrap", func() {
	Describe("BootstrapCode", func() {
		It("should generate correct assembly for bootstrap", func() {
			result, err := code.BootstrapCode("file")
			Expect(err).NotTo(HaveOccurred())

			expected, err := os.ReadFile("testdata/bootstrap/bootstrap.asm")
			Expect(err).NotTo(HaveOccurred())

			Expect(result).To(Equal(string(expected)))
		})
	})
})
