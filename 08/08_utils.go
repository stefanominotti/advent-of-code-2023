package main

import (
	"bufio"
	"strings"
)

type Network struct {
	path         string
	nodes        map[string][2]string
}

// Create a Network object from the file
func parseFile(fileScanner *bufio.Scanner) Network {
	fileScanner.Scan()
	path := fileScanner.Text()

	nodes := map[string][2]string{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			continue
		}

		nodeName, moves := parseLine(line)
		nodes[nodeName] = moves
	}
	return Network{path, nodes}
}

func parseLine(line string) (string, [2]string) {
	splitLine := strings.Split(line, " = ")
	nodeName := splitLine[0]
	moves := [2]string{}
	copy(moves[:], strings.Split(splitLine[1][1 : len(splitLine[1])-1], ", "))
	return nodeName, moves
}

// Apply the given move from the given position
func applyMove(network Network, move rune, position string) string {
		moveIdx := 0
		if move == 'R' {
			moveIdx = 1
		}
		return network.nodes[position][moveIdx]
}
