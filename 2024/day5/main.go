package main

import (
	"fmt"
	"reflect"
	"sort"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	rules, updates := getRulesAndUpdates(input)
	// ans1, ans2 := ans(updates, getRulesMap(rules))
	ans1, ans2 := alternativeAns(getRulesMap(rules), updates)
	fmt.Println("answer for part 1: ", ans1)
	fmt.Println("answer for part 2: ", ans2)
}

func getRulesAndUpdates(input []string) ([][]int, [][]int) {
	var rules, updates [][]int

	isRules := true
	for _, row := range input {
		if len(row) == 0 {
			isRules = false
			continue
		}

		if isRules {
			rules = append(rules, aoc.FetchSliceOfIntsInString(row))
		} else {
			updates = append(updates, aoc.FetchSliceOfIntsInString(row))
		}
	}
	return rules, updates
}

func getRulesMap(rules [][]int) map[int][]int {
	ruleMap := make(map[int][]int)

	for _, rule := range rules {
		if _, ok := ruleMap[rule[0]]; ok {
			ruleMap[rule[0]] = append(ruleMap[rule[0]], rule[1])
		} else {
			ruleMap[rule[0]] = []int{rule[1]}
		}
	}
	return ruleMap
}

func isValidUpdate(ruleMap map[int][]int, update []int) bool {
	for i := 0; i < len(update)-1; i++ {
		if _, ok := ruleMap[update[i]]; !ok {
			return false
		}
		for j := i + 1; j < len(update); j++ {
			found := false
			for _, element := range ruleMap[update[i]] {
				if element == update[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

func ans(updates [][]int, ruleMap map[int][]int) (int, int) {
	sum, fixedSum := 0, 0
	for _, update := range updates {
		if isValidUpdate(ruleMap, update) {
			sum += update[len(update)/2]
		} else {
			fixedSum += getMiddleNumberAfterFixingUpdate(update, ruleMap)
		}
	}
	return sum, fixedSum
}

func getMiddleNumberAfterFixingUpdate(update []int, ruleMap map[int][]int) int {
	fixedUpdate := make([]int, len(update))
	for i := 0; i < len(update); i++ {
		countFound := 0
		if _, ok := ruleMap[update[i]]; !ok {
			// means it should not come before any and should be last
			countFound = 0
		} else {
			// iterate over all the numbers in the update (except itself)
			// find count of all numbers that should come after it according to the rules
			for j := 0; j < len(update); j++ {
				if i == j {
					continue
				}
				found := false
				for _, element := range ruleMap[update[i]] {
					if element == update[j] {
						found = true
						break
					}
				}
				if found {
					countFound++
				}
			}
		}

		// the count of numbers that are found is the count of numbers that should come after it
		// so the current number's position = length of update - number of characters found
		// this is such that all the found characters have space to come after it.
		fixedUpdate[len(update)-countFound-1] = update[i]

		// if the middle number is found, exit early
		if fixedUpdate[len(update)/2] != 0 {
			return fixedUpdate[len(update)/2]
		}
	}
	return fixedUpdate[len(update)/2]
}

// alternative approach: the problem is essentially a sorting problem and
// the rules given at the beginning are the rules for the sort.
// we just need to write a custom sort function using the rules.
func alternativeAns(rulesMap map[int][]int, updates [][]int) (int, int) {
	sum, fixedSum := 0, 0
	for _, update := range updates {

		sortedUpdate := make([]int, len(update))
		copy(sortedUpdate, update)

		sort.Slice(sortedUpdate, func(i, j int) bool {
			return customLess(rulesMap, sortedUpdate, i, j)
		})

		if reflect.DeepEqual(update, sortedUpdate) {
			sum += update[len(update)/2]
		} else {
			fixedSum += sortedUpdate[len(update)/2]
		}
	}
	return sum, fixedSum
}

func customLess(ruleMap map[int][]int, update []int, i, j int) bool {
	if _, ok := ruleMap[update[i]]; ok {
		for _, char := range ruleMap[update[i]] {
			if char == update[j] {
				return true
			}
		}
	}
	return false
}
