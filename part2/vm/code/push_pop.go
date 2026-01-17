package code

import (
	"fmt"
	"strings"
	"vm/internal/utils"
	"vm/parser"
)

func StackPush(segment string, index int, fileName string) (string, error) {
	sb := &strings.Builder{}
	utils.WriteSBf(sb, "// push segment: %s, index %d", segment, index)
	utils.WriteSBf(sb, "")

	// 1. Read value from a memory segment
	switch segment {
	case parser.SegmentPointer:
		if index == 0 {
			utils.WriteSBf(sb, "  %s", parser.SegmentsMnemonics[parser.SegmentThis])
		} else if index == 1 {
			utils.WriteSBf(sb, "  %s", parser.SegmentsMnemonics[parser.SegmentThat])
		} else {
			return "", fmt.Errorf("impossible index '%d' for pointer segment in file '%s'", index, fileName)
		}
		utils.WriteSBf(sb, "  D=M")

	case parser.SegmentArgument, parser.SegmentLocal, parser.SegmentThis, parser.SegmentThat:
		utils.WriteSBf(sb, "  %s", parser.SegmentsMnemonics[segment])
		utils.WriteSBf(sb, "  D=M")
		utils.WriteSBf(sb, "  @%d", index)
		utils.WriteSBf(sb, "  A=D+A")
		utils.WriteSBf(sb, "  D=M")

	case parser.SegmentStatic:
		utils.WriteSBf(sb, "  @%s.%d", fileName, index)
		utils.WriteSBf(sb, "  D=M")

	case parser.SegmentConstant:
		utils.WriteSBf(sb, "  @%d", index)
		utils.WriteSBf(sb, "  D=A")

	case parser.SegmentTemp:
		utils.WriteSBf(sb, "  @%d", index)
		utils.WriteSBf(sb, "  D=A")
		utils.WriteSBf(sb, "  @5")
		utils.WriteSBf(sb, "  A=D+A")
		utils.WriteSBf(sb, "  D=M")

	default:
		return "", fmt.Errorf("unknown segment type for push command '%s'", segment)
	}

	// 2. Write value form segment to stack
	utils.WriteSBf(sb, "  @SP")
	utils.WriteSBf(sb, "  A=M")
	utils.WriteSBf(sb, "  M=D")

	// 3. Increment stack pointer
	utils.WriteSBf(sb, "  D=A+1")
	utils.WriteSBf(sb, "  @SP")
	utils.WriteSBf(sb, "  M=D")
	utils.WriteSBf(sb, "")

	return sb.String(), nil
}

func StackPop(segment string, index int, fileName string) (string, error) {
	sb := &strings.Builder{}
	utils.WriteSBf(sb, "// pop segment: %s, index %d", segment, index)
	utils.WriteSBf(sb, "")

	// 1. Define address in memory
	switch segment {
	case parser.SegmentPointer:
		if index == 0 {
			utils.WriteSBf(sb, "  %s", parser.SegmentsMnemonics[parser.SegmentThis])
		} else if index == 1 {
			utils.WriteSBf(sb, "  %s", parser.SegmentsMnemonics[parser.SegmentThat])
		} else {
			return "", fmt.Errorf("impossible index '%d' for pointer segment in file '%s'", index, fileName)
		}
		utils.WriteSBf(sb, "  D=A")

	case parser.SegmentArgument, parser.SegmentLocal, parser.SegmentThis, parser.SegmentThat:
		utils.WriteSBf(sb, "  %s", parser.SegmentsMnemonics[segment])
		utils.WriteSBf(sb, "  D=M")
		utils.WriteSBf(sb, "  @%d", index)
		utils.WriteSBf(sb, "  D=D+A")

	case parser.SegmentStatic:
		utils.WriteSBf(sb, "  @%s.%d", fileName, index)
		utils.WriteSBf(sb, "  D=A")

	case parser.SegmentTemp:
		utils.WriteSBf(sb, "  @%d", index)
		utils.WriteSBf(sb, "  D=A")
		utils.WriteSBf(sb, "  @5")
		utils.WriteSBf(sb, "  D=D+A")

	default:
		return "", fmt.Errorf("unknown segment type for pop command '%s'", segment)
	}

	// 2. Save the current address for future writing
	utils.WriteSBf(sb, "  @R13")
	utils.WriteSBf(sb, "  M=D")

	// 4. Read values from the stack
	utils.WriteSBf(sb, "  @SP")
	utils.WriteSBf(sb, "  A=M-1")
	utils.WriteSBf(sb, "  D=M")

	// 5. Write value to segment memory
	utils.WriteSBf(sb, "  @R13")
	utils.WriteSBf(sb, "  A=M")
	utils.WriteSBf(sb, "  M=D")

	// 6. Decrement stack pointer
	utils.WriteSBf(sb, "  @SP")
	utils.WriteSBf(sb, "  M=M-1")

	utils.WriteSBf(sb, "")

	return sb.String(), nil
}
