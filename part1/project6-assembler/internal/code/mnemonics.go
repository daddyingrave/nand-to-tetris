package code

import (
	"assembler/internal/parser"
	"fmt"
	"strings"
)

func ConvertC(instruction *parser.Instruction) (string, error) {
	if instruction.Type() != parser.C {
		return "", fmt.Errorf("expected C instruction but given '%v'", instruction.Type())
	}

	comp := CompToBinary(instruction.Comp())
	dest := DestToBinary(instruction.Dest())
	jump := JumpToBinary(instruction.Jump())

	return fmt.Sprintf("111%s%s%s", comp, dest, jump), nil
}

func DestToBinary(mnemonic string) string {
	dest := []string{"0", "0", "0"}
	if strings.Contains(mnemonic, "M") {
		dest[2] = "1"
	}
	if strings.Contains(mnemonic, "D") {
		dest[1] = "1"
	}
	if strings.Contains(mnemonic, "A") {
		dest[0] = "1"
	}

	return strings.Join(dest, "")
}

var compTranslations = map[string]string{
	// a = 0
	"0":   "0101010",
	"1":   "0111111",
	"-1":  "0111010",
	"D":   "0001100",
	"A":   "0110000",
	"!D":  "0001101",
	"!A":  "0110001",
	"-D":  "0001111",
	"-A":  "0110011",
	"D+1": "0011111",
	"A+1": "0110111",
	"D-1": "0001110",
	"A-1": "0110010",
	"D+A": "0000010",
	"D-A": "0010011",
	"A-D": "0000111",
	"D&A": "0000000",
	"D|A": "0010101",

	// a = 1
	"M":   "1110000",
	"!M":  "1110001",
	"-M":  "1110011",
	"M+1": "1110111",
	"M-1": "1110010",
	"D+M": "1000010",
	"D-M": "1010011",
	"M-D": "1000111",
	"D&M": "1000000",
	"D|M": "1010101",
}

func CompToBinary(mnemonic string) string {
	return compTranslations[mnemonic]
}

func JumpToBinary(mnemonic string) string {
	jump := []string{"0", "0", "0"}

	switch mnemonic {
	case "JGT":
		jump[2] = "1"
	case "JEQ":
		jump[1] = "1"
	case "JGE":
		jump[1] = "1"
		jump[2] = "1"
	case "JLT":
		jump[0] = "1"
	case "JNE":
		jump[0] = "1"
		jump[2] = "1"
	case "JLE":
		jump[0] = "1"
		jump[1] = "1"
	case "JMP":
		jump[0] = "1"
		jump[1] = "1"
		jump[2] = "1"
	}

	return strings.Join(jump, "")
}
