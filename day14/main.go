package main

import (
	"fmt"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")

	fmt.Println("answer for part 1: ", findWeight(getNorthTitledGrid(input)))
	fmt.Println("answer for part 2: ", findWeight(getGridAfterXCycles(input, 1000000000)))
}

var cache map[string]string

func getNorthTitledGrid(input []string) []string {
	grid := aoc.Get2DGrid(input)

	for i := 0; i < len(input[0]); i++ {
		index := -1
		for j := 0; j < len(input); j++ {
			if grid[j][i] == "." && index == -1 {
				index = j
			}

			if grid[j][i] == "#" {
				index = -1
			}

			if index != -1 && grid[j][i] == "O" {
				grid[index][i] = "O"
				grid[j][i] = "."
				index++
			}
		}
	}
	for i, line := range grid {
		input[i] = strings.Join(line, "")
	}

	return input
}

// Basic Idea:
//   - Repeated cycling (ie after all N, W, S and E rotation) of the grid
//     is bound to give a repeated pattern again. The iteration at which we
//     encounter the first repeated pattern is the start of the loop.
//   - We also need to count the length of the loop. This the number of iterations
//     it takes to arrive again at the same pattern.
//   - Now that we have both: the iteration of start of the loop and the loop length,
//     we can now figure out what will be the end state at the end of the total number of
//     iterations.
//
// we use a cache to save the entries after each 4 direction spin to reduce calculations.
//
// Note: an assumption made here is that the beginning of the loop is encoutered in a reasonable
// number of iterations. If this is not the case, ie if beginning of the loop is encountered after
// a significant number of iterations, this code will be too slow to arrive at a solution in a reasonable time.
func getGridAfterXCycles(input []string, times int) []string {
	cache = make(map[string]string)
	grid := aoc.Get2DGrid(input)
	cacheEntry := computeCacheEntry(grid)

	loopStartsAt := -1
	loopEntries := []string{}

	for i := 0; i < times; i++ {

		// check for cache hit, ie if current grid has already been evaluated.
		// if yes:
		// 		- if this is the first cache hit: current iteration is start of the loop
		// 		- for each cache hit after the first one, we keep storing the results in a loop slice
		// 		- if the same cache hit is encountered again, we just completed a full loop.
		if value, ok := cache[cacheEntry]; ok {
			if len(loopEntries) != 0 && value == loopEntries[0] {
				break
			}

			if loopStartsAt == -1 {
				loopStartsAt = i
			}

			loopEntries = append(loopEntries, value)
			cacheEntry = value
			continue
		}

		grid = getGridAfterCycle(grid, input)
		cache[cacheEntry] = computeCacheEntry(grid)
		cacheEntry = cache[cacheEntry]
	}

	occurance := (times - 1 - loopStartsAt) % len(loopEntries)
	cacheEntry = loopEntries[occurance]
	return aoc.SplitStringAfter(cacheEntry, len(input[0]))
}

func getGridAfterCycle(grid [][]string, input []string) [][]string {
	// north tilt
	for i := 0; i < len(input[0]); i++ {
		index := -1
		for j := 0; j < len(input); j++ {
			if grid[j][i] == "." && index == -1 {
				index = j
			}

			if grid[j][i] == "#" {
				index = -1
			}

			if index != -1 && grid[j][i] == "O" {
				grid[index][i] = "O"
				grid[j][i] = "."
				index++
			}
		}
	}

	// west tilt
	for i := 0; i < len(input); i++ {
		index := -1
		for j := 0; j < len(input[0]); j++ {
			if grid[i][j] == "." && index == -1 {
				index = j
			}

			if grid[i][j] == "#" {
				index = -1
			}

			if index != -1 && grid[i][j] == "O" {
				grid[i][index] = "O"
				grid[i][j] = "."
				index++
			}
		}
	}

	// south tilt
	for i := len(input[0]) - 1; i >= 0; i-- {
		index := -1
		for j := len(input) - 1; j >= 0; j-- {
			if grid[j][i] == "." && index == -1 {
				index = j
			}

			if grid[j][i] == "#" {
				index = -1
			}

			if index != -1 && grid[j][i] == "O" {
				grid[index][i] = "O"
				grid[j][i] = "."
				index--
			}
		}
	}

	// east tilt
	for i := len(input) - 1; i >= 0; i-- {
		index := -1
		for j := len(input[0]) - 1; j >= 0; j-- {
			if grid[i][j] == "." && index == -1 {
				index = j
			}

			if grid[i][j] == "#" {
				index = -1
			}

			if index != -1 && grid[i][j] == "O" {
				grid[i][index] = "O"
				grid[i][j] = "."
				index--
			}
		}
	}

	return grid
}

func findWeight(input []string) int {
	sum := 0
	for i, line := range input {
		for _, char := range line {
			if char == 'O' {
				sum += (len(input) - i)
			}
		}
	}
	return sum
}

func computeCacheEntry(grid [][]string) (answer string) {
	for _, line := range grid {
		answer += strings.Join(line, "")
	}
	return
}
