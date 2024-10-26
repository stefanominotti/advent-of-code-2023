package main

import (
	"bufio"
)

func processPartAFile(fileScanner *bufio.Scanner) (int, error) {
	seeds, steps, err := parseFile(fileScanner)
	if err != nil {
		return 0, err
	}

	// For each seed calculate the location and get the minimum
	result := -1 
	for _, seed := range seeds {
		seedValue := calculateLocationFromSeed(seed, steps)
		if result == -1 || seedValue < result {
			result = seedValue
		}
	}
	return result, nil
}

// Calculate the location from a seed by following the steps
func calculateLocationFromSeed(seed int, stepsMap map[string][][]int) int {
	for _, stepName := range steps() {
		step := stepsMap[stepName]
		for _, possibleConversion := range step {
			sourceDelta := seed - possibleConversion[1]
			if (sourceDelta >= 0 && sourceDelta < possibleConversion[2]) {
				seed = possibleConversion[0] + sourceDelta
				break
			}
		}
	}
	return seed
}
