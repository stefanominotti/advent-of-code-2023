package main

import (
	"advent-of-code/utils"
	"strconv"
	"unicode"
)

func processPartALine(prevLine string, line string, nextLine string) (int, error) {
	// Find the indexes of all numbers in the line
	re := utils.NumberRegex()
	matches := re.FindAllIndex([]byte(line), -1)

	// For each number checks whether in the current, previous or next line there is
	// an adjacent symbol. If so sum the number to the line result.
	result := 0
	linesArray := []string{prevLine, line, nextLine}
	for _, numberIndexes := range matches {
		for _, checkLine := range linesArray {
			if isThereAdjacentSymbol(numberIndexes[0], numberIndexes[1], checkLine) {
				number, err := strconv.Atoi(string(line[numberIndexes[0]:numberIndexes[1]]))
				if err != nil {
					return 0, err
				}
	
				result += number
				continue
			}
		} 
	}

	return result, nil
}

// Given a range of indexes (start inclusive, end exclusive) and a line
// returns true if the line contains a symbol different from '.' adjacent
// or inside the range
func isThereAdjacentSymbol(firstIdx int, lastIdx int, line string) bool {
	for i := firstIdx-1; i < lastIdx+1; i++ {
		if i < 0 || i >= len(line) {
			continue
		}
		if !unicode.IsDigit(rune(line[i])) && !(string(line[i]) == ".") {
			return true
		}
	}
	return false
}
