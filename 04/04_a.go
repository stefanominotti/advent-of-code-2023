package main

import (
	"math"
)

func processPartALine(line string) (int, error) {
	winningNumbersCount, err := getLineWinningNumbersCount(line)
	if err != nil {
		return 0, err
	}

	// Return 2^(number of winning numbers)
	return int(math.Pow(2, float64(winningNumbersCount-1))), nil
}
