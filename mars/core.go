package mars

import (
	"math/rand"
	"sort"
)

type (
	opcode              int8
	addressMode         int8
	instructionModifier int8
)

const (
	insnDAT opcode = iota
	insnMOV
	insnADD
	insnSUB
	insnJMP
	insnJMZ
	insnDJZ
	insnCMP
	insnSPL
)

const (
	addrIMMEDIATE addressMode = iota
	addrRELATIVE
	addrINDIRECT
)

const (
	modifierAB instructionModifier = iota
	modifierB
	modifierI
	modifierF
)

type location struct {
	opcode         opcode
	modifier       instructionModifier
	aAddr, bAddr   addressMode
	aField, bField int
}

type process struct {
	redcode    *Redcode
	nextThread int
	threads    []int
}

// Core is a Redcode execution core
type Core struct {
	size        int
	minInterval int
	cells       []location
	processes   []*process
	winnerIndex int
	running     bool
}

// Redcode is a ready-to-run Redcode program
type Redcode struct {
	name         string
	start        int
	instructions []location
}

func (location *location) clampValues(coreSize int) {
	location.aField = clampValue(coreSize, location.aField)
	location.bField = clampValue(coreSize, location.bField)
}

func (program *Redcode) Name() string {
	return program.name
}

func (program *Redcode) PrepareAddresses(coreSize int) {
	instructions := program.instructions
	for index := range instructions {
		instructions[index].clampValues(coreSize)
	}
}

// NewCore creates a new Core instance with the specified size
func NewCore(size int) *Core {
	return &Core{
		winnerIndex: -1,
		running:     true,
		size:        size,
		minInterval: 100,
		cells:       make([]location, size),
	}
}

func (core *Core) LoadPrograms(programs []*Redcode, rnd *rand.Rand) {
	programCount := len(programs)
	baseAddress := rnd.Intn(core.size)

	// Fisher-Yates shuffle
	shuffled := make([]*Redcode, programCount)
	copy(shuffled, programs)
	for i := programCount - 1; i > 0; i-- {
		j := rnd.Intn(i + 1)
		temp := shuffled[j]
		shuffled[j] = shuffled[i]
		shuffled[i] = temp
	}

	available := core.size
	available -= core.minInterval * programCount
	for _, program := range programs {
		available -= len(program.instructions)
	}

	// generate random numbers summing to available
	sequence := make([]int, programCount+1)
	for i := 0; i < programCount-1; i++ {
		sequence[i] = rnd.Intn(available)
	}
	sequence[programCount-1] = 0
	sequence[programCount] = available
	sort.Ints(sequence)

	for index, program := range shuffled {
		interval := sequence[index+1] - sequence[index]
		copy(core.cells[baseAddress:], program.instructions)
		if baseAddress+len(program.instructions) > core.size {
			copy(core.cells, program.instructions[core.size-baseAddress:])
		}
		core.processes = append(core.processes, &process{
			redcode:    program,
			nextThread: 0,
			threads:    []int{baseAddress + program.start},
		})
		baseAddress = (baseAddress + core.minInterval + len(program.instructions) + interval) % core.size
	}
}

func (core *Core) Step() {
	runningCount := 0
	winnerIndex := 0
	for index, process := range core.processes {
		if len(process.threads) == 0 {
			continue
		}
		process.step(core)
		if len(process.threads) != 0 {
			runningCount++
			winnerIndex = index
		}
	}
	if runningCount == 1 {
		core.winnerIndex = winnerIndex
	}
	core.running = runningCount > 1
}

func (core *Core) IsComplete() bool {
	return !core.running
}

func (core *Core) Winner() *Redcode {
	if core.winnerIndex == -1 {
		return nil
	}
	return core.processes[core.winnerIndex].redcode
}

func (core *Core) Run(maxCycles int) *Redcode {
	if !core.running {
		panic("Can't call Run() twice")
	}

	for i := 0; i < maxCycles && core.running; i++ {
		core.Step()
	}

	return core.Winner()
}

func (core *Core) address(base int, mode addressMode, field int) int {
	if mode == addrRELATIVE {
		address := core.clampValue(base + field)
		return address
	} else if mode == addrINDIRECT {
		pointer := core.clampValue(base + field)
		address := core.clampValue(pointer + core.cells[pointer].bField)
		return address
	} else if mode == addrIMMEDIATE {
		return base
	}
	panic("Illegal address mode")
}

func (core *Core) clampValue(address int) int {
	return (address%core.size + core.size) % core.size
}

func clampValue(coreSize, address int) int {
	return (address%coreSize + coreSize) % coreSize
}
