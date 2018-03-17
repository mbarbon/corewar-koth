package mars

import (
	"testing"
)

const baseAddr = 2
const coreSize = 10

var NOP = location{
	opcode:   insnJMP,
	modifier: modifierB,
	aAddr:    addrRELATIVE,
	aField:   1,
}

func TestTermination(t *testing.T) {
	core := loadCore(coreSize, location{opcode: insnDAT, modifier: modifierF})
	core.Step()
	if len(core.processes[0].threads) != 0 {
		t.Error("Program did not terminate")
	}
}

func TestMovSanity(t *testing.T) {
	core := loadCore(coreSize, location{
		opcode:   insnMOV,
		modifier: modifierI,
		aAddr:    addrRELATIVE,
		aField:   0,
		bAddr:    addrRELATIVE,
		bField:   1,
	})
	core.Step()
	checkProgramState(t, core, baseAddr+1)
	if core.cells[baseAddr] != core.cells[baseAddr+1] {
		t.Error("The copy did not succeed")
	}
}

func TestAddSanity(t *testing.T) {
	core := loadCore(coreSize, location{
		opcode:   insnADD,
		modifier: modifierAB,
		aAddr:    addrIMMEDIATE,
		aField:   7,
		bAddr:    addrRELATIVE,
		bField:   1,
	}, location{
		opcode: insnDAT,
		aAddr:  addrIMMEDIATE,
		aField: 4,
		bAddr:  addrIMMEDIATE,
		bField: 9,
	})
	core.Step()
	checkProgramState(t, core, baseAddr+1)
	if core.cells[baseAddr+1].aField != 4 {
		t.Error("A-field incorrectly modified")
	}
	if core.cells[baseAddr+1].bField != 16%coreSize {
		t.Error("B-field correctly modified")
	}
}

func TestSubSanity(t *testing.T) {
	core := loadCore(coreSize, location{
		opcode:   insnSUB,
		modifier: modifierAB,
		aAddr:    addrIMMEDIATE,
		aField:   77,
		bAddr:    addrRELATIVE,
		bField:   1,
	}, location{
		opcode: insnDAT,
		aAddr:  addrIMMEDIATE,
		aField: 4,
		bAddr:  addrIMMEDIATE,
		bField: 9,
	})
	core.Step()
	checkProgramState(t, core, baseAddr+1)
	if core.cells[baseAddr+1].aField != 4 {
		t.Error("A-field incorrectly modified")
	}
	if core.cells[baseAddr+1].bField != 2 {
		t.Error("B-field correctly modified")
	}
}

func TestJmpSanity(t *testing.T) {
	core := loadCore(coreSize, location{
		opcode:   insnJMP,
		modifier: modifierB,
		aAddr:    addrRELATIVE,
		aField:   -1,
	})
	core.Step()
	checkProgramState(t, core, baseAddr-1)
}

func TestSplSanity(t *testing.T) {
	core := loadCore(coreSize, location{
		opcode:   insnSPL,
		modifier: modifierB,
		bAddr:    addrRELATIVE,
		bField:   -1,
	})
	core.Step()
	checkProgramState(t, core, baseAddr-1, baseAddr+1)
}

func TestThreads(t *testing.T) {
	core := loadCore(coreSize, location{
		opcode:   insnSPL,
		modifier: modifierB,
		bAddr:    addrRELATIVE,
		bField:   2,
	},
		NOP,
		location{
			opcode:   insnJMP,
			modifier: modifierB,
			aAddr:    addrRELATIVE,
			aField:   -1,
		})
	core.Step()
	checkProgramState(t, core, baseAddr+2, baseAddr+1)
	core.Step()
	checkProgramState(t, core, baseAddr+1, baseAddr+1)
	core.Step()
	checkProgramState(t, core, baseAddr+1, baseAddr+2)
	core.Step()
	checkProgramState(t, core, baseAddr+2, baseAddr+2)
	core.Step()
	checkProgramState(t, core, baseAddr+2, baseAddr+1)
}

func TestIPWrap(t *testing.T) {
	size := 3
	cells := []location{NOP, NOP, location{
		opcode:   insnADD,
		modifier: modifierAB,
		aAddr:    addrIMMEDIATE,
		aField:   1,
		bAddr:    addrRELATIVE,
		bField:   0,
	}}
	core := &Core{
		running: true,
		size:    size,
		cells:   cells,
		processes: []*process{
			&process{
				nextThread: 0,
				threads:    []int{2},
			},
		},
	}

	if core.cells[2].bField != 0 {
		t.Error("Failed precondition")
	}

	core.Step()
	checkProgramState(t, core, 0)
	if core.cells[2].bField != 1 {
		t.Error("Add did not execute correctly")
	}

	core.Step()
	checkProgramState(t, core, 1)
}

func loadCore(size int, instructions ...location) *Core {
	core := &Core{
		running: true,
		size:    size,
		cells:   make([]location, size),
		processes: []*process{
			&process{
				nextThread: 0,
				threads:    []int{baseAddr},
			},
		},
	}
	copy(core.cells[baseAddr:], instructions)

	return core
}

func checkProgramState(t *testing.T, core *Core, addresses ...int) {
	if count := len(core.processes[0].threads); count != len(addresses) {
		t.Errorf("Incorrect thread count %d", count)
		return
	}
	for index, address := range addresses {
		if core.processes[0].threads[index] != address {
			t.Errorf("Incorrect program counters: %v", core.processes[0].threads)
			return
		}
	}
}
