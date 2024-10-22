package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	result, err := processFile(processPartALine)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part A:", result)

	result, err = processFile(processPartBLine)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part B:", result)
}

func processFile(processingFunction func(line string) (int, error)) (int, error) {
	filePath := "input.txt"
	readFile, err := os.Open(filePath)

	if err != nil {
		return 0, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	result := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lineResult, err := processingFunction(line)
		if err != nil {
			return 0, err
		}
		result += lineResult
	}

	readFile.Close()

	return result, err
}
