package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	inputMap, evalMap := createInputMap(input)
	evaluateWires(inputMap, evalMap)
	ans1 := getDecimal10FromGatesStartingWith(inputMap, "z")
	fmt.Println("answer for part 1: ", ans1)
	fmt.Println("answer for part 2: ", validateRippleCarryAdderRules(getGateConnections(input), evalMap))
}

func createInputMap(input []string) (map[string]int, map[string][][]string) {
	inputMap := make(map[string]int)
	evalMap := make(map[string][][]string)

	isInitialState := true
	for _, line := range input {
		if len(line) == 0 {
			isInitialState = false
		}

		if isInitialState {
			inputMap[line[:3]] = aoc.FetchNumFromStringIgnoringNonNumeric(line[3:])
		}

		if !isInitialState && len(line) > 3 {
			vals := strings.Split(line, " ")
			evalMap[vals[0]] = append(evalMap[vals[0]], []string{vals[2], vals[1], vals[4]})
			evalMap[vals[2]] = append(evalMap[vals[2]], []string{vals[0], vals[1], vals[4]})
		}
	}
	return inputMap, evalMap
}

func evalGate(a, gate, b, output string, inputMap map[string]int) {
	switch gate {
	case "AND":
		inputMap[output] = inputMap[a] & inputMap[b]
	case "OR":
		inputMap[output] = inputMap[a] | inputMap[b]
	case "XOR":
		inputMap[output] = inputMap[a] ^ inputMap[b]
	}
}

func getDecimal10FromGatesStartingWith(inputMap map[string]int, char string) int64 {
	resultZ, result := []string{}, []string{}
	for key, _ := range inputMap {
		if string(key[0]) == char {
			resultZ = append(resultZ, key)
		}
	}

	sort.Strings(resultZ)
	for i := len(resultZ) - 1; i >= 0; i-- {
		result = append(result, fmt.Sprintf("%d", inputMap[resultZ[i]]))
	}
	intValue, _ := strconv.ParseInt(strings.Join(result, ""), 2, 64)
	return intValue
}

func evaluateWires(inputMap map[string]int, evalInst map[string][][]string) {
	next := []string{}
	visited := make(map[string]struct{})
	for key, _ := range inputMap {
		next = append(next, key)
	}

	for len(next) != 0 {
		// fmt.Println(next)
		current := next[0]
		next = next[1:]

		if _, ok := visited[current]; ok {
			continue
		}

		if _, ok := evalInst[current]; !ok {
			continue
		}

		if _, ok := inputMap[current]; !ok {
			next = append(next, current)
			continue
		}

		found := false
		for _, combination := range evalInst[current] {
			if _, ok := inputMap[combination[0]]; !ok {
				continue
			}
			found = true
			evalGate(current, combination[1], combination[0], combination[2], inputMap)
			if _, ok := visited[combination[2]]; !ok {
				next = append(next, combination[2])
			}
		}

		if !found {
			next = append(next, current)
			continue
		}

		visited[current] = struct{}{}
	}
}

func expectedOutput(inputMap map[string]int) int {
	a := getDecimal10FromGatesStartingWith(inputMap, "x")
	b := getDecimal10FromGatesStartingWith(inputMap, "y")
	return int(a + b)
}

func getGateConnections(input []string) []string {
	for index, line := range input {
		if len(line) == 0 {
			return input[index+1:]
		}
	}
	return input
}

// validateRippleCarryAdderRules checks the following rules:
// 1. If the output of a gate is z, then the operation has to be XOR unless it is the last bit.
// 2. If the output of a gate is not z and the inputs are not x, y then it has to be AND / OR, but not XOR.
// 3. If you have a XOR gate with inputs x, y, there must be another XOR gate with this gate as an input.
// Search through all gates for an XOR-gate with this gate as an input; if it does not exist, your (original) XOR gate is faulty.
// 4. Similarly, if you have an AND-gate, there must be an OR-gate with this gate as an input. If that gate doesn't exist, the original AND gate is faulty.
// from: https://www.reddit.com/r/adventofcode/comments/1hla5ql/2024_day_24_part_2_a_guide_on_the_idea_behind_the/
func validateRippleCarryAdderRules(input []string, instMap map[string][][]string) string {
	faulty := []string{}
	for _, line := range input {
		split := strings.Split(line, " ")
		a, gate, b, output := split[0], split[1], split[2], split[4]

		//  If the output of a gate is z, then the operation has to be XOR unless it is the last bit.
		if output[0] == 'z' && gate != "XOR" && output != "z45" {
			faulty = append(faulty, output)
			continue
		}

		// If the output of a gate is not z and the inputs are not x, y then it has to be AND / OR, but not XOR.
		if output[0] != 'z' && a[0] != 'x' && a[0] != 'y' && b[0] != 'x' && b[0] != 'y' && gate == "XOR" {
			faulty = append(faulty, output)
			continue
		}

		// If you have a XOR gate with inputs x, y, there must be another XOR gate with this gate as an input.
		// Search through all gates for an XOR-gate with this gate as an input; if it does not exist, your (original) XOR gate is faulty.
		if gate == "XOR" && ((a[0] == 'x' && b[0] == 'y') || (a[0] == 'y' && b[0] == 'x')) &&
			a != "x00" && b != "x00" && a != "y00" && b != "y00" {
			if _, ok := instMap[output]; !ok {
				faulty = append(faulty, output)
				continue
			}

			isValid := false
			for _, poss := range instMap[output] {
				if poss[1] == "XOR" {
					isValid = true
					break
				}
			}

			if !isValid {
				faulty = append(faulty, output)
				continue
			}
		}

		// if you have an AND-gate, there must be an OR-gate with this gate as an input.
		// If that gate doesn't exist, the original AND gate is faulty.
		if gate == "AND" && ((a[0] == 'x' && b[0] == 'y') || (a[0] == 'y' && b[0] == 'x')) &&
			a != "x00" && b != "x00" && a != "y00" && b != "y00" {
			if _, ok := instMap[output]; !ok {
				faulty = append(faulty, output)
				continue
			}

			isValid := false
			for _, poss := range instMap[output] {
				if poss[1] == "OR" {
					isValid = true
					break
				}
			}

			if !isValid {
				faulty = append(faulty, output)
				continue
			}
		}
	}
	sort.Strings(faulty)
	return strings.Join(faulty, ",")
}
