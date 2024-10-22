package main

import (
	"regexp"
	"strconv"
)

const partARegexPattern = `\d`

func processPartALine(line string) (int, error) {
	re := regexp.MustCompile(partARegexPattern)
	matches := re.FindAllString(line, -1)

	firstNumber, err := strconv.Atoi(matches[0])
	if err != nil {
		return 0, err
	}

	lastNumber, err := strconv.Atoi(matches[len(matches)-1])
	if err != nil {
		return 0, err
	}

	return firstNumber*10 + lastNumber, nil
}
