package main

import (
	"bufio"
)

func processPartBFile(fileScanner *bufio.Scanner) (int, error) {
	seeds, steps, err := parseFile(fileScanner)
	if err != nil {
		return 0, err
	}

	// Starting from location 0, calculate the corresponding original
	// seed of each location until a valid seed is found
	result := -1 
	currentLocation := 0
	for result == -1 {
		seed := calculateSeedFromLocation(currentLocation, steps)
		for i := 0; i < len(seeds); i+=2 {
			if (seed >= seeds[i] && seed < seeds[i] + seeds[i+1]) {
				result = currentLocation
				break
			}
		}
		currentLocation += 1
	}
	return result, nil
}

// Calculate the seed from a location by following the steps in reverse order
func calculateSeedFromLocation(location int, stepsMap map[string][][]int) int {
	seed := location
	for _, stepName := range stepsReversed() {
		step := stepsMap[stepName]
		for _, possibleConversion := range step {
			sourceDelta := seed - possibleConversion[0]
			if (sourceDelta >= 0 && sourceDelta < possibleConversion[2]) {
				seed = possibleConversion[1] + sourceDelta
				break
			}
		}
	}
	return seed
}
