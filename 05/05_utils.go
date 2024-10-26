package main

import (
	"advent-of-code/utils"
	"bufio"
	"slices"
	"strconv"
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

func parseFile(fileScanner *bufio.Scanner) ([]int, map[string][][]int, error) {
	steps := steps()

	stepsMap := map[string][][]int{}
	for _, step := range steps {
		stepsMap[step] = [][]int{}
	}

	var seeds []int
	var currentTransition string
	for fileScanner.Scan() {
		line := fileScanner.Text()	
		if line == "" {
			continue
		}

		line = strings.ReplaceAll(line, " map:", "")
		if slices.Contains(steps, line) {
			currentTransition = line
			continue
		}

		if len(seeds) == 0 {
			parsedSeeds, err := parseSeedsLine(line)
			if err != nil {
				return nil, nil, err
			}
			seeds = parsedSeeds
			continue
		}

		stepValues, err := parseStepLine(line)
		if err != nil {
			return nil, nil, err
		}
		stepsMap[currentTransition] = append(stepsMap[currentTransition], stepValues)
	}

	return seeds, stepsMap, nil
}

func parseSeedsLine(line string) ([]int, error) {
	var seeds []int
	seedsString := strings.Split(line, ":")[1]
	seedsStringValues := utils.TrimSpacesAndSplit(seedsString, " ")
	for _, seedString := range seedsStringValues {
		seed, err := strconv.Atoi(seedString)
		if err != nil {
			return nil, err
		}
		seeds = append(seeds, seed)
	}
	return seeds, nil
}

func parseStepLine(line string) ([]int, error) {
	stepStringValues := utils.TrimSpacesAndSplit(line, " ")
	var stepValues []int
	for _, valueString := range stepStringValues {
		value, err := strconv.Atoi(valueString)
		if err != nil {
			return nil, err
		}
		stepValues = append(stepValues, value)
	}
	return stepValues, nil
}
