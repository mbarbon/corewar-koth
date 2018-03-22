package arena

// RunFullTournament runs all possible 1-against-1 match combinations
func RunFullTournament(arena *Arena) {
	programCount := arena.ProgramCount()
	match := make([]int, 2)

	for i := 0; i < programCount-1; i++ {
		for j := i + 1; j < programCount; j++ {
			match[0], match[1] = i, j
			arena.RunMatch(match)
		}
	}
}
