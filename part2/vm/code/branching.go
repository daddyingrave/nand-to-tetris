package code

import (
	"fmt"
	"strings"
)

func Label(fileName string, functionName string, label string) string {
	sb := &strings.Builder{}

	sb.WriteString(fmt.Sprintf("// label: %s, func: %s, file: %s\n\n", label, functionName, fileName))
	sb.WriteString(fmt.Sprintf("(%s)\n", labelName(fileName, functionName, label)))
	sb.WriteString("\n")

	return sb.String()
}

func GoTo(fileName string, functionName string, label string) string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("// goto: %s, func: %s, file: %s\n\n", label, functionName, fileName))
	sb.WriteString(fmt.Sprintf("  @%s\n", labelName(fileName, functionName, label)))
	sb.WriteString("  0;JMP\n")
	sb.WriteString("\n")

	return sb.String()
}

func IfGoTo(fileName string, functionName string, label string) string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("// if-goto: %s, func: %s, file: %s\n\n", label, functionName, fileName))
	sb.WriteString("  @SP\n")
	sb.WriteString("  AM=M-1\n")
	sb.WriteString("  D=M\n")
	sb.WriteString(fmt.Sprintf("  @%s\n", labelName(fileName, functionName, label)))
	sb.WriteString("  D;JNE\n")
	sb.WriteString("\n")

	return sb.String()
}

func labelName(fileName string, functionName string, label string) string {
	return fmt.Sprintf("%s.%s$%s", fileName, functionName, label)
}
