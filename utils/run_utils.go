package utils

import (
	"bufio"
	"fmt"
	"os"
)

func RunChallenge(processingFunctions ...func(fileScanner *bufio.Scanner) int) {
	for idx, function := range processingFunctions {
		result := processFile(function)
		fmt.Printf("Part %s: %d\n", string('A'+idx), result)
	}
}

func ProcessFileLineByLineFunction(processLineFunction func(line string) int) func(*bufio.Scanner) int {
	return func(fileScanner *bufio.Scanner) int {
		return processFileLineByLine(processLineFunction)
	}
}

func processFile(processingFunction func(fileScanner *bufio.Scanner) int) int {
	filePath := "input.txt"
	readFile, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	result := processingFunction(fileScanner)

	readFile.Close()

	return result
}

func processFileLineByLine(processingFunction func(line string) int) int {
	lineByLineProcessing := func(fileScanner *bufio.Scanner) int {
		result := 0
		for fileScanner.Scan() {
			line := fileScanner.Text()
			lineResult := processingFunction(line)
			result += lineResult
		}

		return result
	}
	return processFile(lineByLineProcessing)
}
