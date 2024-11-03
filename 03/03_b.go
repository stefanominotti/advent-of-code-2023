package main

import (
	"advent-of-code/utils"
	"regexp"
	"strconv"
)

const starRegexPattern = `\*`

func processPartBLine(prevLine string, line string, nextLine string) int {
	// Find the indexes of all '*' in the line
	re := regexp.MustCompile(starRegexPattern)
	matches := re.FindAllIndex([]byte(line), -1)

	// For each '*' checks get all the adjacent numbers from the current, previous and
	// next line. If there are exactly two numbers sum their product to the
	// current line result.
	result := 0
	linesArray := []string{prevLine, line, nextLine}
	for _, numberIndexes := range matches {
		var adjacentNumbers []int
		for _, checkLine := range linesArray {
			lineNumbers := getAdjacentNumbers(numberIndexes[0], checkLine)
			adjacentNumbers = append(adjacentNumbers, lineNumbers...)
		} 
		if (len(adjacentNumbers) == 2) {
			result += adjacentNumbers[0] * adjacentNumbers[1]
		}
	}

	return result
}

// Given an index and a line returns all the numbers in the line
// which are adjacent to the given index.
func getAdjacentNumbers(idx int, line string) []int {
	re := utils.NumberRegex()
	matches := re.FindAllIndex([]byte(line), -1)

	var numbers []int
	for _, match := range matches {
		if (idx >= match[0]-1 && idx < match[1]+1) {
			number, err := strconv.Atoi(string(line[match[0]:match[1]]))
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, number)
		}
	}

	return numbers
}
