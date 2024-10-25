package main

import (
	"advent-of-code/utils"
	"strconv"
)

func processPartALine(line string) (int, error) {
	// Find all digits in the line
	re := utils.DigitRegex()
	matches := re.FindAllString(line, -1)

	// Get the first digit matched
	firstNumber, err := strconv.Atoi(matches[0])
	if err != nil {
		return 0, err
	}

	// Get the last digit matched
	lastNumber, err := strconv.Atoi(matches[len(matches)-1])
	if err != nil {
		return 0, err
	}

	// Concat the digits
	return firstNumber*10 + lastNumber, nil
}
