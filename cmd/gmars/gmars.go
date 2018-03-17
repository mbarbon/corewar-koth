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
	rounds   = kingpin.Flag("r", "Number of rounds").Default("1").Int()
	coreSize = kingpin.Flag("s", "Core size").Default("8000").Int()
	cycles   = kingpin.Flag("c", "Cycles until tie").Default("80000").Int()
	programs = kingpin.Arg("program", "Redcode program").Required().Strings()
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

	redcodePrograms := make([]*mars.Redcode, len(*programs))
	for index, programName := range *programs {
		program, err := mars.LoadRedcodeFile(programName)
		if err != nil {
			return err
		}
		program.PrepareAddresses(*coreSize)
		redcodePrograms[index] = program
	}

	for i := 0; i < *rounds; i++ {
		core := mars.NewCore(*coreSize)
		core.LoadPrograms(redcodePrograms, rnd)
		winner := core.Run(*cycles)
		if winner == nil {
			fmt.Printf("Tie\n")
		} else {
			fmt.Printf("Winner is '%s'\n", winner.Name())
		}
	}

	return nil
}
