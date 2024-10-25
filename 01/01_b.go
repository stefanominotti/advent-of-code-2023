package main

import (
	"regexp"
	"strconv"
)

const extendedDigitPattern = `one|two|three|four|five|six|seven|eight|nine|\d`
const invertedExtendedDigitPattern = `eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|\d`

func processPartBLine(line string) (int, error) {
	// Find the first digit in the line also in text format
	re := regexp.MustCompile(extendedDigitPattern)
	match := re.FindString(line)
	firstNumber, err := stringNumberToInt(match)
	if err != nil {
		return 0, err
	}

	// Find the last digit in the line also in text format
	// by matching text numbers reversed to avoid cases like
	// 'twone' in which the non-reversed regex would match 'two'
	// instead of 'one' as number
	invertedRe := regexp.MustCompile(invertedExtendedDigitPattern)
	match = invertedRe.FindString(reverse(line))
	lastNumber, err := stringNumberToInt(reverse(match))
	if err != nil {
		return 0, err
	}

	// Concat the digits
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
