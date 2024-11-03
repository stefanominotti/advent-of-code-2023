package main

import (
	"bufio"
	"slices"
)

func processPartBFile(fileScanner *bufio.Scanner) int {
	network := parseFile(fileScanner)

	startingNodes := []string{}
	for node := range network.nodes {
		if (node[2] == 'A') {
			startingNodes = append(startingNodes, node)
		}
	}

	// Calculate the minimum number of cycles for all the 
	// starting nodes to reach an ending node
	minCyclesPerNode := []int{}
	for _, startingNode := range startingNodes {
		position := startingNode
		cycles := 0
		for position[2] != 'Z' {
			for _, move := range network.path {
				position = applyMove(network, move, position)
			}
			cycles += 1
		}
		minCyclesPerNode = append(minCyclesPerNode, cycles)
	}

	// The number of cycles is the LCM between the minimum 
	// cycles number of all the starting nodes
	minCyclesNeeded := slices.Max(minCyclesPerNode)
	cycles := minCyclesNeeded
	for !areCyclesValid(cycles, minCyclesPerNode) {
		cycles += minCyclesNeeded
	}

	return cycles*len(network.path)
}

// The function checks whether cycles is a multiplier of all 
// the minimum cycles numbers of the starting nodes
func areCyclesValid(cycles int, minCyclesPerNode []int) bool{
	for _, nodeMinCycles := range minCyclesPerNode {
		if cycles % nodeMinCycles != 0 {
			return false
		}
	}
	return true
}
