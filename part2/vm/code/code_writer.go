package code

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"vm/parser"
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

var Labelled = map[string]int{
	"eq": 0,
	"gt": 0,
	"lt": 0,
}

func (r writer) WriteArithmetic(command string) error {
	count, exist := Labelled[command]
	if exist {
		count++
		Labelled[command] = count
	}

	var code string

	switch command {
	case "add":
		code = Add()
	case "sub":
		code = Sub()
	case "neg":
		code = Neg()
	case "eq":
		code = Eq(count)
	case "gt":
		code = Gt(count)
	case "lt":
		code = Lt(count)
	case "and":
		code = And()
	case "or":
		code = Or()
	case "not":
		code = Not()
	default:
		return fmt.Errorf("unknown command '%s'", command)
	}

	if _, err := r.asmFile.WriteString(code); err != nil {
		return fmt.Errorf("fail to write command '%s' to file %w", command, err)
	}

	return nil
}

var staticMappings = map[string]int{}

func (r writer) WritePushPop(commandType parser.CommandType, segment string, index int) error {
	var err error
	var generatedCode string

	asmFileName := r.asmFile.Name()
	_, asmFileName = filepath.Split(asmFileName)
	if segment == parser.SegmentStatic {
		staticMappings[asmFileName]++
	}

	switch commandType {
	case parser.Push:
		generatedCode, err = StackPush(segment, index, asmFileName)
	case parser.Pop:
		generatedCode, err = StackPop(segment, index, asmFileName)
	default:
		return fmt.Errorf("unsupported command type '%s'", commandType.String())
	}

	if err != nil {
		return fmt.Errorf(
			"fail to write command '%s', segment '%s', index '%d', to file '%s'",
			commandType.String(), segment, index, asmFileName,
		)
	}

	if _, err := r.asmFile.WriteString(generatedCode); err != nil {
		return fmt.Errorf(
			"fail to write command '%s' segment '%s', index '%d', to file '%s' %w",
			commandType.String(), segment, index, asmFileName, err,
		)
	}

	return nil
}

func (r writer) Close() error {
	return r.asmFile.Close()
}
