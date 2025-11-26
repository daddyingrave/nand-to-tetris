package code

import (
	"fmt"
	"io"
	"os"
	"project7-vm1/parser"
	"strings"
)

type Writer interface {
	WriteArithmetic(command string) error
	WritePushPop(commandType parser.CommandType, segment string, index int) error
	io.Closer
}

type writer struct {
	asmFile *os.File
}

// 0 - stack pointer (SP)
// 1 - base address of local segment (LCL)
// 2 - base address of argument segment (ARG)
// 3 - base address of this segment (THIS)
// 4 - base address of that segment (THAT)
// 5 - 12 temp segment (TEMP)
// 13 - 15 could be used as variables in generated assembly code

// 16 - 255 - static variables
// 256 - 2047 - stack

func NewWriter(path string) (Writer, error) {
	targetPath := strings.Replace(path, ".vm", ".asm", 1)
	asmFile, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return nil, fmt.Errorf("fail to open target file '%s' %w", targetPath, err)
	}

	return writer{asmFile: asmFile}, nil
}

func (r writer) WriteArithmetic(command string) error {
	writeF := Mappings[command]
	if writeF == nil {
		return fmt.Errorf("unknown command '%s'", command)
	}

	if _, err := r.asmFile.WriteString(writeF()); err != nil {
		return fmt.Errorf("fail to write command '%s' to file %w", command, err)
	}

	return nil
}

var staticMappings = map[string]int{}

func (r writer) WritePushPop(commandType parser.CommandType, segment string, index int) error {
	var err error
	var generatedCode string

	if segment == parser.SegmentStatic {
		staticMappings[r.asmFile.Name()]++
	}

	switch commandType {
	case parser.Push:
		generatedCode, err = StackPush(segment, index, r.asmFile.Name())
	case parser.Pop:
		generatedCode, err = StackPop(segment, index, r.asmFile.Name())
	default:
		return fmt.Errorf("unsupported command type '%s'", commandType.String())
	}

	if err != nil {
		return fmt.Errorf(
			"fail to write command '%s', segment '%s', index '%d', to file '%s'",
			commandType.String(), segment, index, r.asmFile.Name(),
		)
	}

	if _, err := r.asmFile.WriteString(generatedCode); err != nil {
		return fmt.Errorf(
			"fail to write command '%s' segment '%s', index '%d', to file '%s' %w",
			commandType.String(), segment, index, r.asmFile.Name(), err,
		)
	}

	return nil
}

func (r writer) Close() error {
	return r.asmFile.Close()
}
