package main

import (
	"advent-of-code/utils"
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
func createGame(line string) Game {
	// Retrieve game number
	re := regexp.MustCompile(gameRegexPattern)
	gameNumberString := re.FindStringSubmatch(line)[1]

	// Retrieve the list of the extractions
	extractions := utils.TrimSpacesAndSplit(strings.Split(line, ":")[1], ";")

	// For each extracted ball in each extraction build the
	// BallExtraction object with color and number
	var balls []BallExtraction
	for _, extraction := range extractions {
		extractedBalls := utils.TrimSpacesAndSplit(extraction, ",")

		for _, extractedBall := range extractedBalls {
			splitBall := utils.TrimSpacesAndSplit(extractedBall, " ")

			color := splitBall[1]
			number, err := strconv.Atoi(splitBall[0])
			if err != nil {
				panic(err)
			}
			
			balls = append(balls, BallExtraction{color, number})
		}
	}

	gameNumber, err := strconv.Atoi(gameNumberString)
	if err != nil {
		panic(err)
	}
	return Game{gameNumber, balls}
}
