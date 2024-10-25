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

func processFileFunction(lineProcessingFunction func(prevLine string, line string, nextLine string) (int, error)) func(fileScanner *bufio.Scanner) (int, error) {
	return func(fileScanner *bufio.Scanner) (int, error) {
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
			lineResult, err := lineProcessingFunction(prevLine, line, nextLine)
			if err != nil {
				return 0, err
			}
			result += lineResult
			prevLine = line
			line = nextLine
		}

		return result, nil
	}
}
