package main

func cubes() map[string]int {
	return map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
}

func processPartALine(line string) (int, error) {
	game, err := createGame(line)

	if err != nil {
		return 0, err
	}

	for _, ball := range game.balls {
		if cubes()[ball.color] < ball.number {
			return 0, nil
		}
	}

	return game.number, nil
}
