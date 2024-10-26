package main

import (
	"math"
)

// Calculate the min and max button holding time by solving
// x^2 - x*time + distance < 0
func calculateMinAndMaxButtonHoldingTime(distance int, totalTime int) (int, int) {
	minTime := math.Floor(0.5 * (float64(totalTime) - math.Sqrt(float64(totalTime)*float64(totalTime) - 4*float64(distance))) + 1)
	maxTime := math.Ceil(0.5 * (float64(totalTime) + math.Sqrt(float64(totalTime)*float64(totalTime) - 4*float64(distance))) - 1)
	return int(minTime), int(maxTime)
}