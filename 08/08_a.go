package main

import (
	"bufio"
)

func processPartAFile(fileScanner *bufio.Scanner) (int, error) {
	network := parseFile(fileScanner)

	// Calculate the distance from AAA to ZZZ
	position := "AAA"
	distance := 0
	for position != "ZZZ" {
		for _, move := range network.path {
			position = applyMove(network, move, position)
			distance += 1
			if position == "ZZZ" {
				break
			}
		}		
	}

	// The number of cycles is the LCM between the distance and the length of the path
	cycles := 1
	for len(network.path)*cycles%distance != 0 {
		cycles += 1
	}

	return cycles*len(network.path), nil
}

