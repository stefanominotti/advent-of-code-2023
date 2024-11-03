package main

import (
	"bufio"
)

func cardPowerMapPartB() map[string]int {
	return map[string]int{
		"J": 0,
		"2": 1,
		"3": 2,
		"4": 3,
		"5": 4,
		"6": 5,
		"7": 6,
		"8": 7,
		"9": 8,
		"T": 9,
		"Q": 10,
		"K": 11,
		"A": 12,
	}
}

func processPartBFile(fileScanner *bufio.Scanner) (int, error) {
	return processFile(fileScanner, parseLinePartB)
}

func parseLinePartB(line string) Hand {
	return parseLine(line, cardPowerMapPartB(), calculateHandStrengthPartB)
}

// Call the calculateHandStrength function passing the
// number of J (jollies)
func calculateHandStrengthPartB(handCardsCount []int) int {
	jCount := handCardsCount[0]
	handCardsCount = handCardsCount[1:]

	return calculateHandStrength(handCardsCount, jCount)
}
