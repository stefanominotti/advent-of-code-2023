package main

import (
	"regexp"
	"strconv"
)

const partBRegexPattern = `one|two|three|four|five|six|seven|eight|nine|\d`
const invertedPartBRegexPattern = `eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|\d`

func processPartBLine(line string) (int, error) {
	re := regexp.MustCompile(partBRegexPattern)
	match := re.FindString(line)
	firstNumber, err := stringNumberToInt(match)
	if err != nil {
		return 0, err
	}

	invertedRe := regexp.MustCompile(invertedPartBRegexPattern)
	match = invertedRe.FindString(reverse(line))
	lastNumber, err := stringNumberToInt(reverse(match))
	if err != nil {
		return 0, err
	}

	return firstNumber*10 + lastNumber, nil
}

func stringNumberToInt(number string) (int, error) {
	switch number {
	case "one":
		return 1, nil
	case "two":
		return 2, nil
	case "three":
		return 3, nil
	case "four":
		return 4, nil
	case "five":
		return 5, nil
	case "six":
		return 6, nil
	case "seven":
		return 7, nil
	case "eight":
		return 8, nil
	case "nine":
		return 9, nil
	default:
		return strconv.Atoi(number)
	}
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
