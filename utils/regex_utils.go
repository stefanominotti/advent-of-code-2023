package utils

import "regexp"

const digitRegexPattern = `\d`

func DigitRegex() (*regexp.Regexp) {
	return regexp.MustCompile(digitRegexPattern)
}

const numberRegexPattern = `\d+`

func NumberRegex() (*regexp.Regexp) {
	return regexp.MustCompile(numberRegexPattern)
}