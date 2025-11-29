package code

import (
	"fmt"
	"strings"
)

func Add() string {
	sb := &strings.Builder{}
	sb.WriteString("// add\n\n")
	popBinary(sb)

	sb.WriteString("D=D+M\n")

	writeBackBinary(sb)
	sb.WriteString("\n")

	return sb.String()
}

func Sub() string {
	sb := &strings.Builder{}
	sb.WriteString("// sub\n\n")
	popBinary(sb)

	sb.WriteString("D=M-D\n")

	writeBackBinary(sb)
	sb.WriteString("\n")

	return sb.String()
}

func Neg() string {
	sb := &strings.Builder{}
	sb.WriteString("// neg\n\n")
	popUnary(sb)

	sb.WriteString("D=-D\n")

	writeBackUnary(sb)
	sb.WriteString("\n")

	return sb.String()
}

func Eq(counter int) string {
	sb := &strings.Builder{}
	sb.WriteString("// eq\n\n")
	popBinary(sb)

	sb.WriteString("D=D-M\n")
	sb.WriteString(fmt.Sprintf("@EQ_%d\n", counter))
	sb.WriteString("D;JEQ\n")
	sb.WriteString(fmt.Sprintf("@NOT_EQ_%d\n", counter))
	sb.WriteString("D;JNE\n")
	sb.WriteString(fmt.Sprintf("(EQ_%d)\n", counter))
	sb.WriteString("D=-1\n")
	sb.WriteString(fmt.Sprintf("@EQ_END_%d\n", counter))
	sb.WriteString("D;JMP\n")
	sb.WriteString(fmt.Sprintf("(NOT_EQ_%d)\n", counter))
	sb.WriteString("D=0\n")
	sb.WriteString(fmt.Sprintf("(EQ_END_%d)\n", counter))

	writeBackBinary(sb)
	sb.WriteString("\n")

	return sb.String()
}

func Gt(counter int) string {
	sb := &strings.Builder{}
	sb.WriteString("// gt\n\n")
	popBinary(sb)

	sb.WriteString("D=M-D\n")
	sb.WriteString(fmt.Sprintf("@GT_%d\n", counter))
	sb.WriteString("D;JGT\n")
	sb.WriteString(fmt.Sprintf("@NOT_GT_%d\n", counter))
	sb.WriteString("D;JMP\n")
	sb.WriteString(fmt.Sprintf("(GT_%d)\n", counter))
	sb.WriteString("D=-1\n")
	sb.WriteString(fmt.Sprintf("@GT_END_%d\n", counter))
	sb.WriteString("D;JMP\n")
	sb.WriteString(fmt.Sprintf("(NOT_GT_%d)\n", counter))
	sb.WriteString("D=0\n")
	sb.WriteString(fmt.Sprintf("(GT_END_%d)\n", counter))

	writeBackBinary(sb)
	sb.WriteString("\n")

	return sb.String()
}

func Lt(counter int) string {
	sb := &strings.Builder{}
	sb.WriteString("// lt\n\n")
	popBinary(sb)

	sb.WriteString("D=M-D\n")
	sb.WriteString(fmt.Sprintf("@LT_%d\n", counter))
	sb.WriteString("D;JLT\n")
	sb.WriteString(fmt.Sprintf("@NOT_LT_%d\n", counter))
	sb.WriteString("D;JMP\n")
	sb.WriteString(fmt.Sprintf("(LT_%d)\n", counter))
	sb.WriteString("D=-1\n")
	sb.WriteString(fmt.Sprintf("@LT_END_%d\n", counter))
	sb.WriteString("D;JMP\n")
	sb.WriteString(fmt.Sprintf("(NOT_LT_%d)\n", counter))
	sb.WriteString("D=0\n")
	sb.WriteString(fmt.Sprintf("(LT_END_%d)\n", counter))

	writeBackBinary(sb)
	sb.WriteString("\n")

	return sb.String()
}

func And() string {
	sb := &strings.Builder{}
	sb.WriteString("// and\n\n")
	popBinary(sb)

	sb.WriteString("D=D&M\n")

	writeBackBinary(sb)
	sb.WriteString("\n")

	return sb.String()
}

func Or() string {
	sb := &strings.Builder{}
	sb.WriteString("// or\n\n")
	popBinary(sb)

	sb.WriteString("D=D|M\n")

	writeBackBinary(sb)
	sb.WriteString("\n")

	return sb.String()
}

func Not() string {
	sb := &strings.Builder{}
	sb.WriteString("// not\n\n")
	popUnary(sb)

	sb.WriteString("D=-D\n")
	sb.WriteString("D=D-1\n")

	writeBackUnary(sb)
	sb.WriteString("\n")

	return sb.String()
}

func popBinary(sb *strings.Builder) *strings.Builder {
	sb.WriteString("@SP\n")
	sb.WriteString("A=M\n")
	sb.WriteString("A=A-1\n")
	sb.WriteString("D=M\n")
	sb.WriteString("A=A-1\n")

	return sb
}

func popUnary(sb *strings.Builder) *strings.Builder {
	sb.WriteString("@SP\n")
	sb.WriteString("A=M\n")
	sb.WriteString("A=A-1\n")
	sb.WriteString("D=M\n")

	return sb
}

func writeBackBinary(sb *strings.Builder) *strings.Builder {
	sb.WriteString("@SP\n")
	sb.WriteString("A=M\n")
	sb.WriteString("A=A-1\n")
	sb.WriteString("A=A-1\n")
	sb.WriteString("M=D\n")
	sb.WriteString("D=A+1\n")
	sb.WriteString("@SP\n")
	sb.WriteString("M=D\n")

	return sb
}

func writeBackUnary(sb *strings.Builder) *strings.Builder {
	sb.WriteString("@SP\n")
	sb.WriteString("A=M\n")
	sb.WriteString("A=A-1\n")
	sb.WriteString("M=D\n")

	return sb
}
