package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/mbarbon/koth/arena"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	rounds   = kingpin.Flag("rounds", "Number of rounds").Short('r').Default("1").Int()
	coreSize = kingpin.Flag("core-size", "Core size").Short('s').Default("8000").Int()
	cycles   = kingpin.Flag("max-cycles", "Cycles until tie").Short('c').Default("80000").Int()
	programs = kingpin.Arg("program", "Redcode program(s)").Required().Strings()
)

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
	match := make([]int, len(*programs))

	for index, programName := range *programs {
		err := cwArena.LoadRedcodeFile(programName)
		if err != nil {
			return err
		}
		match[index] = index
	}

	cwArena.RunMatch(match)

	// print results
	programCount := cwArena.ProgramCount()
	for i, max := 0, programCount; i < max; i++ {
		program := cwArena.Program(i)
		score := cwArena.Score(i)
		fmt.Printf("%s by %s scores %d\n", program.Name(), program.Author(), score.Score())
		if programCount > 2 {
			fmt.Print("  Results:")
			for survivors := 1; survivors <= programCount; survivors++ {
				fmt.Printf(" %d", score.Wins(survivors))
			}
			fmt.Printf(" %d\n", score.Losses())
		}
	}
	if programCount == 2 {
		score1, score2 := cwArena.Score(0), cwArena.Score(1)
		fmt.Printf("Results: %d %d %d\n", score1.Wins(1), score2.Wins(1), score2.Wins(2))
	}

	return nil
}
