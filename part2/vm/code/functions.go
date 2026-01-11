package code

import (
	"fmt"
	"strings"
	"vm/internal/utils"
	"vm/parser"
)

// Call

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

		//pop, err := StackPop(parser.SegmentLocal, i, fileName)
		//if err != nil {
		//	return "", fmt.Errorf("fail to build pop local for function '%s/%d', in file '%s' %w", fnName, nVars, fileName, err)
		//}
		//utils.WriteSBf(sb, pop)
	}

	return sb.String(), nil
}

var callerSegments = map[string]int{
	"THAT": 1,
	"THIS": 2,
	"ARG":  3,
	"LCL":  4,
}

func Return(fileName string, function string) (string, error) {
	sb := &strings.Builder{}

	utils.WriteSBf(sb, "// return for function: %s, file: %s\n", function, fileName)

	// endFrame
	utils.WriteSBf(sb, "// endFrame")
	utils.WriteSBf(sb, "  @LCL")
	utils.WriteSBf(sb, "  D=M")
	utils.WriteSBf(sb, "  @endFrame")
	utils.WriteSBf(sb, "  M=D")
	utils.WriteSBf(sb, "")

	// retAddr
	utils.WriteSBf(sb, "// retAddr")
	utils.WriteSBf(sb, "  @5")
	utils.WriteSBf(sb, "  D=D-A") // D preserved from endFrame
	utils.WriteSBf(sb, "  D=M")
	utils.WriteSBf(sb, "  @retAddr")
	utils.WriteSBf(sb, "  M=D")
	utils.WriteSBf(sb, "")

	// reposition return val
	utils.WriteSBf(sb, "// reposition return val")
	utils.WriteSBf(sb, "  @SP")
	utils.WriteSBf(sb, "  AM=M-1")
	utils.WriteSBf(sb, "  D=M")
	utils.WriteSBf(sb, "  @ARG")
	utils.WriteSBf(sb, "  M=D")
	utils.WriteSBf(sb, "")

	// reposition SP for caller
	utils.WriteSBf(sb, "// reposition SP for caller")
	utils.WriteSBf(sb, "  @ARG")
	utils.WriteSBf(sb, "  D=M+1")
	utils.WriteSBf(sb, "  @SP")
	utils.WriteSBf(sb, "  M=D")
	utils.WriteSBf(sb, "")

	// restore caller segments
	utils.WriteSBf(sb, "// restore caller segments")
	for segment, offset := range callerSegments {
		utils.WriteSBf(sb, "  @%d", offset)
		utils.WriteSBf(sb, "  D=A")
		utils.WriteSBf(sb, "  @endFrame")
		utils.WriteSBf(sb, "  D=M-D")
		utils.WriteSBf(sb, "  @%s", segment)
		utils.WriteSBf(sb, "  M=D")
		utils.WriteSBf(sb, "")
	}

	// goto retAddr
	utils.WriteSBf(sb, "// goto retAddr")
	utils.WriteSBf(sb, "  @retAddr")
	utils.WriteSBf(sb, "  A=M")
	utils.WriteSBf(sb, "  0;JMP")
	utils.WriteSBf(sb, "")

	return sb.String(), nil
}
