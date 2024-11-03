package main

import (
	"advent-of-code/utils"
	"bufio"
	"slices"
	"strings"
)

const seedToSoil = "seed-to-soil"
const soilToFertilizer = "soil-to-fertilizer"
const fertilizerToWater = "fertilizer-to-water"
const waterToLight = "water-to-light"
const lightToTemperature = "light-to-temperature"
const temperatureToHumidity = "temperature-to-humidity"
const humidityToLocation = "humidity-to-location"

func steps() []string {
	return []string{seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation}
}

func stepsReversed() []string {
	steps := steps()
	slices.Reverse(steps)
	return steps
}

func parseFile(fileScanner *bufio.Scanner) ([]int, map[string][][]int) {
	steps := steps()

	stepsMap := map[string][][]int{}
	for _, step := range steps {
		stepsMap[step] = [][]int{}
	}

	var seeds []int
	var currentTransition string
	for fileScanner.Scan() {
		// Skip empty lined
		line := fileScanner.Text()	
		if line == "" {
			continue
		}

		// Update current transition if transition start line is matched
		line = strings.ReplaceAll(line, " map:", "")
		if slices.Contains(steps, line) {
			currentTransition = line
			continue
		}

		// Parse seed line (first file line)
		if len(seeds) == 0 {
			parsedSeeds := parseSeedsLine(line)
			seeds = parsedSeeds
			continue
		}

		// Parse step line
		stepValues, err := utils.ParseStringIntoIntArray(line, " ")
		if err != nil {
			panic(err)
		}
		stepsMap[currentTransition] = append(stepsMap[currentTransition], stepValues)
	}

	return seeds, stepsMap
}

func parseSeedsLine(line string) []int {
	seedsString := strings.Split(line, ":")[1]
	seedsInt, err := utils.ParseStringIntoIntArray(seedsString, " ")
	if err != nil {
		panic(err)
	}
	return seedsInt
}
