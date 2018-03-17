package redcode

import (
	"errors"
)

//go:generate ragel -Z -G2 -o lex.go redcode.rl
//go:generate goyacc redcode.y

// ParseString parses a Redcode program
func ParseString(text string) ([]Instruction, error) {
	return ParseBytes([]byte(text))
}

// ParseBytes parses a Redcode program
func ParseBytes(text []byte) ([]Instruction, error) {
	lex := newLexer(text)
	e := yyParse(lex)
	if lex.err != nil {
		return nil, lex.err
	} else if e != 0 {
		return nil, errors.New("Unknown error during parsing")
	}
	for index, instruction := range lex.instructions {
		var err error
		lex.instructions[index], err = checkInstruction(instruction)
		if err != nil {
			return nil, err
		}
	}
	return lex.instructions, nil
}

func checkInstruction(instruction Instruction) (Instruction, error) {
	switch instruction.Opcode {
	case OpSpl:
		if hasB(instruction) {
			return instruction, errors.New("Invalid 2-operand SPL")
		}
		// SPL is weird in '86
		instruction.B = instruction.A
		instruction.A = Operand{}
		return instruction, nil
	}
	return instruction, nil
}

func hasA(instruction Instruction) bool {
	return instruction.A.Expression != nil
}

func hasB(instruction Instruction) bool {
	return instruction.B.Expression != nil
}
