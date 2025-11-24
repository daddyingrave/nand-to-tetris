package code

import (
	"fmt"
	"project7-vm1/parser"
	"strings"
)

func LineFromPush(segment string, index int, fileName string) (string, error) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("// push segment: %s, index %d\n\n", segment, index))

	if parser.SegmentPointer == segment && index == 0 {
		segment = parser.SegmentThis
	} else if parser.SegmentPointer == segment && index == 1 {
		segment = parser.SegmentThat
	}

	// 1. Read value from a memory segment
	switch segment {
	case parser.SegmentArgument, parser.SegmentLocal, parser.SegmentThis, parser.SegmentThat:
		sb.WriteString(fmt.Sprintf("@%d\n", index))
		sb.WriteString("D=A\n")
		sb.WriteString(parser.SegmentsMnemonics[segment] + "\n")
		sb.WriteString("A=D+M\n")
		sb.WriteString("D=M\n")

	case parser.SegmentStatic:
		staticLabel := strings.ReplaceAll(fileName, ".asm", fmt.Sprintf(".%d", index))
		staticLabel = fmt.Sprintf("@%s", staticLabel)
		sb.WriteString(staticLabel + "\n")
		sb.WriteString("D=M\n")

	case parser.SegmentConstant:
		sb.WriteString(fmt.Sprintf("@%d\n", index))
		sb.WriteString("D=A\n")

	case parser.SegmentTemp:
		sb.WriteString(fmt.Sprintf("@%d\n", index))
		sb.WriteString("D=A\n")
		sb.WriteString("@5\n")
		sb.WriteString("A=D+M\n")
		sb.WriteString("D=M\n")

	default:
		return "", fmt.Errorf("unknown segment type for push command '%s'", segment)
	}

	// 2. Write value form segment to stack
	sb.WriteString("@SP\n")
	sb.WriteString("A=M\n")
	sb.WriteString("M=D\n")

	// 3. Increment stack pointer
	sb.WriteString("D=A+1\n")
	sb.WriteString("@SP\n")
	sb.WriteString("M=D\n")
	sb.WriteString("\n")

	return sb.String(), nil
}

func LineFromPop(segment string, index int, fileName string) (string, error) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("// pop segment: %s, index %d\n\n", segment, index))

	if parser.SegmentPointer == segment && index == 0 {
		segment = parser.SegmentThis
	} else if parser.SegmentPointer == segment && index == 1 {
		segment = parser.SegmentThat
	}

	// 1. Define address in memory
	switch segment {
	case parser.SegmentArgument, parser.SegmentLocal, parser.SegmentThis, parser.SegmentThat:
		sb.WriteString(fmt.Sprintf("@%d\n", index))
		sb.WriteString("D=A\n")
		sb.WriteString(parser.SegmentsMnemonics[segment] + "\n")
		sb.WriteString("D=D+M\n")

	case parser.SegmentStatic:
		staticLabel := strings.ReplaceAll(fileName, ".asm", fmt.Sprintf(".%d", index))
		staticLabel = fmt.Sprintf("@%s", staticLabel)
		sb.WriteString(staticLabel + "\n")
		sb.WriteString("D=A\n")

	case parser.SegmentTemp:
		sb.WriteString(fmt.Sprintf("@%d\n", index))
		sb.WriteString("D=A\n")
		sb.WriteString("@5\n")
		sb.WriteString("D=D+M\n")

	default:
		return "", fmt.Errorf("unknown segment type for pop command '%s'", segment)
	}

	// 2. Save the current address for future writing
	sb.WriteString("@R13\n")
	sb.WriteString("M=D\n")

	// 4. Read values from the stack
	sb.WriteString("@SP\n")
	sb.WriteString("A=M-1\n")
	sb.WriteString("D=M\n")

	// 5. Write value to segment memory
	sb.WriteString("@R13\n")
	sb.WriteString("A=M\n")
	sb.WriteString("M=D\n")

	// 6. Decrement stack pointer
	sb.WriteString("@SP\n")
	sb.WriteString("M=M-1\n")

	sb.WriteString("\n")

	return sb.String(), nil
}
