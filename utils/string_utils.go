package utils

import "strings"

func TrimSpacesAndSplit(stringToSplit string, separator string) []string {
	return strings.Split(strings.TrimSpace(stringToSplit), separator)
}