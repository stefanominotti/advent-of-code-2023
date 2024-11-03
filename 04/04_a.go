package main

import (
	"math"
)

func processPartALine(line string) int {
	winningNumbersCount := getLineWinningNumbersCount(line)

	// Return 2^(number of winning numbers)
	return int(math.Pow(2, float64(winningNumbersCount-1)))
}
