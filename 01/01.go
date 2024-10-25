package main

import (
	"advent-of-code/utils"
)

func main() {
	processingFunctionPartA := utils.ProcessFileLineByLineFunction(processPartALine)
	processingFunctionPartB := utils.ProcessFileLineByLineFunction(processPartBLine)
	utils.RunChallenge(processingFunctionPartA, processingFunctionPartB)
}
