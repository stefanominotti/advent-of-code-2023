package main

func processPartALine(line string) int {
	game := createGame(line)

	cubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	// For each ball extraction check wether the number of balls
	// is lower than the max number
	for _, ball := range game.balls {
		if cubes[ball.color] < ball.number {
			return 0
		}
	}

	// In case all the extractions respect the condition,
	// returns the game number
	return game.number
}
