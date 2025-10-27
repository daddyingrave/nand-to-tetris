package assembler

import (
	"assembler/internal/code"
	"assembler/internal/parser"
	"bufio"
	"errors"
	"fmt"
	"log"
	"maps"
	"os"
	"strings"
)

func Translate(path string) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		log.Fatalf("failed to open file '%s' with program code %s", path, err)
	}

	var instructions []*parser.Instruction
	symbols := map[string]int{}
	maps.Copy(symbols, code.PredefinedSymbols)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		instruction, err := parser.ParseLine(scanner.Text())
		if err != nil {
			if errors.Is(err, parser.ErrLineNotInstruction) {
				continue
			}
			log.Fatalf("fail to parse program %s", err)
		}

		if instruction.Type() == parser.L {
			symbols[instruction.Symbol()] = len(instructions)
		} else {
			instructions = append(instructions, instruction)
		}
	}

	targetPath := strings.Replace(path, ".asm", ".hack", 1)
	binaryFile, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("fail to open target file '%s' %s", targetPath, err)
	}
	addressToAllocate := 16

	for i, instruction := range instructions {
		var binaryInstruction string
		if instruction.Type() == parser.A {
			var address int
			isNumber, num := instruction.IsNumber()
			if isNumber {
				address = num
			} else {
				symbolAddress, exist := symbols[instruction.Symbol()]
				if !exist {
					symbols[instruction.Symbol()] = addressToAllocate
					symbolAddress = addressToAllocate
					addressToAllocate++
				}

				address = symbolAddress
			}

			binaryInstruction = fmt.Sprintf("%016b", address)
		} else {
			binaryC, err := code.ConvertC(instruction)
			if err != nil {
				log.Fatalf("failed to convert C instruction %s", err)
			}

			binaryInstruction = binaryC
		}

		if i != len(instructions)-1 {
			binaryInstruction += "\n"
		}
		_, err := binaryFile.WriteString(binaryInstruction)
		if err != nil {
			log.Fatalf("fail to write instruction '%s' to the target file", binaryInstruction)
		}
	}
}
