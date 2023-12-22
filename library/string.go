package aoc

import (
	"strconv"
	"strings"
	"unicode"
)

func FetchSliceOfIntsInString(line string) []int {
	nums := []int{}
	var build strings.Builder
	isNegative := false
	for _, char := range line {
		if unicode.IsDigit(char) {
			build.WriteRune(char)
		}

		if char == '-' {
			isNegative = true
		}

		if (char == ' ' || char == ',' || char == '~') && build.Len() != 0 {
			localNum, err := strconv.ParseInt(build.String(), 10, 64)
			if err != nil {
				panic(err)
			}
			if isNegative {
				localNum *= -1
			}
			nums = append(nums, int(localNum))
			build.Reset()
			isNegative = false
		}
	}
	if build.Len() != 0 {
		localNum, err := strconv.ParseInt(build.String(), 10, 64)
		if err != nil {
			panic(err)
		}
		if isNegative {
			localNum *= -1
		}
		nums = append(nums, int(localNum))
		build.Reset()
	}
	return nums
}

func FetchNumFromStringIgnoringNonNumeric(line string) int {
	var build strings.Builder
	for _, char := range line {
		if unicode.IsDigit(char) {
			build.WriteRune(char)
		}
	}
	if build.Len() != 0 {
		localNum, err := strconv.ParseInt(build.String(), 10, 64)
		if err != nil {
			panic(err)
		}
		return int(localNum)
	}
	return 0
}

func SplitStringAfter(input string, length int) (output []string) {
	startIndex := 0
	for startIndex < len(input) {
		output = append(output, input[startIndex:startIndex+length])
		startIndex += length
	}
	return
}

func Get2DGrid(input []string) (grid [][]string) {
	for _, line := range input {
		grid = append(grid, strings.Split(line, ""))
	}
	return
}
