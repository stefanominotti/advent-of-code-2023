package main

import (
	"regexp"
	"strconv"
	"strings"
)

const gameRegexPattern = `Game (\d+):.*`

type BallExtraction struct {
    color string
	number int
}

type Game struct {
	number int
	balls []BallExtraction
}

func createGame(line string) (Game, error) {
	re := regexp.MustCompile(gameRegexPattern)
	gameNumberString := re.FindStringSubmatch(line)[1]

	extractions := trimSpacesAndSplit(strings.Split(line, ":")[1], ";")

	var balls []BallExtraction

	for _, extraction := range extractions {
		extractedBalls := trimSpacesAndSplit(extraction, ",")

		for _, extractedBall := range extractedBalls {
			splitBall := trimSpacesAndSplit(extractedBall, " ")

			color := splitBall[1]
			number, err := strconv.Atoi(splitBall[0])
			if err != nil {
				return Game{}, err
			}
			
			balls = append(balls, BallExtraction{color, number})
		}
	}

	gameNumber, err := strconv.Atoi(gameNumberString)
	if err != nil {
		return Game{}, err
	}
	return Game{gameNumber, balls}, nil
}

func trimSpacesAndSplit(stringToSplit string, separator string) []string {
	return strings.Split(strings.TrimSpace(stringToSplit), separator)
}