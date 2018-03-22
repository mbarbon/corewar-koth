package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/mbarbon/koth/arena"
	"github.com/mbarbon/koth/mars"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	rounds   = kingpin.Flag("rounds", "Number of rounds").Short('r').Default("1").Int()
	coreSize = kingpin.Flag("core-size", "Core size").Short('s').Default("8000").Int()
	cycles   = kingpin.Flag("max-cycles", "Cycles until tie").Short('c').Default("80000").Int()
	programs = kingpin.Arg("program", "Redcode program(s)").Required().Strings()
)

type scoredProgram struct {
	score   int
	program *mars.Redcode
}

func main() {
	kingpin.Parse()

	err := run()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	rnd := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	cwArena := arena.NewArena(*coreSize, *cycles, *rounds, arena.WithRng(rnd))

	for _, programName := range *programs {
		err := cwArena.LoadRedcodeFile(programName)
		if err != nil {
			return err
		}
	}

	arena.RunFullTournament(cwArena)

	// print results
	programCount := cwArena.ProgramCount()
	scoredPrograms := make([]scoredProgram, programCount)
	for index := 0; index < programCount; index++ {
		scoredPrograms[index] = scoredProgram{
			score:   cwArena.Score(index).Score(),
			program: cwArena.Program(index),
		}
	}

	sort.Slice(scoredPrograms, func(i, j int) bool {
		return scoredPrograms[i].score > scoredPrograms[j].score
	})

	for _, scoredProgram := range scoredPrograms {
		score, program := scoredProgram.score, scoredProgram.program
		fmt.Printf("%4d: %s by %s\n", score, program.Name(), program.Author())
	}

	return nil
}
