package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"vm/parser"
)

type Writer interface {
	WriteArithmetic(command string) error
	WritePushPop(commandType parser.CommandType, segment string, index int) error
	WriteLabel(label string) error
	WriteGoTo(label string) error
	WriteIfGoTo(label string) error
	WriteFunction(functionName string, nVars int) error
	WriteCall(functionName string, nArgs int) error
	WriteReturn() error
}

type writer struct {
	asmFile         *os.File
	vmFileName      string
	currentFunction string
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

func NewWriter(asmFile *os.File, vmFileName string) (Writer, error) {
	return writer{
		asmFile:         asmFile,
		vmFileName:      strings.TrimSuffix(vmFileName, ".vm"),
		currentFunction: "global",
	}, nil
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

	_, vmFileNameOnly := filepath.Split(r.vmFileName)
	if segment == parser.SegmentStatic {
		staticMappings[vmFileNameOnly]++
	}

	switch commandType {
	case parser.Push:
		generatedCode, err = StackPush(segment, index, vmFileNameOnly)
	case parser.Pop:
		generatedCode, err = StackPop(segment, index, vmFileNameOnly)
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

func (r writer) WriteLabel(label string) error {
	_, fn := filepath.Split(r.vmFileName)
	fileName := strings.ReplaceAll(fn, ".asm", "")

	if _, err := r.asmFile.WriteString(Label(fileName, r.currentFunction, label)); err != nil {
		return fmt.Errorf("fail to write command label '%s', to file '%s' %w", label, r.asmFile.Name(), err)
	}

	return nil
}

func (r writer) WriteGoTo(label string) error {
	_, fn := filepath.Split(r.vmFileName)
	fileName := strings.ReplaceAll(fn, ".asm", "")

	if _, err := r.asmFile.WriteString(GoTo(fileName, r.currentFunction, label)); err != nil {
		return fmt.Errorf("fail to write command goto '%s', to file '%s' %w", label, r.asmFile.Name(), err)
	}

	return nil
}

func (r writer) WriteIfGoTo(label string) error {
	_, fn := filepath.Split(r.vmFileName)
	fileName := strings.ReplaceAll(fn, ".asm", "")

	if _, err := r.asmFile.WriteString(IfGoTo(fileName, r.currentFunction, label)); err != nil {
		return fmt.Errorf("fail to write command if-goto '%s', to file '%s' %w", label, r.asmFile.Name(), err)
	}

	return nil
}

func (r writer) WriteFunction(functionName string, nVars int) error {
	//TODO implement me
	panic("implement me")
}

func (r writer) WriteCall(functionName string, nArgs int) error {
	//TODO implement me
	panic("implement me")
}

func (r writer) WriteReturn() error {
	//TODO implement me
	panic("implement me")
}
