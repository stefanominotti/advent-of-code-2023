package main

import (
	"bufio"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cardValues [5]int
	strength   int
	bet        int
}

func processFile(fileScanner *bufio.Scanner, parseLine func(line string) Hand) int {
	hands := parseFile(fileScanner, parseLine)

	// Sort hands by strength
	slices.SortFunc(hands, handsSortFunc)

	result := 0
	for idx, hand := range hands {
		result += (idx + 1) * hand.bet
	}
	return result
}

// Given two hands returns -1 if a is stronger than b,
// 1 if a is weaker than b, 0 if they have the same strength;
// the check is done first on hand strength and in case of tie
// comparing the card values
func handsSortFunc(a Hand, b Hand) int {
	if a.strength < b.strength {
		return -1
	}

	if a.strength > b.strength {
		return 1
	}

	for i := 0; i < 5; i++ {
		if a.cardValues[i] < b.cardValues[i] {
			return -1
		}

		if a.cardValues[i] > b.cardValues[i] {
			return 1
		}
	}

	return 0
}

func parseFile(fileScanner *bufio.Scanner, parseLine func(line string) Hand) []Hand {
	var hands []Hand

	for fileScanner.Scan() {
		line := fileScanner.Text()

		hand := parseLine(line)
		hands = append(hands, hand)
	}
	return hands
}

// Create an Hand object from a line considering the card power
// from the cardPowerMap parameter and calculating the strength
// of the hand using calculateHandStrength function
func parseLine(line string, cardPowerMap map[string]int, calculateHandStrength func(handCardsCount []int) int) Hand {
	splitLine := strings.Split(line, " ")

	cardsCount := [13]int{}
	cards := [5]int{}
	for i := 0; i < 5; i++ {
		cardString := string(splitLine[0][i])
		cardPower := cardPowerMap[cardString]
		cardsCount[cardPower] += 1
		cards[i] = cardPower
	}

	bet, err := strconv.Atoi(splitLine[1])
	if err != nil {
		panic(err)
	}

	return Hand{cards, calculateHandStrength(cardsCount[:]), bet}
}

// Calculate the strength of an hand by checking the number of copies
// of the card with the most copies (eventually plus the jollies, for part B);
// in case of 3 and 2 copies check also the second card with the most copies
// for full house or two pair
func calculateHandStrength(handCardsCount []int, jolliesCount int) int {
	slices.Sort(handCardsCount)
	slices.Reverse(handCardsCount)

	switch handCardsCount[0] + jolliesCount {
	case 1:
		return 1
	case 2:
		switch handCardsCount[1] {
		case 2:
			return 3
		default:
			return 2
		}
	case 3:
		switch handCardsCount[1] {
		case 2:
			return 5
		default:
			return 4
		}
	case 4:
		return 6
	case 5:
		return 7
	default:
		return 0
	}
}