package main

import (
	"fmt"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	fmt.Println("answer for part 1: ", ans1(input))
	fmt.Println("answer for part 2: ", ans2(input))
}

func ans1(input []string) int {
	sum := 0
	for _, line := range input {
		sum += evaluateHistory(line, extendEnd)
	}
	return sum
}

func ans2(input []string) int {
	sum := 0
	for _, line := range input {
		sum += evaluateHistory(line, extendBeginning)
	}
	return sum
}

func evaluateHistory(input string, extend func([]int, int) int) int {
	inputSlice := aoc.FetchSliceOfIntsInString(input)
	next := extend(inputSlice, predictNext(inputSlice, extend))
	return next

}

func extendEnd(inputSlice []int, last int) int {
	return inputSlice[len(inputSlice)-1] + last
}

func extendBeginning(inputSlice []int, last int) int {
	return inputSlice[0] - last
}

func predictNext(input []int, extend func([]int, int) int) int {
	if isAllZero(input) {
		return 0
	}

	if len(input) <= 1 {
		return input[0]
	}

	diff := findDiff(input)
	last := predictNext(diff, extend)
	return extend(diff, last)
}

func findDiff(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	output := []int{}
	for i := 0; i < len(input)-1; i++ {
		output = append(output, input[i+1]-input[i])
	}
	return output
}

func isAllZero(input []int) bool {
	output := 0
	for i := 0; i < len(input); i++ {
		if input[i] == 0 {
			output++
		}
	}

	return output == len(input)
}
