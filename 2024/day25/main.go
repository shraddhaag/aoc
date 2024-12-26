package main

import (
	"fmt"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	keys, locks := getKeysAndLocks(input)
	fmt.Println(checkAllLockAndKeys(locks, keys))
}

func getKeysAndLocks(input []string) (keys [][]int, locks [][]int) {
	current := []string{}

	for _, line := range input {
		if len(line) == 0 && len(current) != 0 {
			pattern, isKey := processPattern(current)
			if isKey {
				keys = append(keys, pattern)
			} else {
				locks = append(locks, pattern)
			}

			current = []string{}
			continue
		}
		current = append(current, line)
	}
	pattern, isKey := processPattern(current)
	if isKey {
		keys = append(keys, pattern)
	} else {
		locks = append(locks, pattern)
	}
	return
}

func processPattern(input []string) (heights []int, isKey bool) {
	switch input[0][0] {
	case '.':
		// this is a key
		isKey = true
		for i := 0; i < len(input[0]); i++ {
			for j := 1; j < len(input); j++ {
				if input[j][i] == '#' {
					heights = append(heights, j-1)
					break
				}
			}
		}
	case '#':
		// this is a lock
		for i := 0; i < len(input[0]); i++ {
			for j := 1; j < len(input); j++ {
				if input[j][i] == '.' {
					heights = append(heights, j-1)
					break
				}
			}
		}
	}
	// fmt.Println(input, heights)
	return
}

func isLockAndKeyFit(lock, key []int) bool {
	for index, lockHeight := range lock {
		if key[index] < lockHeight {
			return false
		}
	}
	return true
}

func checkAllLockAndKeys(locks, keys [][]int) int {
	count := 0
	for _, lock := range locks {
		for _, key := range keys {
			if isLockAndKeyFit(lock, key) {
				count++
			}
		}
	}
	return count
}
