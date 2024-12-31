package main

import (
	"fmt"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	ans1, ans2 := getPossibleCount(getTowelsAndDesigns(input))
	fmt.Println("answer to part 1: ", ans1)
	fmt.Println("answer to part 2: ", ans2)
}

func getTowelsAndDesigns(input []string) (towels []string, designs []string) {
	for _, row := range input {

		if len(row) == 0 {
			continue
		}

		if strings.Contains(row, ",") {
			towels = append(towels, strings.Split(row, ",")...)
		} else {
			designs = append(designs, strings.TrimSpace(row))
		}
	}

	for i, t := range towels {
		towels[i] = strings.TrimSpace(t)
	}
	return
}

func isTowelPossible(towelDesign string, towelPatterns []string, possible map[string]int) int {
	if val, ok := possible[towelDesign]; ok {
		return val
	}

	isPossibleCount := 0

	for _, t := range towelPatterns {
		if len(t) > len(towelDesign) {
			continue
		}

		if strings.HasPrefix(towelDesign, t) {
			if len(t) == len(towelDesign) {
				isPossibleCount++
				continue
			}
			isPossibleCount += isTowelPossible(towelDesign[len(t):], towelPatterns, possible)
		}
	}
	possible[towelDesign] = isPossibleCount
	return isPossibleCount
}

func getPossibleCount(towels []string, designs []string) (int, int) {
	count1, count2 := 0, 0
	possible := make(map[string]int)
	for _, d := range designs {
		isPossibleCount := isTowelPossible(d, towels, possible)
		count2 += isPossibleCount
		if isPossibleCount > 0 {
			count1 += 1
		}
	}
	return count1, count2
}
