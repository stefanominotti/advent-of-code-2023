package main

import (
	"advent-of-code/utils"
	"bufio"
	"strconv"
	"strings"
)

func processPartBFile(fileScanner *bufio.Scanner) int {
	time, distance := parsePartBFile(fileScanner)
	minTime, maxTime := calculateMinAndMaxButtonHoldingTime(distance, time)
	return maxTime - minTime + 1
}

func parsePartBFile(fileScanner *bufio.Scanner) (int, int) {
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
			panic(err)
		}

		if time == -1 {
			time = number
		} else {
			distance = number
		}
	}
	return time, distance
}
