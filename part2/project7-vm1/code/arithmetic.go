package code

import "strings"

var Mappings = map[string]func() string{
	"add": Add,
	"sub": Sub,
	"neg": Neg,
	"eq":  Eq,
	"gt":  Gt,
	"lt":  Lt,
	"and": And,
	"or":  Or,
	"not": Not,
}

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

func Eq() string {
	sb := &strings.Builder{}
	sb.WriteString("// eq\n\n")
	popBinary(sb)

	sb.WriteString("D=D&M\n")
	sb.WriteString("D=!D\n")
	sb.WriteString("D=D&M\n")
	sb.WriteString("@EQ\n")
	sb.WriteString("D;JEQ\n")
	sb.WriteString("@NOT_EQ\n")
	sb.WriteString("D;JNE\n")
	sb.WriteString("(EQ)\n")
	sb.WriteString("D=1\n")
	sb.WriteString("(NOT_EQ)\n")
	sb.WriteString("D=-1\n")

	writeBackBinary(sb)
	sb.WriteString("\n")

	return sb.String()
}

func Gt() string {
	sb := &strings.Builder{}
	sb.WriteString("// gt\n\n")
	popBinary(sb)

	sb.WriteString("D=M-D\n")
	sb.WriteString("@GT\n")
	sb.WriteString("D;JGT\n")
	sb.WriteString("@NOT_GT\n")
	sb.WriteString("D;JMP\n")
	sb.WriteString("(GT)\n")
	sb.WriteString("D=1\n")
	sb.WriteString("D=-1\n")

	writeBackBinary(sb)
	sb.WriteString("\n")

	return sb.String()
}

func Lt() string {
	sb := &strings.Builder{}
	sb.WriteString("// lt\n\n")
	popBinary(sb)

	sb.WriteString("D=M-D\n")
	sb.WriteString("@LT\n")
	sb.WriteString("D;JLT\n")
	sb.WriteString("@NOT_LT\n")
	sb.WriteString("D;JMP\n")
	sb.WriteString("(LT)\n")
	sb.WriteString("D=1\n")
	sb.WriteString("(NOT_LT)\n")
	sb.WriteString("D=-1\n")

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
	sb.WriteString("M=D\n")
	sb.WriteString("D=A\n")
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
