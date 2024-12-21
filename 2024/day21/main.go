package main

import (
	"fmt"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	numericalMap := make(map[string]aoc.Coordinates)
	numericalMap["A"] = aoc.Coordinates{2, 0}
	numericalMap["0"] = aoc.Coordinates{1, 0}
	numericalMap["1"] = aoc.Coordinates{0, 1}
	numericalMap["2"] = aoc.Coordinates{1, 1}
	numericalMap["3"] = aoc.Coordinates{2, 1}
	numericalMap["4"] = aoc.Coordinates{0, 2}
	numericalMap["5"] = aoc.Coordinates{1, 2}
	numericalMap["6"] = aoc.Coordinates{2, 2}
	numericalMap["7"] = aoc.Coordinates{0, 3}
	numericalMap["8"] = aoc.Coordinates{1, 3}
	numericalMap["9"] = aoc.Coordinates{2, 3}

	directionalMap := make(map[string]aoc.Coordinates)
	directionalMap["A"] = aoc.Coordinates{2, 1}
	directionalMap["^"] = aoc.Coordinates{1, 1}
	directionalMap["<"] = aoc.Coordinates{0, 0}
	directionalMap["v"] = aoc.Coordinates{1, 0}
	directionalMap[">"] = aoc.Coordinates{2, 0}

	fmt.Println("answer to part 1: ", getSequence(input, numericalMap, directionalMap, 2))
	fmt.Println("answer to part 1: ", getSequence(input, numericalMap, directionalMap, 25))
}

/*
// +---+---+---+
// | 7 | 8 | 9 |
// +---+---+---+
// | 4 | 5 | 6 |
// +---+---+---+
// | 1 | 2 | 3 |
// +---+---+---+
//	   | 0 | A |
//	   +---+---+
*/
func getPressesForNumericPad(input []string, start string, numericalMap map[string]aoc.Coordinates) []string {
	current := numericalMap[start]
	output := []string{}

	for _, char := range input {
		dest := numericalMap[char]
		diffX, diffY := dest.X-current.X, dest.Y-current.Y

		horizontal, vertical := []string{}, []string{}

		for i := 0; i < aoc.Abs(diffX); i++ {
			if diffX >= 0 {
				horizontal = append(horizontal, ">")
			} else {
				horizontal = append(horizontal, "<")
			}
		}

		for i := 0; i < aoc.Abs(diffY); i++ {
			if diffY >= 0 {
				vertical = append(vertical, "^")
			} else {
				vertical = append(vertical, "v")
			}
		}

		// prioritisation order:
		// 1. moving with least turns
		// 2. moving < over ^ over v over >

		if current.Y == 0 && dest.X == 0 {
			output = append(output, vertical...)
			output = append(output, horizontal...)
		} else if current.X == 0 && dest.Y == 0 {
			output = append(output, horizontal...)
			output = append(output, vertical...)
		} else if diffX < 0 {
			output = append(output, horizontal...)
			output = append(output, vertical...)
		} else if diffX >= 0 {
			output = append(output, vertical...)
			output = append(output, horizontal...)
		}

		current = dest
		output = append(output, "A")
	}
	return output
}

/*
//     +---+---+
//     | ^ | A |
// +---+---+---+
// | < | v | > |
// +---+---+---+
*/
func getPressesForDirectionalPad(input []string, start string, directionlMap map[string]aoc.Coordinates) []string {
	current := directionlMap[start]
	output := []string{}

	for _, char := range input {
		dest := directionlMap[char]
		diffX, diffY := dest.X-current.X, dest.Y-current.Y

		horizontal, vertical := []string{}, []string{}

		for i := 0; i < aoc.Abs(diffX); i++ {
			if diffX >= 0 {
				horizontal = append(horizontal, ">")
			} else {
				horizontal = append(horizontal, "<")
			}
		}

		for i := 0; i < aoc.Abs(diffY); i++ {
			if diffY >= 0 {
				vertical = append(vertical, "^")
			} else {
				vertical = append(vertical, "v")
			}
		}

		// prioritisation order:
		// 1. moving with least turns
		// 2. moving < over ^ over v over >

		if current.X == 0 && dest.Y == 1 {
			output = append(output, horizontal...)
			output = append(output, vertical...)
		} else if current.Y == 1 && dest.X == 0 {
			output = append(output, vertical...)
			output = append(output, horizontal...)
		} else if diffX < 0 {
			output = append(output, horizontal...)
			output = append(output, vertical...)
		} else if diffX >= 0 {
			output = append(output, vertical...)
			output = append(output, horizontal...)
		}
		current = dest
		output = append(output, "A")
	}
	return output
}

func getSequence(input []string, numericalMap, directionalMap map[string]aoc.Coordinates, robots int) int {
	count := 0
	cache := make(map[string][]int)
	for _, line := range input {
		row := strings.Split(line, "")
		seq1 := getPressesForNumericPad(row, "A", numericalMap)
		num := getCountAfterRobots(seq1, robots, 1, cache, directionalMap)
		count += aoc.FetchNumFromStringIgnoringNonNumeric(line) * num
	}
	return count
}

func getCountAfterRobots(input []string, maxRobots int, robot int, cache map[string][]int, directionalMap map[string]aoc.Coordinates) int {
	if val, ok := cache[strings.Join(input, "")]; ok {
		if val[robot-1] != 0 {
			return val[robot-1]
		}
	} else {
		cache[strings.Join(input, "")] = make([]int, maxRobots)
	}

	seq := getPressesForDirectionalPad(input, "A", directionalMap)
	cache[strings.Join(input, "")][0] = len(seq)

	if robot == maxRobots {
		return len(seq)
	}

	splitSeq := getIndividualSteps(seq)

	count := 0
	for _, s := range splitSeq {
		c := getCountAfterRobots(s, maxRobots, robot+1, cache, directionalMap)
		if _, ok := cache[strings.Join(s, "")]; !ok {
			cache[strings.Join(s, "")] = make([]int, maxRobots)
		}
		cache[strings.Join(s, "")][0] = c
		count += c
	}

	cache[strings.Join(input, "")][robot-1] = count
	return count
}

func getIndividualSteps(input []string) [][]string {
	output := [][]string{}
	current := []string{}
	for _, char := range input {
		current = append(current, char)

		if char == "A" {
			output = append(output, current)
			current = []string{}
		}
	}
	return output
}
