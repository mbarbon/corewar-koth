package redcode

import (
	"errors"
	"regexp"
	"strings"
)

//go:generate ragel -Z -G2 -o lex.go redcode.rl
//go:generate goyacc redcode.y

// Directives map Redcode directive names to values
//
// e.g. ";name Imp"
type Directives map[string]string

var scanDirective *regexp.Regexp

func init() {
	var err error
	scanDirective, err = regexp.Compile("^\\s*(name|author)\\s+(.*)")
	if err != nil {
		panic(err)
	}
}

// ParseString parses a Redcode program
func ParseString(text, filename string) ([]Instruction, Directives, error) {
	return ParseBytes([]byte(text), filename)
}

// ParseBytes parses a Redcode program
func ParseBytes(text []byte, filename string) ([]Instruction, Directives, error) {
	lex := newLexer(text, filename)
	e := yyParse(lex)
	if lex.err != nil {
		return nil, nil, lex.err
	} else if e != 0 {
		return nil, nil, errors.New("Unknown error during parsing")
	}
	for index, instruction := range lex.instructions {
		var err error
		lex.instructions[index], err = checkInstruction(instruction)
		if err != nil {
			return nil, nil, err
		}
	}
	return lex.instructions, lex.directives, nil
}

func checkInstruction(instruction Instruction) (Instruction, error) {
	switch instruction.Opcode {
	case OpDat:
		if !hasB(instruction) {
			instruction.A, instruction.B = instruction.B, instruction.A
		}
	}
	return instruction, nil
}

func hasA(instruction Instruction) bool {
	return instruction.A.Expression != nil
}

func hasB(instruction Instruction) bool {
	return instruction.B.Expression != nil
}

func parseDirective(lexer *lexer, comment string) {
	parts := scanDirective.FindStringSubmatch(comment)
	if parts == nil {
		return
	}
	lexer.directives[parts[1]] = strings.Trim(parts[2], " \t")
}
