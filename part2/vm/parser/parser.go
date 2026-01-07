package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Parser interface {
	Commands(yield func(*Command, error) bool)
}

type Command struct {
	Type CommandType
	Arg1 string
	Arg2 int
}

type CommandType int

const (
	Arithmetic CommandType = iota
	Push
	Pop
	Label
	Goto
	IfGoto
	Function
	Return
	Call
)

func (r CommandType) String() string {
	switch r {
	case Arithmetic:
		return "arithmetic"
	case Push:
		return "push"
	case Pop:
		return "pop"
	case Label:
		return "label"
	case Goto:
		return "goto"
	case IfGoto:
		return "if-goto"
	case Function:
		return "function"
	case Return:
		return "return"
	case Call:
		return "call"
	default:
		return ""
	}
}

var ArithmeticalCommands = map[string]struct{}{
	"add": {}, "sub": {}, "neg": {},
	"eq": {}, "gt": {}, "lt": {},
	"and": {}, "or": {}, "not": {},
}

const (
	SegmentArgument = "argument"
	SegmentLocal    = "local"
	SegmentStatic   = "static"
	SegmentConstant = "constant"
	SegmentThis     = "this"
	SegmentThat     = "that"
	SegmentPointer  = "pointer"
	SegmentTemp     = "temp"
)

var SegmentsMnemonics = map[string]string{
	SegmentArgument: "@ARG",
	SegmentLocal:    "@LCL",
	SegmentThis:     "@THIS",
	SegmentThat:     "@THAT",
	SegmentTemp:     "@TEMP",
}

var pushes = map[string]struct{}{
	SegmentArgument: {},
	SegmentLocal:    {},
	SegmentStatic:   {},
	SegmentConstant: {},
	SegmentThis:     {},
	SegmentThat:     {},
	SegmentPointer:  {},
	SegmentTemp:     {},
}

var pops = map[string]struct{}{
	SegmentArgument: {},
	SegmentLocal:    {},
	SegmentStatic:   {},
	SegmentThis:     {},
	SegmentThat:     {},
	SegmentPointer:  {},
	SegmentTemp:     {},
}

type parser struct {
	path string
}

func NewParser(path string) Parser {
	return &parser{path: path}
}

func (r *parser) Commands(yield func(*Command, error) bool) {
	file, err := os.OpenFile(r.path, os.O_RDONLY, 0600)
	if err != nil {
		log.Fatalf("failed to open file '%s' with program code %s", r.path, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lineNumber := 0
	for scanner.Scan() {
		lineNumber++

		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		inlineCommentIndex := strings.Index(line, "//")
		if inlineCommentIndex > 0 {
			line = strings.TrimSpace(line[:inlineCommentIndex])
		}

		commandSplit := strings.Split(line, " ")
		switch {
		case len(commandSplit) == 3:
			typ := strings.TrimSpace(commandSplit[0])
			arg1 := strings.TrimSpace(commandSplit[1])
			arg2 := strings.TrimSpace(commandSplit[2])

			var actualType CommandType
			switch {
			case typ == Push.String():
				if _, exist := pushes[arg1]; !exist {
					yield(nil, fmt.Errorf("unsupported push segment '%s' on line '%s:%d'", arg1, r.path, lineNumber))
					return
				}
				actualType = Push
			case typ == Pop.String():
				if _, exist := pops[arg1]; !exist {
					yield(nil, fmt.Errorf("unsupported pop segment '%s' on line '%s:%d'", arg1, r.path, lineNumber))
					return
				}
				actualType = Pop
			case typ == Function.String():
				actualType = Function
			case typ == Call.String():
				actualType = Call
			default:
				yield(nil, fmt.Errorf("unexpected command type '%s' on line '%s:%d'", typ, r.path, lineNumber))
				return
			}

			arg2Typed, err := strconv.Atoi(arg2)
			if err != nil {
				yield(nil, fmt.Errorf("fail to convert arg2 to int on line '%s:%d' error: %w", r.path, lineNumber, err))
				return
			}

			cmd := &Command{
				Type: actualType,
				Arg1: arg1,
				Arg2: arg2Typed,
			}
			if !yield(cmd, nil) {
				return
			}

		case len(commandSplit) == 2:
			typ := strings.TrimSpace(commandSplit[0])
			arg1 := strings.TrimSpace(commandSplit[1])

			var actualType CommandType
			switch {
			case typ == Label.String():
				actualType = Label
			case typ == Goto.String():
				actualType = Goto
			case typ == IfGoto.String():
				actualType = IfGoto
			default:
				yield(nil, fmt.Errorf("unexpected command type '%s' on line '%s:%d'", typ, r.path, lineNumber))
				return
			}

			cmd := &Command{
				Type: actualType,
				Arg1: arg1,
			}
			if !yield(cmd, nil) {
				return
			}

		case len(commandSplit) == 1:
			cmdRaw := strings.TrimSpace(commandSplit[0])

			var actualType CommandType
			if cmdRaw == Return.String() {
				actualType = Return
			} else {
				if _, exist := ArithmeticalCommands[cmdRaw]; !exist {
					yield(nil, fmt.Errorf("unsupported command '%s' on line '%s:%d'", cmdRaw, r.path, lineNumber))
					return
				}
				actualType = Arithmetic
			}

			cmd := &Command{
				Type: actualType,
				Arg1: cmdRaw,
			}
			if !yield(cmd, nil) {
				return
			}

		default:
			yield(nil, fmt.Errorf(
				"can't parse line '%s:%d', expected exactly 1 or 3 parts, but given '%+v",
				r.path, lineNumber, commandSplit,
			))

			return
		}
	}

	return
}
