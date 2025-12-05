package main

import (
	"fmt"
	"sort"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	validIDs, ids := getValidRangesAndAvailable(input)

	fmt.Println("answer for part 1: ", getFreshAvailable(validIDs, ids))
	fmt.Println("answer for part 2: ", getTotalFresh(validIDs))
}

func getValidRangesAndAvailable(input []string) ([][]int, []int) {
	validRanges := make([][]int, 0)
	ids := make([]int, 0)
	breakLine := 0
	for i, b := range input {
		if len(b) == 0 {
			breakLine = i
			break
		}
		a := strings.Split(b, "-")
		r := []int{aoc.FetchNumFromStringIgnoringNonNumeric(a[0]), aoc.FetchNumFromStringIgnoringNonNumeric(a[1])}
		validRanges = append(validRanges, r)
	}

	for _, b := range input[breakLine+1:] {
		ids = append(ids, aoc.FetchNumFromStringIgnoringNonNumeric(b))
	}
	return validRanges, ids
}

func getTotalFresh(validRanges [][]int) int {
	nonOverlappingRange := ensureNonOverLapping(validRanges)
	totalValidIDs := 0

	for _, a := range nonOverlappingRange {
		totalValidIDs += a[1] - a[0] + 1
	}
	return totalValidIDs
}

func ensureNonOverLapping(s [][]int) [][]int {
	// sort the slice first
	sort.Slice(s, func(i, j int) bool {
		// edge cases
		if len(s[i]) == 0 && len(s[j]) == 0 {
			return false // two empty slices - so one is not less than other i.e. false
		}
		if len(s[i]) == 0 || len(s[j]) == 0 {
			return len(s[i]) == 0 // empty slice listed "first" (change to != 0 to put them last)
		}

		// both slices len() > 0, so can test this now:
		return s[i][0] < s[j][0]
	})

	output := [][]int{}
	for _, current := range s {
		if len(output) == 0 {
			output = append(output, current)
			continue
		}

		last := output[len(output)-1]

		// if both are smaller than the last range, move on
		// since we've sorted, this is not really possible
		if current[0] < last[0] && current[1] < last[0] {
			output = append(output, current)
			continue
		}

		// if both are bigger than last range, move on
		if current[0] > last[1] && current[1] > last[1] {
			output = append(output, current)
			continue
		}

		// if we are still here that means there is def some overlap
		if current[0] < last[0] {
			last[0] = current[0]
		}

		if current[1] > last[1] {
			last[1] = current[1]
		}

		output[len(output)-1] = last
	}
	return output
}

func getFreshAvailable(validIds [][]int, ids []int) int {
	freshAvail := 0
	for _, num := range ids {
		for _, r := range validIds {
			if num >= r[0] && num <= r[1] {
				freshAvail++
				break
			}
		}
	}
	return freshAvail
}
