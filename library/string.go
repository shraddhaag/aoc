package aoc

import (
	"strconv"
	"strings"
	"unicode"
)

func FetchSliceOfIntsInString(line string) []int {
	nums := []int{}
	var build strings.Builder
	for _, char := range line {
		if unicode.IsDigit(char) {
			build.WriteRune(char)
		}

		if char == ' ' && build.Len() != 0 {
			localNum, err := strconv.ParseInt(build.String(), 10, 64)
			if err != nil {
				panic(err)
			}
			nums = append(nums, int(localNum))
			build.Reset()
		}
	}
	if build.Len() != 0 {
		localNum, err := strconv.ParseInt(build.String(), 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, int(localNum))
		build.Reset()
	}
	return nums
}
