package main

import (
	"regexp"
	"strconv"
)

const extendedDigitPattern = `one|two|three|four|five|six|seven|eight|nine|\d`
const invertedExtendedDigitPattern = `eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|\d`

func processPartBLine(line string) int {
	// Find the first digit in the line also in text format
	re := regexp.MustCompile(extendedDigitPattern)
	match := re.FindString(line)
	firstNumber := stringNumberToInt(match)

	// Find the last digit in the line also in text format
	// by matching text numbers reversed to avoid cases like
	// 'twone' in which the non-reversed regex would match 'two'
	// instead of 'one' as number
	invertedRe := regexp.MustCompile(invertedExtendedDigitPattern)
	match = invertedRe.FindString(reverse(line))
	lastNumber := stringNumberToInt(reverse(match))

	// Concat the digits
	return firstNumber*10 + lastNumber
}

func stringNumberToInt(number string) int {
	switch number {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		result, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		return result
	}
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
