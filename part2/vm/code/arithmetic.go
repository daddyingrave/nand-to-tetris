package code

import (
	"strings"
	"vm/internal/utils"
)

func Add() string {
	sb := &strings.Builder{}
	utils.WriteSBf(sb, "// add\n")
	popBinary(sb)

	utils.WriteSBf(sb, "  D=D+M")

	writeBackBinary(sb)
	utils.WriteSBf(sb, "")

	return sb.String()
}

func Sub() string {
	sb := &strings.Builder{}
	utils.WriteSBf(sb, "// sub\n")
	popBinary(sb)

	utils.WriteSBf(sb, "  D=M-D")

	writeBackBinary(sb)
	utils.WriteSBf(sb, "")

	return sb.String()
}

func Neg() string {
	sb := &strings.Builder{}
	utils.WriteSBf(sb, "// neg\n")
	popUnary(sb)

	utils.WriteSBf(sb, "  D=-D")

	writeBackUnary(sb)
	utils.WriteSBf(sb, "")

	return sb.String()
}

func Eq(counter int) string {
	sb := &strings.Builder{}
	utils.WriteSBf(sb, "// eq\n")
	popBinary(sb)

	utils.WriteSBf(sb, "  D=D-M")
	utils.WriteSBf(sb, "  @EQ_%d", counter)
	utils.WriteSBf(sb, "  D;JEQ")
	utils.WriteSBf(sb, "  @NOT_EQ_%d", counter)
	utils.WriteSBf(sb, "  D;JNE")
	utils.WriteSBf(sb, "(EQ_%d)", counter)
	utils.WriteSBf(sb, "  D=-1")
	utils.WriteSBf(sb, "  @EQ_END_%d", counter)
	utils.WriteSBf(sb, "  D;JMP")
	utils.WriteSBf(sb, "(NOT_EQ_%d)", counter)
	utils.WriteSBf(sb, "  D=0")
	utils.WriteSBf(sb, "(EQ_END_%d)", counter)

	writeBackBinary(sb)
	utils.WriteSBf(sb, "")

	return sb.String()
}

func Gt(counter int) string {
	sb := &strings.Builder{}
	utils.WriteSBf(sb, "// gt\n")
	popBinary(sb)

	utils.WriteSBf(sb, "  D=M-D")
	utils.WriteSBf(sb, "  @GT_%d", counter)
	utils.WriteSBf(sb, "  D;JGT")
	utils.WriteSBf(sb, "  @NOT_GT_%d", counter)
	utils.WriteSBf(sb, "  D;JMP")
	utils.WriteSBf(sb, "(GT_%d)", counter)
	utils.WriteSBf(sb, "  D=-1")
	utils.WriteSBf(sb, "  @GT_END_%d", counter)
	utils.WriteSBf(sb, "  D;JMP")
	utils.WriteSBf(sb, "(NOT_GT_%d)", counter)
	utils.WriteSBf(sb, "  D=0")
	utils.WriteSBf(sb, "(GT_END_%d)", counter)

	writeBackBinary(sb)
	utils.WriteSBf(sb, "")

	return sb.String()
}

func Lt(counter int) string {
	sb := &strings.Builder{}
	utils.WriteSBf(sb, "// lt\n")
	popBinary(sb)

	utils.WriteSBf(sb, "  D=M-D")
	utils.WriteSBf(sb, "  @LT_%d", counter)
	utils.WriteSBf(sb, "  D;JLT")
	utils.WriteSBf(sb, "  @NOT_LT_%d", counter)
	utils.WriteSBf(sb, "  D;JMP")
	utils.WriteSBf(sb, "(LT_%d)", counter)
	utils.WriteSBf(sb, "  D=-1")
	utils.WriteSBf(sb, "  @LT_END_%d", counter)
	utils.WriteSBf(sb, "  D;JMP")
	utils.WriteSBf(sb, "(NOT_LT_%d)", counter)
	utils.WriteSBf(sb, "  D=0")
	utils.WriteSBf(sb, "(LT_END_%d)", counter)

	writeBackBinary(sb)
	utils.WriteSBf(sb, "")

	return sb.String()
}

func And() string {
	sb := &strings.Builder{}
	utils.WriteSBf(sb, "// and\n")
	popBinary(sb)

	utils.WriteSBf(sb, "  D=D&M")

	writeBackBinary(sb)
	utils.WriteSBf(sb, "")

	return sb.String()
}

func Or() string {
	sb := &strings.Builder{}
	utils.WriteSBf(sb, "// or\n")
	popBinary(sb)

	utils.WriteSBf(sb, "  D=D|M")

	writeBackBinary(sb)
	utils.WriteSBf(sb, "")

	return sb.String()
}

func Not() string {
	sb := &strings.Builder{}
	utils.WriteSBf(sb, "// not\n")
	popUnary(sb)

	utils.WriteSBf(sb, "  D=-D")
	utils.WriteSBf(sb, "  D=D-1")

	writeBackUnary(sb)
	utils.WriteSBf(sb, "")

	return sb.String()
}

func popBinary(sb *strings.Builder) *strings.Builder {
	utils.WriteSBf(sb, "  @SP")
	utils.WriteSBf(sb, "  A=M")
	utils.WriteSBf(sb, "  A=A-1")
	utils.WriteSBf(sb, "  D=M")
	utils.WriteSBf(sb, "  A=A-1")

	return sb
}

func popUnary(sb *strings.Builder) *strings.Builder {
	utils.WriteSBf(sb, "  @SP")
	utils.WriteSBf(sb, "  A=M")
	utils.WriteSBf(sb, "  A=A-1")
	utils.WriteSBf(sb, "  D=M")

	return sb
}

func writeBackBinary(sb *strings.Builder) *strings.Builder {
	utils.WriteSBf(sb, "  @SP")
	utils.WriteSBf(sb, "  A=M")
	utils.WriteSBf(sb, "  A=A-1")
	utils.WriteSBf(sb, "  A=A-1")
	utils.WriteSBf(sb, "  M=D")
	utils.WriteSBf(sb, "  D=A+1")
	utils.WriteSBf(sb, "  @SP")
	utils.WriteSBf(sb, "  M=D")

	return sb
}

func writeBackUnary(sb *strings.Builder) *strings.Builder {
	utils.WriteSBf(sb, "  @SP")
	utils.WriteSBf(sb, "  A=M")
	utils.WriteSBf(sb, "  A=A-1")
	utils.WriteSBf(sb, "  M=D")

	return sb
}
