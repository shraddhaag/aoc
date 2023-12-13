package main

import (
	"fmt"
	"reflect"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	sum := 0
	sum1 := 0
	cache := map[string]int{}
	for _, pattern := range getPatterns(input) {
		count := evaluatePattern(pattern, 0)
		count2 := getAlternatePattern(pattern, 0, 0, cache, count)
		sum += count
		sum1 += count2
	}

	fmt.Println("answer for part 1: ", sum)
	fmt.Println("answer for part 2: ", sum1)
}

func getPatterns(input []string) [][]string {
	output := [][]string{}
	local := []string{}
	for _, line := range input {
		if len(line) != 0 {
			local = append(local, line)
			continue
		}
		output = append(output, local)
		local = []string{}
	}
	if len(local) != 0 {
		output = append(output, local)
	}
	return output
}

func evaluatePattern(input []string, old int) int {
	palindromIndexes := getInitialIndexesToCheck(len(input[0]))
	for _, line := range input {
		new := []int{}
		for i := 0; i < len(palindromIndexes); i++ {
			if isPalendromOnIndex(strings.Split(line, ""), palindromIndexes[i]) {
				new = append(new, palindromIndexes[i])
			}
		}

		palindromIndexes = new
		if len(palindromIndexes) == 0 {
			break
		}
	}

	for _, p := range palindromIndexes {
		if p != old {
			return p
		}
	}

	for i := 0; i < len(input); i++ {
		if isPalendromOnIndex(input, i) {
			palindromIndexes = append(palindromIndexes, i*100)
		}
	}

	for _, p := range palindromIndexes {
		if p != old {
			return p
		}
	}

	return 0
}

func isPalendromOnIndex(input []string, midIndex int) bool {
	if midIndex == 0 {
		return false
	}

	var actualString []string
	if midIndex <= (len(input) / 2) {
		actualString = input[:(midIndex * 2)]
	} else {
		actualString = input[2*midIndex-len(input):]
	}
	var reverseString []string
	j := len(actualString) - 1
	for j >= 0 {
		if j >= 0 {
			reverseString = append(reverseString, actualString[j])
			j--
		}
	}

	return reflect.DeepEqual(actualString, reverseString)
}

func getInitialIndexesToCheck(length int) []int {
	output := []int{}
	for i := 0; i < length; i++ {
		output = append(output, i)
	}
	return output
}

func getAlternatePattern(input []string, i, j int, cache map[string]int, old int) int {

	// flip (i, j) bit
	flipChar(input, i, j)

	// check if entry is present in cache
	if v, ok := cache[strings.Join(input, "")]; ok {
		return v
	}

	e := evaluatePattern(input, old)
	cache[strings.Join(input, "")] = e
	if e != 0 {
		return e
	}

	// flip (i, j) bit back
	flipChar(input, i, j)

	if j == len(input[i])-1 {
		j = 0
		i = i + 1
	} else {
		j++
	}

	return getAlternatePattern(input, i, j, cache, old)
}

func flipChar(input []string, i, j int) {
	line := strings.Split(input[i], "")
	if line[j] == "#" {
		line[j] = "."
	} else {
		line[j] = "#"
	}
	input[i] = strings.Join(line, "")
}
