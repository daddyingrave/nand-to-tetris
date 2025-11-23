package code

import (
	"fmt"
	"project7-vm1/parser"
	"strings"
)

func LineFromPush(segment string, index int, fileName string, staticIndex int) (string, error) {
	sb := strings.Builder{}

	if parser.SegmentPointer == segment && index == 0 {
		segment = parser.SegmentThis
	} else if parser.SegmentPointer == segment && index == 1 {
		segment = parser.SegmentThat
	}

	// 1. Read value from a memory segment
	switch segment {
	case parser.SegmentArgument, parser.SegmentLocal, parser.SegmentThis, parser.SegmentThat:
		sb.WriteString(fmt.Sprintf("@%d", index))
		sb.WriteString("D=A")
		sb.WriteString(parser.SegmentsMnemonics[segment])
		sb.WriteString("A=D+M")
		sb.WriteString("D=M")

	case parser.SegmentStatic:
		staticLabel := strings.ReplaceAll(fileName, ".asm", fmt.Sprintf(".%d", staticIndex))
		staticLabel = fmt.Sprintf("@%s", staticLabel)
		sb.WriteString(staticLabel)
		sb.WriteString("D=M")

	case parser.SegmentConstant:
		sb.WriteString(fmt.Sprintf("@%d", index))
		sb.WriteString("D=A")

	case parser.SegmentTemp:
		sb.WriteString(fmt.Sprintf("@%d", index))
		sb.WriteString("D=A")
		sb.WriteString("@5")
		sb.WriteString("A=D+M")
		sb.WriteString("D=M")

	default:
		return "", fmt.Errorf("unknown segment type for push command '%s'", segment)
	}

	// 2. Write value form segment to stack
	sb.WriteString("@SP")
	sb.WriteString("A=M")
	sb.WriteString("M=D")

	// 3. Increment stack pointer
	sb.WriteString("D=A+1")
	sb.WriteString("@SP")
	sb.WriteString("M=D")

	return sb.String(), nil
}

func LineFromPop(segment string, index int, fileName string, staticIndex int) (string, error) {
	sb := strings.Builder{}

	if parser.SegmentPointer == segment && index == 0 {
		segment = parser.SegmentThis
	} else if parser.SegmentPointer == segment && index == 1 {
		segment = parser.SegmentThat
	}

	// 1. Define address in memory
	switch segment {
	case parser.SegmentArgument, parser.SegmentLocal, parser.SegmentThis, parser.SegmentThat:
		sb.WriteString(fmt.Sprintf("@%d", index))
		sb.WriteString("D=A")
		sb.WriteString(parser.SegmentsMnemonics[segment])
		sb.WriteString("D=D+M")

	case parser.SegmentStatic:
		staticLabel := strings.ReplaceAll(fileName, ".asm", fmt.Sprintf(".%d", staticIndex))
		staticLabel = fmt.Sprintf("@%s", staticLabel)
		sb.WriteString(staticLabel)
		sb.WriteString("D=A")

	case parser.SegmentTemp:
		sb.WriteString(fmt.Sprintf("@%d", index))
		sb.WriteString("D=A")
		sb.WriteString("@5")
		sb.WriteString("D=D+M")

	default:
		return "", fmt.Errorf("unknown segment type for pop command '%s'", segment)
	}

	// 2. Save the current address for future writing
	sb.WriteString("@R13")
	sb.WriteString("M=D")

	// 4. Read values from the stack
	sb.WriteString("@SP")
	sb.WriteString("A=M-1")
	sb.WriteString("D=M")

	// 5. Write value to segment memory
	sb.WriteString("@R13")
	sb.WriteString("A=M")
	sb.WriteString("M=D")

	// 6. Decrement stack pointer
	sb.WriteString("@SP")
	sb.WriteString("M=M-1")

	return sb.String(), nil
}
