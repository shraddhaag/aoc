package main

import (
	"fmt"
	"sort"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	intArrays := getNumArrayFromColumns(input)

	sort.Ints(intArrays[0])
	sort.Ints(intArrays[1])

	fmt.Println("answer for part 1: ", getTotalDistance(intArrays))
	fmt.Println("answer for part 2: ", getSimilarityScore(intArrays))
}

func getNumArrayFromColumns(input []string) [][]int {
	output := make([][]int, 2)

	for _, val := range input {
		nums := aoc.FetchSliceOfIntsInString(val)
		output[0] = append(output[0], nums[0])
		output[1] = append(output[1], nums[1])
	}

	return output
}

func getTotalDistance(intSlice [][]int) int {
	length := len(intSlice[0])
	sum := 0

	for i := 0; i < length; i++ {
		diff := intSlice[0][i] - intSlice[1][i]
		if diff < 0 {
			diff *= -1
		}

		sum += diff
	}

	return sum
}

func getSimilarityScore(intSlice [][]int) int {
	score := 0
	length := len(intSlice[0])

	reoccuraneCount := make(map[int]int)

	for i := 0; i < length; i++ {
		reoccuraneCount[intSlice[1][i]] += 1
	}

	for _, key := range intSlice[0] {
		if val, ok := reoccuraneCount[key]; ok {
			score += key * val
		}
	}
	return score
}
