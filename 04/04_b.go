package main

import (
	"bufio"
)

func processPartBFile(fileScanner *bufio.Scanner) int {
	var scratchcardCounts []int
	var currentScratchCard int
	var result int
	for fileScanner.Scan() {
		// Increment by 1 the count for the current scratchcard
		scratchcardCounts = incrementScratchcardCount(scratchcardCounts, currentScratchCard, 1)

		// Get winning numbers for current scratchcard
		line := fileScanner.Text()
		lineResult := getLineWinningNumbersCount(line)

		// Increment the count of the next lineResult scratchcards
		// by the number of the current scratchcard
		for i := currentScratchCard+1; i < currentScratchCard+lineResult+1; i++ {
			scratchcardCounts = incrementScratchcardCount(scratchcardCounts, i, scratchcardCounts[currentScratchCard])
		}

		// Sum the count of the current scratchcard to the result
		result += scratchcardCounts[currentScratchCard]
		
		// Go to next scratchcard
		currentScratchCard += 1
	}

	return result
}

// Increment the specified position of the scratchcards count array by
// the given increment number, if the index is out of bound append
// the increment number to the array
func incrementScratchcardCount(scratchcardCounts []int, idx int, increment int) []int {
	if len(scratchcardCounts) <= idx {
		scratchcardCounts = append(scratchcardCounts, increment)
	} else {
		scratchcardCounts[idx] += increment
	}
	return scratchcardCounts
}
