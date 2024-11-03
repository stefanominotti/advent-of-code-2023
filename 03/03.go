package main

import (
	"advent-of-code/utils"
	"bufio"
)

func main() {
	processingFunctionPartA := processFileFunction(processPartALine)
	processingFunctionPartB := processFileFunction(processPartBLine)
	utils.RunChallenge(processingFunctionPartA, processingFunctionPartB)
}

func processFileFunction(lineProcessingFunction func(prevLine string, line string, nextLine string) int) func(fileScanner *bufio.Scanner) int {
	return func(fileScanner *bufio.Scanner) int {
		var prevLine string
		var line string
		var nextLine string

		result := 0
		for fileScanner.Scan() || line != "" {
			if line == "" {
				line = fileScanner.Text()
				continue
			}
			nextLine = fileScanner.Text()
			result += lineProcessingFunction(prevLine, line, nextLine)
			prevLine = line
			line = nextLine
		}

		return result
	}
}
