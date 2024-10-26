package main

import (
	"advent-of-code/utils"
	"bufio"
	"strings"
)

func processPartAFile(fileScanner *bufio.Scanner) (int, error) {
	times, distances, err := parsePartAFile(fileScanner)
	if err != nil {
		return 0, err
	}

	result := 1
	for idx, time := range times {
		distance := distances[idx]
		minTime, maxTime := calculateMinAndMaxButtonHoldingTime(distance, time)
		result *= maxTime - minTime + 1
	}
	return result, nil
}

func parsePartAFile(fileScanner *bufio.Scanner) ([]int, []int, error) {
	var times []int
	var distances []int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = strings.ReplaceAll(line, "Time:", "")
		line = strings.ReplaceAll(line, "Distance:", "")

		parsedLine, err := utils.ParseStringIntoIntArray(line, " ")
		if err != nil {
			return nil, nil, err
		}
		if len(times) == 0 {
			times = parsedLine
		} else {
			distances = parsedLine
		}
	}
	return times, distances, nil
}
