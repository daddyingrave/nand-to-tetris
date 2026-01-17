package code

import (
	"fmt"
	"slices"
	"strings"
	"vm/internal/utils"
	"vm/parser"
)

func Call(fileName string, currentFunction string, functionToCall string, nArgs int, i int) (string, error) {
	sb := &strings.Builder{}

	utils.WriteSBf(sb, "// call: %s from %s, nArgs %d, file: %s\n", functionToCall, currentFunction, nArgs, fileName)

	returnLabelName := fmt.Sprintf("%s.%s.$ret.%d", fileName, currentFunction, i)

	utils.WriteSBf(sb, "// return address label")
	utils.WriteSBf(sb, "  @%s", returnLabelName)
	utils.WriteSBf(sb, "  D=A")
	utils.WriteSBf(sb, "  @SP")
	utils.WriteSBf(sb, "  A=M")
	utils.WriteSBf(sb, "  M=D")
	utils.WriteSBf(sb, "  D=A+1")
	utils.WriteSBf(sb, "  @SP")
	utils.WriteSBf(sb, "  M=D")
	utils.WriteSBf(sb, "")

	utils.WriteSBf(sb, "// save callers segments")
	for _, segment := range slices.Backward(callerSegments) {
		utils.WriteSBf(sb, "// push %s", segment)
		utils.WriteSBf(sb, "  @%s", segment)
		utils.WriteSBf(sb, "  D=M")
		utils.WriteSBf(sb, "  @SP")
		utils.WriteSBf(sb, "  A=M")
		utils.WriteSBf(sb, "  M=D")
		utils.WriteSBf(sb, "  D=A+1")
		utils.WriteSBf(sb, "  @SP")
		utils.WriteSBf(sb, "  M=D")
		utils.WriteSBf(sb, "")
	}

	utils.WriteSBf(sb, "// Reposition ARG")
	utils.WriteSBf(sb, "  @SP")
	utils.WriteSBf(sb, "  D=M")
	utils.WriteSBf(sb, "  @5")
	utils.WriteSBf(sb, "  D=D-A")
	utils.WriteSBf(sb, "  @%d", nArgs)
	utils.WriteSBf(sb, "  D=D-A")
	utils.WriteSBf(sb, "  @ARG")
	utils.WriteSBf(sb, "  M=D")
	utils.WriteSBf(sb, "")

	utils.WriteSBf(sb, "// Reposition LCL")
	utils.WriteSBf(sb, "  @SP")
	utils.WriteSBf(sb, "  D=M")
	utils.WriteSBf(sb, "  @LCL")
	utils.WriteSBf(sb, "  M=D")
	utils.WriteSBf(sb, "")

	utils.WriteSBf(sb, "// goto function")
	utils.WriteSBf(sb, "  @%s", functionToCall)
	utils.WriteSBf(sb, "  0;JMP")
	utils.WriteSBf(sb, "(%s)", returnLabelName)
	utils.WriteSBf(sb, "")

	return sb.String(), nil
}

func Function(fileName string, fnName string, nVars int) (string, error) {
	sb := &strings.Builder{}

	utils.WriteSBf(sb, "// function: %s, nVars %d, file: %s\n", fnName, nVars, fileName)
	utils.WriteSBf(sb, "(%s)\n", fnName)

	for range nVars {
		push, err := StackPush(parser.SegmentConstant, 0, fileName)
		if err != nil {
			return "", fmt.Errorf("fail to build push constant for function '%s/%d', in file '%s' %w", fnName, nVars, fileName, err)
		}
		utils.WriteSBf(sb, push)
	}

	return sb.String(), nil
}

var callerSegments = []string{
	"THAT",
	"THIS",
	"ARG",
	"LCL",
}

func Return(fileName string, function string) (string, error) {
	sb := &strings.Builder{}

	utils.WriteSBf(sb, "// return for function: %s, file: %s\n", function, fileName)

	utils.WriteSBf(sb, "// endFrame")
	utils.WriteSBf(sb, "  @LCL")
	utils.WriteSBf(sb, "  D=M")
	utils.WriteSBf(sb, "  @endFrame")
	utils.WriteSBf(sb, "  M=D")
	utils.WriteSBf(sb, "")

	utils.WriteSBf(sb, "// retAddr")
	utils.WriteSBf(sb, "  @5")
	utils.WriteSBf(sb, "  D=D-A") // D preserved from endFrame
	utils.WriteSBf(sb, "  A=D")   // Dereference return address
	utils.WriteSBf(sb, "  D=M")
	utils.WriteSBf(sb, "  @retAddr")
	utils.WriteSBf(sb, "  M=D")
	utils.WriteSBf(sb, "")

	utils.WriteSBf(sb, "// reposition return val")
	utils.WriteSBf(sb, "  @SP")
	utils.WriteSBf(sb, "  AM=M-1")
	utils.WriteSBf(sb, "  D=M")
	utils.WriteSBf(sb, "  @ARG")
	utils.WriteSBf(sb, "  A=M")
	utils.WriteSBf(sb, "  M=D")
	utils.WriteSBf(sb, "")

	utils.WriteSBf(sb, "// reposition SP for caller")
	utils.WriteSBf(sb, "  @ARG")
	utils.WriteSBf(sb, "  D=M+1")
	utils.WriteSBf(sb, "  @SP")
	utils.WriteSBf(sb, "  M=D")
	utils.WriteSBf(sb, "")

	utils.WriteSBf(sb, "// restore caller segments")
	for offset, segment := range callerSegments {
		utils.WriteSBf(sb, "  @%d", offset+1)
		utils.WriteSBf(sb, "  D=A")
		utils.WriteSBf(sb, "  @endFrame")
		utils.WriteSBf(sb, "  D=M-D")
		utils.WriteSBf(sb, "  A=D")
		utils.WriteSBf(sb, "  D=M")
		utils.WriteSBf(sb, "  @%s", segment)
		utils.WriteSBf(sb, "  M=D")
		utils.WriteSBf(sb, "")
	}

	utils.WriteSBf(sb, "// goto retAddr")
	utils.WriteSBf(sb, "  @retAddr")
	utils.WriteSBf(sb, "  A=M")
	utils.WriteSBf(sb, "  0;JMP")
	utils.WriteSBf(sb, "")

	return sb.String(), nil
}
