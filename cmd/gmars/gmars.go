package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/mbarbon/koth/mars"
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

type matchScore struct {
	matchResult []int
}

func (matchScore *matchScore) computeScore() int {
	score := 0
	w := len(matchScore.matchResult)
	for index, count := range matchScore.matchResult {
		s := index + 1
		f := (w*w - 1) / s
		score += count * f
	}
	return score
}

func run() error {
	rnd := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	redcodePrograms := make([]*mars.Redcode, len(*programs))
	for index, programName := range *programs {
		program, err := mars.LoadRedcodeFile(programName)
		if err != nil {
			return err
		}
		program.PrepareAddresses(*coreSize)
		redcodePrograms[index] = program
	}

	scores := make(map[*mars.Redcode]matchScore)
	for _, program := range redcodePrograms {
		scores[program] = matchScore{
			matchResult: make([]int, len(redcodePrograms)),
		}
	}

	for i := 0; i < *rounds; i++ {
		core := mars.NewCore(*coreSize)
		core.LoadPrograms(redcodePrograms, rnd)
		core.Run(*cycles)
		winnerCount := core.RunningCount()
		for _, program := range core.RunningPrograms() {
			scores[program].matchResult[winnerCount-1]++
		}
	}

	for _, program := range redcodePrograms {
		score := scores[program]
		fmt.Printf("%s by %s scores %d\n", program.Name(), program.Author(), score.computeScore())
		if len(redcodePrograms) > 2 {
			losses := *rounds
			fmt.Print("  Results:")
			for _, wins := range score.matchResult {
				fmt.Printf(" %d", wins)
				losses -= wins
			}
			fmt.Printf(" %d\n", losses)
		}
	}
	if len(redcodePrograms) == 2 {
		wins1 := scores[redcodePrograms[0]].matchResult[0]
		wins2 := scores[redcodePrograms[1]].matchResult[0]
		ties := *rounds - wins1 - wins2
		fmt.Printf("Results: %d %d %d\n", wins1, wins2, ties)
	}

	return nil
}
