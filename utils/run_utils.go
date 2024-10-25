package utils

import (
	"bufio"
	"fmt"
	"os"
)

func RunChallenge(processingFunctionPartA func(fileScanner *bufio.Scanner) (int, error), processingFunctionPartB func(fileScanner *bufio.Scanner) (int, error)) {
	for _, part := range []string{"A", "B"} {
		var result int
		var err error
		if (part == "A") {
			result, err = processFile(processingFunctionPartA)
		} else {
			result, err = processFile(processingFunctionPartB)
		}

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Part %s: %d\n", part, result)
	}
}

func ProcessFileLineByLineFunction(processLineFunction func(line string) (int, error)) func(*bufio.Scanner) (int, error) {
	return func(fileScanner *bufio.Scanner) (int, error) {
		return processFileLineByLine(processLineFunction)
	}
}

func processFile(processingFunction func(fileScanner *bufio.Scanner) (int, error)) (int, error) {
	filePath := "input.txt"
	readFile, err := os.Open(filePath)

	if err != nil {
		return 0, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	result, err := processingFunction(fileScanner)
	if err != nil {
		return 0, err
	}

	readFile.Close()

	return result, nil
}

func processFileLineByLine(processingFunction func(line string) (int, error)) (int, error) {
	lineByLineProcessing := func (fileScanner *bufio.Scanner) (int, error) {
		result := 0
		for fileScanner.Scan() {
			line := fileScanner.Text()
			lineResult, err := processingFunction(line)
			if err != nil {
				return 0, err
			}
			result += lineResult
		}

		return result, nil
	}
	return processFile(lineByLineProcessing)
}
