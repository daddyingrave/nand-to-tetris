package parser

import (
	"errors"
	"strconv"
	"strings"
)

// 1   x   x   a   c   c    c   c   c   c   d   d   d   j   j   j
// 15  14  13  12  11  10   9   8   7   6   5   4   3   2   1   0
// 0   1   2   3   4   5    6   7   8   9   10  11  12  13  14  15

type InstructionType int

const (
	A InstructionType = iota
	C
	L
)

type Instruction struct {
	raw string
	typ InstructionType
}

func NewInstruction(line string) *Instruction {
	var typ InstructionType
	switch line[0] {
	case '@':
		typ = A
	case '(':
		typ = L
	default:
		typ = C
	}

	return &Instruction{
		raw: line,
		typ: typ,
	}
}

func (r *Instruction) Type() InstructionType {
	return r.typ
}

func (r *Instruction) Symbol() string {
	if r.typ == A {
		return strings.TrimPrefix(r.raw, "@")
	}
	if r.typ == L {
		return strings.Trim(r.raw, "()")
	}

	return ""
}

func (r *Instruction) IsNumber() (bool, int) {
	if r.typ == A {
		num, err := strconv.Atoi(r.Symbol())
		if err == nil {
			return true, num
		}
	}

	return false, 0
}

func (r *Instruction) Dest() string {
	if r.typ == C && strings.Contains(r.raw, "=") {
		split := strings.Split(r.raw, "=")
		return strings.TrimSpace(split[0])
	}

	return ""
}

func (r *Instruction) Comp() string {
	if r.typ == C {
		comp := r.raw
		if strings.Contains(r.raw, "=") {
			comp = strings.TrimSpace(strings.SplitAfter(comp, "=")[1])
		}
		if strings.Contains(r.raw, ";") {
			comp = strings.TrimSpace(strings.SplitAfter(comp, ";")[0])
			comp = comp[:len(comp)-1]
		}

		return comp
	}

	return ""
}

func (r *Instruction) Jump() string {
	if r.typ == C && strings.Contains(r.raw, ";") {
		split := strings.Split(r.raw, ";")
		return strings.TrimSpace(split[len(split)-1])
	}

	return ""
}

var ErrLineNotInstruction = errors.New("line is not an instruction")

func ParseLine(line string) (*Instruction, error) {
	trimmedLine := strings.TrimSpace(line)
	if strings.HasPrefix(trimmedLine, "//") {
		return nil, ErrLineNotInstruction
	}

	if len(trimmedLine) == 0 {
		return nil, ErrLineNotInstruction
	}

	lineAndComment := strings.SplitN(trimmedLine, "//", 2)
	instructionLine := strings.TrimSpace(lineAndComment[0])

	return NewInstruction(instructionLine), nil
}
