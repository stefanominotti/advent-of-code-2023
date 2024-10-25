package main

import (
	"advent-of-code/utils"
)

func main() {
	processingFunctionPartA := utils.ProcessFileLineByLineFunction(processPartALine)
	processingFunctionPartB := processPartBFile
	utils.RunChallenge(processingFunctionPartA, processingFunctionPartB)
}
