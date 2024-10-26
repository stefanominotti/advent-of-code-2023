package main

import (
	"advent-of-code/utils"
	"bufio"
	"strconv"
	"strings"
)

func processPartBFile(fileScanner *bufio.Scanner) (int, error) {
	time, distance, err := parsePartBFile(fileScanner)
	if err != nil {
		return 0, err
	}

	minTime, maxTime := calculateMinAndMaxButtonHoldingTime(distance, time)
	return maxTime - minTime + 1, nil
}

func parsePartBFile(fileScanner *bufio.Scanner) (int, int, error) {
	time := -1
	distance := -1
	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = strings.ReplaceAll(line, "Time:", "")
		line = strings.ReplaceAll(line, "Distance:", "")

		splitLine := utils.TrimSpacesAndSplit(line, " ")
		concatenatedLine := ""
		for _, numberString := range splitLine {
			concatenatedLine += strings.Trim(numberString, " ")
		}
		number, err := strconv.Atoi(concatenatedLine)
		if err != nil {
			return 0, 0, err
		}

		if time == -1 {
			time = number
		} else {
			distance = number
		}
	}
	return time, distance, nil
}
