package main

import (
	"bufio"
)

func cardPowerMapPartA() map[string]int {
	return map[string]int{
		"2": 0,
		"3": 1,
		"4": 2,
		"5": 3,
		"6": 4,
		"7": 5,
		"8": 6,
		"9": 7,
		"T": 8,
		"J": 9,
		"Q": 10,
		"K": 11,
		"A": 12,
	}
}

func processPartAFile(fileScanner *bufio.Scanner) int {
	return processFile(fileScanner, parseLinePartA)
}

func parseLinePartA(line string) Hand {
	return parseLine(line, cardPowerMapPartA(), calculateHandStrengthPartA)
}

func calculateHandStrengthPartA(handCardsCount []int) int {
	return calculateHandStrength(handCardsCount, 0)
}
