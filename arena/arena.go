package arena

import (
	"math/rand"
	"time"

	"github.com/mbarbon/koth/mars"
)

type option func(*Arena)

// MatchScore stores the score for a program
type MatchScore struct {
	matchResult []int
	wins        int
	arena       *Arena
}

func (score *MatchScore) recordWin(survivors int) {
	if score.matchResult == nil {
		score.matchResult = make([]int, len(score.arena.programs))
	}
	score.wins++
	score.matchResult[survivors-1]++
}

// Score returns the total program score
func (score *MatchScore) Score() int {
	total := 0
	w := len(score.arena.programs)
	for index, count := range score.matchResult {
		s := index + 1
		f := (w*w - 1) / s
		total += count * f
	}
	return total
}

// Wins returns the number of victories with the given number of survivors
func (score *MatchScore) Wins(survivors int) int {
	if score.wins == 0 {
		return 0
	}
	return score.matchResult[survivors-1]
}

// Losses returns the number of losses
func (score *MatchScore) Losses() int {
	return score.arena.rounds - score.wins
}

// WithRng uses the specified random number generator for the arena
func WithRng(rng *rand.Rand) option {
	return func(arena *Arena) {
		arena.rng = rng
	}
}

// Arena holds a set of pre-processed redcode programs and their scores
type Arena struct {
	rng       *rand.Rand
	rounds    int
	coreSize  int
	maxCycles int
	programs  []*mars.Redcode
	scores    []MatchScore
}

// NewArena initializes a new empty arena
func NewArena(coreSize, maxCycles, rounds int, options ...option) *Arena {
	arena := &Arena{
		rounds:    rounds,
		coreSize:  coreSize,
		maxCycles: maxCycles,
	}
	for _, option := range options {
		option(arena)
	}
	if arena.rng == nil {
		arena.rng = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	}
	return arena
}

// LoadRedcodeFile loads a Redcode assembly file into the arena
func (arena *Arena) LoadRedcodeFile(filepath string) error {
	program, err := mars.LoadRedcodeFile(filepath)
	if err != nil {
		return err
	}
	program.PrepareAddresses(arena.coreSize)
	arena.programs = append(arena.programs, program)
	arena.scores = append(arena.scores, MatchScore{arena: arena})
	return nil
}

// RunMatch runs a single match with two warriors
func (arena *Arena) RunMatch(programIndices []int) {
	warriors := make([]*mars.Redcode, len(programIndices))
	for index, programIndex := range programIndices {
		warriors[index] = arena.programs[programIndex]
	}
	for i := 0; i < arena.rounds; i++ {
		core := mars.NewCore(arena.coreSize)
		core.LoadPrograms(warriors, arena.rng)
		core.Run(arena.maxCycles)
		survivors := core.RunningCount()
		for _, programIndex := range core.RunningProgramIndices() {
			arena.scores[programIndices[programIndex]].recordWin(survivors)
		}
	}
}

// ProgramCount returns the number of programs loaded in this arena
func (arena *Arena) ProgramCount() int {
	return len(arena.programs)
}

// Program returns the program with the specified index
func (arena *Arena) Program(index int) *mars.Redcode {
	return arena.programs[index]
}

// Score returns the score for the program with the specified index
func (arena *Arena) Score(index int) *MatchScore {
	return &arena.scores[index]
}
