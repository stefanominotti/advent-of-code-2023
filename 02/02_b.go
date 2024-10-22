package main

func processPartBLine(line string) (int, error) {
	game, err := createGame(line)

	minColorNumbers := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	if err != nil {
		return 0, err
	}

	for _, ball := range game.balls {
		if minColorNumbers[ball.color] < ball.number {
			minColorNumbers[ball.color] = ball.number
		}
	}

	return minColorNumbers["red"] * minColorNumbers["green"] * minColorNumbers["blue"], nil
}
