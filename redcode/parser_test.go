package redcode

import (
	"testing"

	"github.com/go-test/deep"
)

func TestSanity(t *testing.T) {
	checkInstructions(t, " MOV #7, -1\n", Instruction{
		Opcode: OpMov,
		A:      Operand{Mode: Immediate, Expression: constNum(7)},
		B:      Operand{Mode: Relative, Expression: constNum(-1)},
	})

	checkInstructions(t, "mov #7, -1", Instruction{
		Opcode: OpMov,
		A:      Operand{Mode: Immediate, Expression: constNum(7)},
		B:      Operand{Mode: Relative, Expression: constNum(-1)},
	})
}

func TestComments(t *testing.T) {
	checkInstructions(t, " ; comment")

	checkInstructions(t, "mov #7, -1 ; comment", Instruction{
		Opcode: OpMov,
		A:      Operand{Mode: Immediate, Expression: constNum(7)},
		B:      Operand{Mode: Relative, Expression: constNum(-1)},
	})

	checkInstructions(t, "mov #7, -1\n; comment\nmov #8, 1", Instruction{
		Opcode: OpMov,
		A:      Operand{Mode: Immediate, Expression: constNum(7)},
		B:      Operand{Mode: Relative, Expression: constNum(-1)},
	}, Instruction{
		Opcode: OpMov,
		A:      Operand{Mode: Immediate, Expression: constNum(8)},
		B:      Operand{Mode: Relative, Expression: constNum(1)},
	})
}

func TestDirectives(t *testing.T) {
	checkDirectives(t, ";name Imp \n; author A.K. Dewdney", Directives{
		"name":   "Imp",
		"author": "A.K. Dewdney",
	})
}

func TestLabel(t *testing.T) {
	checkInstructions(t, "imp mov 0, 1", Instruction{
		Label:  "imp",
		Opcode: OpMov,
		A:      Operand{Mode: Relative, Expression: constNum(0)},
		B:      Operand{Mode: Relative, Expression: constNum(1)},
	})
}

func TestImmediate(t *testing.T) {
	checkInstructions(t, "mov #0, 1", Instruction{
		Opcode: OpMov,
		A:      Operand{Mode: Immediate, Expression: constNum(0)},
		B:      Operand{Mode: Relative, Expression: constNum(1)},
	})
}

func TestDirect(t *testing.T) {
	checkInstructions(t, "mov $0, 1", Instruction{
		Opcode: OpMov,
		A:      Operand{Mode: Relative, Expression: constNum(0)},
		B:      Operand{Mode: Relative, Expression: constNum(1)},
	})
}

func TestIndirect(t *testing.T) {
	checkInstructions(t, "mov @0, 1", Instruction{
		Opcode: OpMov,
		A:      Operand{Mode: Indirect, Expression: constNum(0)},
		B:      Operand{Mode: Relative, Expression: constNum(1)},
	})
}

func TestDecrement(t *testing.T) {
	checkInstructions(t, "mov <0, 1", Instruction{
		Opcode: OpMov,
		A:      Operand{Mode: DecrementIndirect, Expression: constNum(0)},
		B:      Operand{Mode: Relative, Expression: constNum(1)},
	})
}

func TestLabelRef(t *testing.T) {
	checkInstructions(t, "imp mov imp, 1", Instruction{
		Label:  "imp",
		Opcode: OpMov,
		A:      Operand{Mode: Relative, Expression: label("imp")},
		B:      Operand{Mode: Relative, Expression: constNum(1)},
	})
}

func TestExpr(t *testing.T) {
	checkInstructions(t, "imp mov imp, imp + 1", Instruction{
		Label:  "imp",
		Opcode: OpMov,
		A:      Operand{Mode: Relative, Expression: label("imp")},
		B:      Operand{Mode: Relative, Expression: &Expression{Operation: Add, Left: label("imp"), Right: constNum(1)}},
	})
}

func TestDjn(t *testing.T) {
	checkInstructions(t, "djn 1, 2", Instruction{
		Opcode: OpDjn,
		A:      Operand{Mode: Relative, Expression: constNum(1)},
		B:      Operand{Mode: Relative, Expression: constNum(2)},
	})
}

func TestSpl(t *testing.T) {
	checkInstructions(t, "spl 2", Instruction{
		Opcode: OpSpl,
		A:      Operand{Mode: Relative, Expression: constNum(2)},
	})
}

func TestDat(t *testing.T) {
	checkInstructions(t, "dat #2", Instruction{
		Opcode: OpDat,
		B:      Operand{Mode: Immediate, Expression: constNum(2)},
	})
}

func TestEnd(t *testing.T) {
	checkInstructions(t, "spl 2\nend blah\ngargle gargle gargle", Instruction{
		Opcode: OpSpl,
		A:      Operand{Mode: Relative, Expression: constNum(2)},
	}, Instruction{
		Opcode: OpEnd,
		A:      Operand{Mode: Relative, Expression: label("blah")},
	})
}

func checkInstructions(t *testing.T, text string, instructions ...Instruction) {
	lines, _, err := ParseString(text, "string")
	if err != nil {
		t.Errorf("Error prsing '%s': %s", text, err)
		return
	}
	if diff := deep.Equal(lines, instructions); diff != nil {
		t.Errorf("Different parse of '%s' %+v", text, diff)
		return
	}
}

func checkDirectives(t *testing.T, text string, values Directives) {
	_, directives, err := ParseString(text, "string")
	if err != nil {
		t.Errorf("Error prsing '%s': %s", text, err)
		return
	}
	if diff := deep.Equal(directives, values); diff != nil {
		t.Errorf("Different parse of '%s' %+v", text, diff)
		return
	}
}

func constNum(n int) *Expression {
	return &Expression{Operation: Number, Number: n}
}

func label(l string) *Expression {
	return &Expression{Operation: Label, Label: l}
}
