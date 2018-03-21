package mars

func runMOV(core *Core, process *process, address int, instruction location) {
	addrA := core.address(address, instruction.aAddr, instruction.aField)
	addrB := core.address(address, instruction.bAddr, instruction.bField)

	switch instruction.modifier {
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

func runOP(core *Core, process *process, address int, instruction location, op func(*Core, int, int) int) {
	addrA := core.address(address, instruction.aAddr, instruction.aField)
	addrB := core.address(address, instruction.bAddr, instruction.bField)

	switch instruction.modifier {
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

func runJMP(core *Core, process *process, address int, instruction location) {
	addrA := core.address(address, instruction.aAddr, instruction.aField)
	process.threads[process.nextThread] = addrA
	process.moveNextThread()
}

func runSPL(core *Core, process *process, address int, instruction location) {
	addrB := core.address(address, instruction.bAddr, instruction.bField)
	process.threads[process.nextThread] = (address + 1) % core.size
	// valid for '86 rules
	process.threads = append(process.threads, 0)
	copy(process.threads[process.nextThread+1:], process.threads[process.nextThread:])
	process.threads[process.nextThread] = addrB
}

func runJMZ(core *Core, process *process, address int, instruction location) {
	addrA := core.address(address, instruction.aAddr, instruction.aField)
	addrB := core.address(address, instruction.bAddr, instruction.bField)

	if core.cells[addrB].bField == 0 {
		process.threads[process.nextThread] = addrA
		process.moveNextThread()
	} else {
		process.moveNext(core)
	}
}

func runDJN(core *Core, process *process, address int, instruction location) {
	addrA := core.address(address, instruction.aAddr, instruction.aField)
	addrB := core.address(address, instruction.bAddr, instruction.bField)

	value := core.clampValue(core.cells[addrB].bField - 1)
	core.cells[addrB].bField = value

	if value != 0 {
		process.threads[process.nextThread] = addrA
		process.moveNextThread()
	} else {
		process.moveNext(core)
	}
}

func runSEQ(core *Core, process *process, address int, instruction location) {
	addrA := core.address(address, instruction.aAddr, instruction.aField)
	addrB := core.address(address, instruction.bAddr, instruction.bField)

	var equal bool

	switch instruction.modifier {
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
	switch instruction.opcode {
	case insnMOV:
		runMOV(core, process, address, instruction)
	case insnADD:
		runOP(core, process, address, instruction, func(core *Core, a, b int) int { return core.clampValue(a + b) })
	case insnSUB:
		runOP(core, process, address, instruction, func(core *Core, a, b int) int { return core.clampValue(a - b) })
	case insnJMP:
		runJMP(core, process, address, instruction)
	case insnJMZ:
		runJMZ(core, process, address, instruction)
	case insnDJN:
		runDJN(core, process, address, instruction)
	case insnCMP:
		runSEQ(core, process, address, instruction)
	case insnSPL:
		runSPL(core, process, address, instruction)
	default:
		process.removeThread(process.nextThread)
	}
}

func (process *process) removeThread(thread int) {
	copy(process.threads[thread+1:], process.threads[thread:])
	process.threads = process.threads[0 : len(process.threads)-1]
	if process.nextThread > thread {
		process.nextThread--
	}
}
