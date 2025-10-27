package code_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMnemonics(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mnemonics Suite")
}

var _ = Describe("Mnemonics", func() {
	Describe("CompToBinary", func() {
		Context("when a=0 (using A register)", func() {
			It("should translate '0' to binary", func() {
				result := CompToBinary("0")
				Expect(result).To(Equal("0101010"))
			})

			It("should translate '1' to binary", func() {
				result := CompToBinary("1")
				Expect(result).To(Equal("0111111"))
			})

			It("should translate '-1' to binary", func() {
				result := CompToBinary("-1")
				Expect(result).To(Equal("0111010"))
			})

			It("should translate 'D' to binary", func() {
				result := CompToBinary("D")
				Expect(result).To(Equal("0001100"))
			})

			It("should translate 'A' to binary", func() {
				result := CompToBinary("A")
				Expect(result).To(Equal("0110000"))
			})

			It("should translate 'D+1' to binary", func() {
				result := CompToBinary("D+1")
				Expect(result).To(Equal("0011111"))
			})

			It("should translate 'A-1' to binary", func() {
				result := CompToBinary("A-1")
				Expect(result).To(Equal("0110010"))
			})

			It("should translate 'D-A' to binary", func() {
				result := CompToBinary("D-A")
				Expect(result).To(Equal("0010011"))
			})
		})

		Context("when a=1 (using M register)", func() {
			It("should translate 'M' to binary", func() {
				result := CompToBinary("M")
				Expect(result).To(Equal("1110000"))
			})

			It("should translate '!M' to binary", func() {
				result := CompToBinary("!M")
				Expect(result).To(Equal("1110001"))
			})

			It("should translate '-M' to binary", func() {
				result := CompToBinary("-M")
				Expect(result).To(Equal("1110011"))
			})

			It("should translate 'M+1' to binary", func() {
				result := CompToBinary("M+1")
				Expect(result).To(Equal("1110111"))
			})

			It("should translate 'D-M' to binary", func() {
				result := CompToBinary("D-M")
				Expect(result).To(Equal("1010011"))
			})
		})
	})

	Describe("DestToBinary", func() {
		It("should translate 'null' (no destination) to binary", func() {
			result := DestToBinary("")
			Expect(result).To(Equal("000"))
		})

		It("should translate 'M' to binary", func() {
			result := DestToBinary("M")
			Expect(result).To(Equal("001"))
		})

		It("should translate 'D' to binary", func() {
			result := DestToBinary("D")
			Expect(result).To(Equal("010"))
		})

		It("should translate 'A' to binary", func() {
			result := DestToBinary("A")
			Expect(result).To(Equal("100"))
		})

		It("should translate 'AM' to binary", func() {
			result := DestToBinary("AM")
			Expect(result).To(Equal("101"))
		})

		It("should translate 'ADM' to binary", func() {
			result := DestToBinary("ADM")
			Expect(result).To(Equal("111"))
		})
	})

	Describe("JumpToBinary", func() {
		It("should translate 'null' (no jump) to binary", func() {
			result := JumpToBinary("")
			Expect(result).To(Equal("000"))
		})

		It("should translate 'JGT' (if comp > 0 jump) to binary", func() {
			result := JumpToBinary("JGT")
			Expect(result).To(Equal("001"))
		})

		It("should translate 'JEQ' (if comp = 0 jump) to binary", func() {
			result := JumpToBinary("JEQ")
			Expect(result).To(Equal("010"))
		})

		It("should translate 'JGE' (if comp >= 0 jump) to binary", func() {
			result := JumpToBinary("JGE")
			Expect(result).To(Equal("011"))
		})

		It("should translate 'JLT' (if comp < 0 jump) to binary", func() {
			result := JumpToBinary("JLT")
			Expect(result).To(Equal("100"))
		})

		It("should translate 'JNE' (if comp != 0 jump) to binary", func() {
			result := JumpToBinary("JNE")
			Expect(result).To(Equal("101"))
		})

		It("should translate 'JMP' (unconditional jump) to binary", func() {
			result := JumpToBinary("JMP")
			Expect(result).To(Equal("111"))
		})
	})
})
