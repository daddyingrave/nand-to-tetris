package code

import (
	"fmt"
	"io"
	"os"
	"project7-vm1/parser"
	"strings"
)

type Writer interface {
	WriteArithmetic(command string)
	WritePushPop(commandType parser.CommandType, segment string, index int)
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

func (r writer) WriteArithmetic(command string) {
	//TODO implement me
	panic("implement me")
}

func (r writer) WritePushPop(commandType parser.CommandType, segment string, index int) {
	//TODO implement me
	panic("implement me")
}

func (r writer) Close() error {
	return r.asmFile.Close()
}
