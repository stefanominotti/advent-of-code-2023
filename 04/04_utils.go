package main

import (
	"advent-of-code/utils"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Scratchcard struct {
	number int
	numbers []int
	winningNumbers []int
}

// Parse a line and retrieve effective winning numbers
func getLineWinningNumbersCount(line string) int {
	scratchcard := createScratchcard(line)

	// Retrieve effective winning numbers, i.e. numbers which are in both
	// numbers and winning numbers lists
	var winningNumbersCount int
	for _, number := range scratchcard.numbers {
		if slices.Contains(scratchcard.winningNumbers, number) {
			winningNumbersCount += 1
		}
	}

	return winningNumbersCount
}

const cardRegexPattern = `Card *(\d+):.*`

// Build a Scratchcard object continaing the number of the scratchcard and 
// the list of all numbers and winning numbers.
func createScratchcard(line string) Scratchcard {
	// Retrieve scratchcard number
	re := regexp.MustCompile(cardRegexPattern)
	cardNumberString := re.FindStringSubmatch(line)[1]

	// Retrieve the strings containing numbers and winning numbers
	numbersStrings := utils.TrimSpacesAndSplit(strings.Split(line, ":")[1], "|")

	// Convert the numbers string into numbers array
	numbers, err := utils.ParseStringIntoIntArray(numbersStrings[0], " ")
	if err != nil {
		panic(err)
	}

	// Convert the winning numbers string into numbers array
	winningNumbers, err := utils.ParseStringIntoIntArray(numbersStrings[1], " ")
	if err != nil {
		panic(err)
	}

	cardNumber, err := strconv.Atoi(cardNumberString)
	if err != nil {
		panic(err)
	}
	return Scratchcard{cardNumber, numbers, winningNumbers}
}
