package utils

import (
	"strconv"
	"strings"
)

// Trim the string and split by separator char
func TrimSpacesAndSplit(stringToSplit string, separator string) []string {
	return strings.Split(strings.TrimSpace(stringToSplit), separator)
}

// Split the string of numbers by separator char and parse them into int type
func ParseStringIntoIntArray(s string, separator string) ([]int, error) {
	splitString := TrimSpacesAndSplit(s, " ")
	numbers := []int{}

	for _, numberString := range splitString {
		if numberString == "" {
			continue
		}
		number, err := strconv.Atoi(strings.Trim(numberString, " "))
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}