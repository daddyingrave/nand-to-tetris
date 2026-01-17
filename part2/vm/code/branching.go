package code

import (
	"fmt"
	"strings"
	"vm/internal/utils"
)

func Label(fileName string, functionName string, label string) string {
	sb := &strings.Builder{}

	utils.WriteSBf(sb, "// label: %s, func: %s, file: %s\n", label, functionName, fileName)
	utils.WriteSBf(sb, "(%s)", labelName(fileName, functionName, label))
	utils.WriteSBf(sb, "")

	return sb.String()
}

func GoTo(fileName string, functionName string, label string) string {
	sb := &strings.Builder{}

	utils.WriteSBf(sb, "// goto: %s, func: %s, file: %s\n", label, functionName, fileName)
	utils.WriteSBf(sb, "  @%s", labelName(fileName, functionName, label))
	utils.WriteSBf(sb, "  0;JMP")
	utils.WriteSBf(sb, "")

	return sb.String()
}

func IfGoTo(fileName string, functionName string, label string) string {
	sb := &strings.Builder{}

	utils.WriteSBf(sb, "// if-goto: %s, func: %s, file: %s\n", label, functionName, fileName)
	utils.WriteSBf(sb, "  @SP")
	utils.WriteSBf(sb, "  AM=M-1")
	utils.WriteSBf(sb, "  D=M")
	utils.WriteSBf(sb, "  @%s", labelName(fileName, functionName, label))
	utils.WriteSBf(sb, "  D;JNE")
	utils.WriteSBf(sb, "")

	return sb.String()
}

func labelName(fileName string, functionName string, label string) string {
	return fmt.Sprintf("%s.%s$%s", fileName, functionName, label)
}
