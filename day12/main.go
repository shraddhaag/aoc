package main

import (
	"fmt"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")

	fmt.Println("answer for part 1: ", ans(input, 1))
	fmt.Println("answer for part 2: ", ans(input, 5))
}

func ans(input []string, times int) int {
	sum := 0
	cache := map[string]int{}
	for _, line := range input {
		sum += processLine(line, cache, times)
	}
	return sum
}

func processLine(input string, cache map[string]int, times int) int {
	stringRep, numRep := fetchInputs(input, times)
	return getCount(stringRep, numRep, cache)
}

func fetchInputs(input string, times int) (string, []int) {
	stringRep := input[:strings.Index(input, " ")]
	numRep := aoc.FetchSliceOfIntsInString(input[strings.Index(input, " "):])

	numRep2 := []int{}
	var build strings.Builder
	for i := 0; i < times; i++ {
		build.WriteString(stringRep)
		if i != (times - 1) {
			build.WriteString("?")
		}
		numRep2 = append(numRep2, numRep...)
	}
	stringRep = build.String()
	return stringRep, numRep2
}

// top level approach used:
// - chose direction to go based on first character of the input string:
//   - '.' -> ignore
//   - '#' -> look ahead if consequetive '#'s satisfy the first group.
//   - '?' -> 2 recursive call by replacing once with '.' and then '#'.
//
// - to reduce the number of recursive calls, we also have a cache to store all evaluated results.
func getCount(line string, group []int, cache map[string]int) int {

	// first, try cache hit ie, check if the current case is already present in cache
	// if yes, return early
	cacheEntry := line + aoc.ConvertIntSliceToString(group, ",")
	if v, ok := cache[cacheEntry]; ok {
		return v
	}

	// check if all groups have been evaluated and input string has no '#' left
	// if yes, this is a valid combination, count can be increased by 1
	if len(group) == 0 && !strings.Contains(line, "#") {
		cache[cacheEntry] = 1
		return 1
		// we can fail early in all the below conditions:
		// 		- input string has '#' left but no groups left
		// 		- input string is exhausted
		// 		- input string length is less than number of groups to satisfy
	} else if (len(group) == 0) || (len(line) == 0) || (len(line) < len(group)) {
		cache[cacheEntry] = 0
		return 0
	}

	var count int
	// if first character of input string is a '?': try both cases -
	// 		- replace with '.'
	// 		- replace with '#'
	if line[0] == '?' {
		count += getCount("#"+line[1:], group, cache)
		count += getCount("."+line[1:], group, cache)
	}

	// if first character of input string is a '.', we can safely ignore it and proceed
	if line[0] == '.' {
		count += getCount(line[1:], group, cache)
	}

	// if first charcter of input string is a '#', then check if we can satisfy the first group
	if line[0] == '#' {

		// count number of consecutive '#' till the first non '#' character
		// check if this count satisfies the first group
		c := 0
		for _, char := range line {
			if char == '#' {
				c++
			} else {
				break
			}
		}

		// if '#' count satisfies the first group and end of input string:
		// invoke the func with an empty string and group minus the first occurance
		// (Note: we don't just register a success here because there might bemore groups to match
		// that the input string does not satisfy)
		if c == group[0] && (len(line) == c) {
			count += getCount("", group[1:], cache)
			// if '#' count satisfies the first group:
			// invoke the func with the rest of the input string and remaining groups left to evaluate
		} else if c == group[0] {
			count += getCount(line[c+1:], group[1:], cache)
		}

		// if '#' count does not satify the first group and:
		// 		- entire input string is '#': ie no further chance to find the first group
		// 		- count of `#` is greater than first group: since order of groups is important, we can fail early
		if c != group[0] && (c == len(line) || c > group[0]) {
			cache[cacheEntry] = 0
			return 0
		}

		// if '#' count is less than first group and the `#`s are followed by a `.`: fail
		if c < group[0] && line[c] == '.' {
			return 0
		}

		// if '#' count is less than first ground and the `#`s are followed by a `?`:
		// replace '?' with '#' and invoke the func again
		if c < group[0] && line[c] == '?' {
			count += getCount(line[:c]+"#"+line[c+1:], group, cache)
		}

	}
	cache[cacheEntry] = count
	return count
}
