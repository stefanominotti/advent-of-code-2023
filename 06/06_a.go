package main

import (
	"advent-of-code/utils"
	"bufio"
	"strings"
)

func processPartAFile(fileScanner *bufio.Scanner) int {
	times, distances := parsePartAFile(fileScanner)

	result := 1
	for idx, time := range times {
		distance := distances[idx]
		minTime, maxTime := calculateMinAndMaxButtonHoldingTime(distance, time)
		result *= maxTime - minTime + 1
	}
	return result
}

func parsePartAFile(fileScanner *bufio.Scanner) ([]int, []int) {
	var times []int
	var distances []int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = strings.ReplaceAll(line, "Time:", "")
		line = strings.ReplaceAll(line, "Distance:", "")

		parsedLine, err := utils.ParseStringIntoIntArray(line, " ")
		if err != nil {
			panic(err)
		}
		if len(times) == 0 {
			times = parsedLine
		} else {
			distances = parsedLine
		}
	}
	return times, distances
}
