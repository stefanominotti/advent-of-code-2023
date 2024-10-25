package main

import (
	"regexp"
	"strconv"
	"strings"
)

type BallExtraction struct {
    color string
	number int
}

type Game struct {
	number int
	balls []BallExtraction
}

const gameRegexPattern = `Game (\d+):.*`

// Build a Game object continaing the number of the game and 
// the list of all extracted balls (BallExtraction) in the game.
func createGame(line string) (Game, error) {
	// Retrieve game number
	re := regexp.MustCompile(gameRegexPattern)
	gameNumberString := re.FindStringSubmatch(line)[1]

	// Retrieve the list of the extractions
	extractions := trimSpacesAndSplit(strings.Split(line, ":")[1], ";")

	// For each extracted ball in each extraction build the
	// BallExtraction object with color and number
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