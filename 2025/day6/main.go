package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	add, mul := parseInput(input)

	fmt.Println("answer for part 1: ", ans1(add, mul))
	fmt.Println("answer for part 2: ", ans2(input))
}

func parseInput(input []string) ([][]int, [][]int) {
	modInput := [][]int{}
	add, mul := [][]int{}, [][]int{}
	for i, nums := range input {
		if i == len(input)-1 {
			break
		}
		modInput = append(modInput, aoc.FetchSliceOfIntsInString(nums))
	}

	signs := strings.ReplaceAll(input[len(input)-1], " ", "")

	for j, sign := range signs {
		nums := []int{}
		for i, _ := range modInput {
			if len(modInput[i]) >= j+1 {
				nums = append(nums, modInput[i][j])
			}
		}
		switch sign {
		case '+':
			add = append(add, nums)
		case '*':
			mul = append(mul, nums)
		}
	}
	return add, mul
}

func ans1(add, mul [][]int) int {
	total := 0
	for _, nums := range add {
		new := 0
		for _, n := range nums {
			new += n
		}
		total += new
	}

	for _, nums := range mul {
		new := 1
		for _, n := range nums {
			new *= n
		}
		total += new
	}
	return total
}

func ans2(input []string) int {
	output := 0
	currSign := '@'
	currMulOrAdd := 0
	for j, sign := range input[len(input)-1] {
		if sign != ' ' {
			output += currMulOrAdd
			currSign = sign
			if sign == '*' {
				currMulOrAdd = 1
			} else {
				currMulOrAdd = 0
			}
		}
		num := 0
		for i := len(input) - 2; i >= 0; i-- {
			if (len(input[i]) >= j+1) && (input[i][j] != ' ') {
				power := 0
				if num != 0 {
					power = len(strconv.Itoa(num))
				}
				num += aoc.FetchNumFromStringIgnoringNonNumeric(string(input[i][j])) * int(math.Pow(10, float64(power)))
			}
		}
		if num == 0 {
			continue
		}
		switch currSign {
		case '*':
			currMulOrAdd *= num
		case '+':
			currMulOrAdd += num
		}
	}
	return output + currMulOrAdd
}
