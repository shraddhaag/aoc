package main

import (
	"fmt"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	inputMap, a := createInputMap(input[2:])
	fmt.Println("answer for part 1: ", countStep(input[0], inputMap, "AAA", isEnd))
	fmt.Println("answer for part 2: ", countStep2(input[0], inputMap, a))
}

type node struct {
	left  string
	right string
}

func createInputMap(input []string) (map[string]node, []string) {
	output := map[string]node{}
	sliceOfNodesEndingInA := []string{}
	for _, line := range input {
		current := line[:3]
		output[current] = node{
			left:  line[7:10],
			right: line[12:15],
		}
		if current[2] == 'A' {
			sliceOfNodesEndingInA = append(sliceOfNodesEndingInA, current)
		}
	}
	return output, sliceOfNodesEndingInA
}

func countStep(inst string, inputMap map[string]node, start string, isEnd func(string) bool) int {
	count := 0
	current := start
	pointer := 0
	for !isEnd(current) {
		if inst[pointer] == 'R' {
			current = inputMap[current].right
		} else {
			current = inputMap[current].left
		}
		if pointer < len(inst)-1 {
			pointer++
		} else {
			pointer = 0
		}
		count++
	}
	return count
}

func countStep2(inst string, inputMap map[string]node, start []string) int {
	individualCount := []int{}
	for _, i := range start {
		individualCount = append(individualCount, countStep(inst, inputMap, i, isEnd2))
	}
	// LCM works here because of the nature of the input:
	// 	1. Each A has exactly one Z in its path.
	// 	2. Number of steps to reach A -> Z = Number of steps to reach Z -> Z
	//	3. All As reach their Zs at the same step in the input instruction. For me its 268.
	//
	// nice reddit thread explanation for this - https://www.reddit.com/r/adventofcode/comments/18dfpub/2023_day_8_part_2_why_is_spoiler_correct/
	// specific answers I liked:
	// 		- https://www.reddit.com/r/adventofcode/comments/18dfpub/comment/kcgyhfi/?utm_source=share&utm_medium=web2x&context=3
	// 		- https://www.reddit.com/r/adventofcode/comments/18dfpub/comment/kch1h0r/?utm_source=share&utm_medium=web2x&context=3
	return aoc.LCM(individualCount)
}

func isEnd(input string) bool {
	return input == "ZZZ"
}

func isEnd2(input string) bool {
	return input[2] == 'Z'
}
