package code

import (
	"fmt"
	"strings"
	"vm/parser"
)

func StackPush(segment string, index int, fileName string) (string, error) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("// push segment: %s, index %d\n\n", segment, index))

	// 1. Read value from a memory segment
	switch segment {
	case parser.SegmentPointer:
		if index == 0 {
			sb.WriteString("  " + parser.SegmentsMnemonics[parser.SegmentThis] + "\n")
		} else if index == 1 {
			sb.WriteString("  " + parser.SegmentsMnemonics[parser.SegmentThat] + "\n")
		} else {
			return "", fmt.Errorf("impossible index '%d' for pointer segment in file '%s'", index, fileName)
		}
		sb.WriteString("  D=M\n")

	case parser.SegmentArgument, parser.SegmentLocal, parser.SegmentThis, parser.SegmentThat:
		sb.WriteString("  " + parser.SegmentsMnemonics[segment] + "\n")
		sb.WriteString("  D=M\n")
		sb.WriteString(fmt.Sprintf("  @%d\n", index))
		sb.WriteString("  A=D+A\n")
		sb.WriteString("  D=M\n")

	case parser.SegmentStatic:
		staticLabel := strings.ReplaceAll(fileName, ".asm", fmt.Sprintf(".%d", index))
		staticLabel = fmt.Sprintf("  @%s", staticLabel)
		sb.WriteString(staticLabel + "\n")
		sb.WriteString("  D=M\n")

	case parser.SegmentConstant:
		sb.WriteString(fmt.Sprintf("  @%d\n", index))
		sb.WriteString("  D=A\n")

	case parser.SegmentTemp:
		sb.WriteString(fmt.Sprintf("  @%d\n", index))
		sb.WriteString("  D=A\n")
		sb.WriteString("  @5\n")
		sb.WriteString("  A=D+A\n")
		sb.WriteString("  D=M\n")

	default:
		return "", fmt.Errorf("unknown segment type for push command '%s'", segment)
	}

	// 2. Write value form segment to stack
	sb.WriteString("  @SP\n")
	sb.WriteString("  A=M\n")
	sb.WriteString("  M=D\n")

	// 3. Increment stack pointer
	sb.WriteString("  D=A+1\n")
	sb.WriteString("  @SP\n")
	sb.WriteString("  M=D\n")
	sb.WriteString("\n")

	return sb.String(), nil
}

func StackPop(segment string, index int, fileName string) (string, error) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("// pop segment: %s, index %d\n\n", segment, index))

	// 1. Define address in memory
	switch segment {
	case parser.SegmentPointer:
		if index == 0 {
			sb.WriteString("  " + parser.SegmentsMnemonics[parser.SegmentThis] + "\n")
		} else if index == 1 {
			sb.WriteString("  " + parser.SegmentsMnemonics[parser.SegmentThat] + "\n")
		} else {
			return "", fmt.Errorf("impossible index '%d' for pointer segment in file '%s'", index, fileName)
		}
		sb.WriteString("  D=A\n")

	case parser.SegmentArgument, parser.SegmentLocal, parser.SegmentThis, parser.SegmentThat:
		sb.WriteString("  " + parser.SegmentsMnemonics[segment] + "\n")
		sb.WriteString("  D=M\n")
		sb.WriteString(fmt.Sprintf("  @%d\n", index))
		sb.WriteString("  D=D+A\n")

	case parser.SegmentStatic:
		staticLabel := strings.ReplaceAll(fileName, ".asm", fmt.Sprintf(".%d", index))
		staticLabel = fmt.Sprintf("  @%s", staticLabel)
		sb.WriteString(staticLabel + "\n")
		sb.WriteString("  D=A\n")

	case parser.SegmentTemp:
		sb.WriteString(fmt.Sprintf("  @%d\n", index))
		sb.WriteString("  D=A\n")
		sb.WriteString("  @5\n")
		sb.WriteString("  D=D+A\n")

	default:
		return "", fmt.Errorf("unknown segment type for pop command '%s'", segment)
	}

	// 2. Save the current address for future writing
	sb.WriteString("  @R13\n")
	sb.WriteString("  M=D\n")

	// 4. Read values from the stack
	sb.WriteString("  @SP\n")
	sb.WriteString("  A=M-1\n")
	sb.WriteString("  D=M\n")

	// 5. Write value to segment memory
	sb.WriteString("  @R13\n")
	sb.WriteString("  A=M\n")
	sb.WriteString("  M=D\n")

	// 6. Decrement stack pointer
	sb.WriteString("  @SP\n")
	sb.WriteString("  M=M-1\n")

	sb.WriteString("\n")

	return sb.String(), nil
}
