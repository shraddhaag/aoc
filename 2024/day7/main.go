package main

import (
	"fmt"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	ans1, ans2 := ans(input)
	fmt.Println("answer for part 1: ", ans1)
	fmt.Println("answer for part 2: ", ans2)
}

func ans(input []string) (int, int) {
	ans1, ans2 := 0, 0
	for _, row := range input {
		nums := aoc.FetchSliceOfIntsInString(row)

		if isSumAMatch(nums[0], 0, nums[1:], false) {
			ans1 += nums[0]
		}

		if isSumAMatch(nums[0], 0, nums[1:], true) {
			ans2 += nums[0]
		}
	}
	return ans1, ans2
}

func calculate(a, b int, operation byte) int {
	calculation := 0
	switch operation {
	case '+':
		calculation = a + b
	case '*':
		calculation = a * b
	case '|':
		mul, q := 10, 10
		for q != 0 {
			q = b / mul
			if q > 0 {
				mul *= 10
			}
		}
		calculation = (a * mul) + b
	}
	return calculation
}

func isSumAMatch(expectedSum, sum int, input []int, isPart2 bool) bool {
	if len(input) == 0 {
		return sum == expectedSum
	}

	if sum > expectedSum {
		return false
	}

	if isSumAMatch(expectedSum, calculate(sum, input[0], '+'), input[1:], isPart2) {
		return true
	}

	if isPart2 && isSumAMatch(expectedSum, calculate(sum, input[0], '|'), input[1:], isPart2) {
		return true
	}
	return isSumAMatch(expectedSum, calculate(sum, input[0], '*'), input[1:], isPart2)
}
