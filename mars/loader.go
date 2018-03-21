package mars

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/mbarbon/koth/redcode"
)

var opcodeMap = map[redcode.Opcode]opcode{
	redcode.OpDat: insnDAT,
	redcode.OpMov: insnMOV,
	redcode.OpAdd: insnADD,
	redcode.OpSub: insnSUB,
	redcode.OpJmp: insnJMP,
	redcode.OpJmz: insnJMZ,
	redcode.OpDjn: insnDJN,
	redcode.OpCmp: insnCMP,
	redcode.OpSpl: insnSPL,
}

var modifierMap = map[redcode.Modifier]instructionModifier{}

var addressModeMap = map[redcode.AddressMode]addressMode{
	redcode.Immediate: addrIMMEDIATE,
	redcode.Relative:  addrRELATIVE,
	redcode.Indirect:  addrINDIRECT,
}

// LoadRedcodeFile parses and assembles a Redcode program
func LoadRedcodeFile(filePath string) (*Redcode, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	instructions, directives, err := redcode.ParseBytes(bytes, filePath)
	if err != nil {
		return nil, err
	}
	assembled, start, err := loadRedcode(instructions)
	if err != nil {
		return nil, err
	}

	filename := path.Base(filePath)
	program := &Redcode{
		filename:     filename,
		name:         directives["name"],
		author:       directives["author"],
		start:        start,
		instructions: assembled,
	}
	return program, nil
}

func loadRedcode(instructions []redcode.Instruction) ([]location, int, error) {
	labelOffsets := make(map[string]int)

	// find label offsets
	for index, instruction := range instructions {
		if instruction.Label != "" {
			labelOffsets[instruction.Label] = index
		}
	}

	startInstruction := 0
	locations := make([]location, len(instructions))
	if lastIns := instructions[len(instructions)-1]; lastIns.Opcode == redcode.OpEnd {
		if lastIns.A.Expression != nil {
			var err error
			startInstruction, err = evaluateOperand(labelOffsets, 0, lastIns.A.Expression)
			if err != nil {
				return nil, 0, err
			}
		}
		// remove END instruction
		instructions = instructions[0 : len(instructions)-1]
		locations = locations[0 : len(locations)-1]
	}

	for index, instruction := range instructions {
		var mappedOpcode opcode
		var mappedModifier instructionModifier

		// map opcode
		mappedOpcode, ok := opcodeMap[instruction.Opcode]
		if !ok {
			return nil, 0, fmt.Errorf("Invalid opcode %d while loading", instruction.Opcode)
		}
		// map or infer modifier
		if instruction.Modifier != redcode.ModUnknown {
			mappedModifier, ok = modifierMap[instruction.Modifier]
			if !ok {
				return nil, 0, errors.New("Invalid modifier")
			}
		} else {
			mappedModifier = inferModifier(mappedOpcode, instruction)
		}
		// address move
		aAddr, aField, err := computeAddress(labelOffsets, index, instruction.A)
		if err != nil {
			return nil, 0, err
		}
		bAddr, bField, err := computeAddress(labelOffsets, index, instruction.B)
		if err != nil {
			return nil, 0, err
		}

		location := location{
			opcode:   mappedOpcode,
			modifier: mappedModifier,
			aAddr:    aAddr,
			aField:   aField,
			bAddr:    bAddr,
			bField:   bField,
		}
		locations[index] = location
	}

	return locations, startInstruction, nil
}

func inferModifier(opcode opcode, instruction redcode.Instruction) instructionModifier {
	// from http://www.koth.org/info/guide.html
	switch opcode {
	case insnDAT:
		return modifierF
	case insnMOV:
		fallthrough
	case insnCMP:
		if instruction.A.Mode == redcode.Immediate {
			return modifierAB
		} else if instruction.B.Mode == redcode.Immediate {
			return modifierB
		} else {
			return modifierI
		}
	case insnADD:
		fallthrough
	case insnSUB:
		if instruction.A.Mode == redcode.Immediate {
			return modifierAB
		} else if instruction.B.Mode == redcode.Immediate {
			return modifierB
		} else {
			return modifierF
		}
	case insnJMP:
		fallthrough
	case insnDJN:
		fallthrough
	case insnSPL:
		fallthrough
	case insnJMZ:
		return modifierB
	default:
		panic("Unknown instruction")
	}
}

func computeAddress(labels map[string]int, instructionOffset int, operand redcode.Operand) (addressMode, int, error) {
	// no operand
	if operand.Expression == nil {
		return addrIMMEDIATE, 0, nil
	}

	mode, ok := addressModeMap[operand.Mode]
	if !ok {
		return 0, 0, errors.New("Invalid address mode")
	}
	value, err := evaluateOperand(labels, instructionOffset, operand.Expression)
	if err != nil {
		return 0, 0, err
	}
	return mode, value, nil
}

func evaluateOperand(labels map[string]int, instructionOffset int, expression *redcode.Expression) (int, error) {
	switch expression.Operation {
	case redcode.Number:
		return expression.Number, nil
	case redcode.Label:
		offset, ok := labels[expression.Label]
		if !ok {
			return 0, fmt.Errorf("Invalid label '%s'", expression.Label)
		}
		return offset - instructionOffset, nil
	case redcode.Add:
		lVal, err := evaluateOperand(labels, instructionOffset, expression.Left)
		if err != nil {
			return 0, err
		}
		rVal, err := evaluateOperand(labels, instructionOffset, expression.Right)
		if err != nil {
			return 0, err
		}
		return lVal + rVal, nil
	case redcode.Sub:
		lVal, err := evaluateOperand(labels, instructionOffset, expression.Left)
		if err != nil {
			return 0, err
		}
		rVal, err := evaluateOperand(labels, instructionOffset, expression.Right)
		if err != nil {
			return 0, err
		}
		return lVal - rVal, nil
	default:
		return 0, errors.New("Invalid expression operation")
	}
}
