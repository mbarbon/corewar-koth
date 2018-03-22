package mars

func runMOV(core *Core, process *process, addrA, addrB int, modifier instructionModifier) {
	switch modifier {
	case modifierAB:
		core.cells[addrB].bField = core.cells[addrA].aField
	case modifierB:
		core.cells[addrB].bField = core.cells[addrA].bField
	case modifierI:
		core.cells[addrB] = core.cells[addrA]
	default:
		panic("Unsupported mode for instruction")
	}

	process.moveNext(core)
}

func runOP(core *Core, process *process, addrA, addrB int, modifier instructionModifier, op func(*Core, int, int) int) {
	switch modifier {
	case modifierAB:
		core.cells[addrB].bField = op(core, core.cells[addrB].bField, core.cells[addrA].aField)
	case modifierB:
		core.cells[addrB].bField = op(core, core.cells[addrB].bField, core.cells[addrA].bField)
	case modifierF:
		core.cells[addrB].aField = op(core, core.cells[addrB].aField, core.cells[addrA].aField)
		core.cells[addrB].bField = op(core, core.cells[addrB].bField, core.cells[addrA].bField)
	default:
		panic("Unsupported mode for instruction")
	}

	process.moveNext(core)
}

func runJMP(core *Core, process *process, addrA, addrB int, modifier instructionModifier) {
	process.threads[process.nextThread] = addrA
	process.moveNextThread()
}

func runSPL(core *Core, process *process, address, addrA, addrB int, modifier instructionModifier) {
	process.threads[process.nextThread] = (address + 1) % core.size
	process.threads = append(process.threads, addrA)
	process.moveNextThread()
}

func runJMZ(core *Core, process *process, addrA, addrB int, modifier instructionModifier) {
	if core.cells[addrB].bField == 0 {
		process.threads[process.nextThread] = addrA
		process.moveNextThread()
	} else {
		process.moveNext(core)
	}
}

func runDJN(core *Core, process *process, addrA, addrB int, modifier instructionModifier) {
	value := core.clampValue(core.cells[addrB].bField - 1)
	core.cells[addrB].bField = value

	if value != 0 {
		process.threads[process.nextThread] = addrA
		process.moveNextThread()
	} else {
		process.moveNext(core)
	}
}

func runSEQ(core *Core, process *process, address, addrA, addrB int, modifier instructionModifier) {
	var equal bool

	switch modifier {
	case modifierAB:
		equal = core.cells[addrB].bField == core.cells[addrA].aField
	case modifierB:
		equal = core.cells[addrB].bField == core.cells[addrA].bField
	case modifierI:
		equal = core.cells[addrB] == core.cells[addrA]
	default:
		panic("Unsupported mode for instruction")
	}

	if equal {
		process.threads[process.nextThread] = (address + 2) % core.size
		process.moveNextThread()
	} else {
		process.moveNext(core)
	}
}

func (process *process) moveNext(core *Core) {
	process.threads[process.nextThread] = (process.threads[process.nextThread] + 1) % core.size
	process.nextThread = (process.nextThread + 1) % len(process.threads)
}

func (process *process) moveNextThread() {
	process.nextThread = (process.nextThread + 1) % len(process.threads)
}

func (process *process) step(core *Core) {
	address := process.threads[process.nextThread]
	instruction := core.cells[address]
	addrA := core.address(address, instruction.aAddr, instruction.aField)
	addrB := core.address(address, instruction.bAddr, instruction.bField)
	switch instruction.opcode {
	case insnMOV:
		runMOV(core, process, addrA, addrB, instruction.modifier)
	case insnADD:
		runOP(core, process, addrA, addrB, instruction.modifier, func(core *Core, a, b int) int { return core.clampValue(a + b) })
	case insnSUB:
		runOP(core, process, addrA, addrB, instruction.modifier, func(core *Core, a, b int) int { return core.clampValue(a - b) })
	case insnJMP:
		runJMP(core, process, addrA, addrB, instruction.modifier)
	case insnJMZ:
		runJMZ(core, process, addrA, addrB, instruction.modifier)
	case insnDJN:
		runDJN(core, process, addrA, addrB, instruction.modifier)
	case insnCMP:
		runSEQ(core, process, address, addrA, addrB, instruction.modifier)
	case insnSPL:
		runSPL(core, process, address, addrA, addrB, instruction.modifier)
	default:
		process.removeThread(process.nextThread)
	}
}

func (process *process) removeThread(thread int) {
	threadCount := len(process.threads) - 1
	copy(process.threads[thread:], process.threads[thread+1:])
	process.threads = process.threads[0:threadCount]
	if process.nextThread >= thread && threadCount > 0 {
		process.nextThread = (process.nextThread + threadCount - 1) % threadCount
	}
}
