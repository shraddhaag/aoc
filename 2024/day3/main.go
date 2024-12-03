package main

import (
	"fmt"
	"regexp"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")

	sum := 0
	for _, line := range input {
		sum += findAnswer1(findValidMuls(line))
	}
	fmt.Println("answer for part 1: ", sum)

	var occurances [][]string
	for _, line := range input {
		occurances = append(occurances, findValidMulsDoDonts(line)...)
	}

	fmt.Println("answer for part 2: ", findAnswer2(occurances))
}

func findValidMuls(input string) [][]string {
	r := regexp.MustCompile(`(mul\()([\d]{1,3})[,]([0-9]{1,3})\)`)
	return r.FindAllStringSubmatch(input, -1)
}

func findValidMulsDoDonts(input string) [][]string {
	r := regexp.MustCompile(`((mul\()([\d]{1,3})[,]([0-9]{1,3})\))|(do\(\))|(don\'t\(\))`)
	return r.FindAllStringSubmatch(input, -1)
}

func findProductInPattern(input string) int {
	nums := aoc.FetchSliceOfIntsInString(input)
	return nums[0] * nums[1]
}

func findAnswer1(input [][]string) int {
	sum := 0
	for _, match := range input {
		sum += findProductInPattern(match[0])
	}
	return sum
}

func findAnswer2(input [][]string) int {
	sum := 0
	active := true
	for _, match := range input {
		if strings.Contains(match[0], "don") {
			active = false
			continue
		} else if strings.Contains(match[0], "do(") {
			active = true
			continue
		} else {
			if active {
				sum += findProductInPattern(match[0])
			}
		}
	}

	return sum
}
