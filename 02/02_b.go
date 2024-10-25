package main

func processPartBLine(line string) (int, error) {
	game, err := createGame(line)

	if err != nil {
		return 0, err
	}

	// For each ball color store the minimim number of balls
	// with that color extracted in any extraction
	minColorNumbers := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, ball := range game.balls {
		if minColorNumbers[ball.color] < ball.number {
			minColorNumbers[ball.color] = ball.number
		}
	}

	// Returns the product of the minimums
	return minColorNumbers["red"] * minColorNumbers["green"] * minColorNumbers["blue"], nil
}
