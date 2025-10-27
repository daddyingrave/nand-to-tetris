package assembler_test

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAssembler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Assembler Suite")
}

var _ = Describe("Assembler", func() {
	Describe("Translate", func() {
		DescribeTable("should translate .asm files to .hack files correctly",
			func(asmFile string, expectedHackFile string) {
				tmpDir := GinkgoT().TempDir()

				asmContent, err := os.ReadFile(asmFile)
				Expect(err).NotTo(HaveOccurred())

				tmpAsmPath := filepath.Join(tmpDir, filepath.Base(asmFile))
				err = os.WriteFile(tmpAsmPath, asmContent, 0600)
				Expect(err).NotTo(HaveOccurred())

				Translate(tmpAsmPath)

				tmpHackPath := filepath.Join(tmpDir, filepath.Base(asmFile[:len(asmFile)-4]+".hack"))
				generatedContent, err := os.ReadFile(tmpHackPath)
				Expect(err).NotTo(HaveOccurred())

				expectedContent, err := os.ReadFile(expectedHackFile)
				Expect(err).NotTo(HaveOccurred())

				Expect(string(generatedContent)).To(Equal(string(expectedContent)))
			},
			Entry("Add.asm", "testdata/Add.asm", "testdata/Add-expected.hack"),
			Entry("Max.asm", "testdata/Max.asm", "testdata/Max-expected.hack"),
			Entry("Rect.asm", "testdata/Rect.asm", "testdata/Rect-expected.hack"),
			Entry("Pong.asm", "testdata/Pong.asm", "testdata/Pong-expected.hack"),
		)
	})
})
