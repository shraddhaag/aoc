package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")

	var sum int
	for _, line := range input {
		count, _ := findWinPointAndCount(line)
		sum += count
	}
	fmt.Println("answer for part 1: ", sum)
	fmt.Println("answer for part 2: ", findNumberOfCards(input))
}

func findWinPointAndCount(input string) (int, int) {
	winMap := make(map[int64]bool)
	var build strings.Builder
	isWin := true
	winPoints := 0
	winCount := 0
	for _, char := range input[strings.Index(input, ":"):] {
		if unicode.IsDigit(char) {
			build.WriteRune(char)
		}

		if char == '|' {
			isWin = false
		}

		if char == ' ' && build.Len() != 0 {
			winPoints, winCount, build = updateWinPointAndCount(build, winMap, winPoints, winCount, isWin)
		}
	}

	if build.Len() != 0 {
		winPoints, winCount, build = updateWinPointAndCount(build, winMap, winPoints, winCount, isWin)
	}

	return winPoints, winCount
}

func updateWinPointAndCount(build strings.Builder, winMap map[int64]bool, winPoints, winCount int, isWinNum bool) (int, int, strings.Builder) {
	localNum, err := strconv.ParseInt(build.String(), 10, 12)
	if err != nil {
		panic(err)
	}

	if isWinNum {
		winMap[localNum] = true
		build.Reset()
		return winPoints, winCount, build
	}

	if _, ok := winMap[localNum]; ok {
		winCount++
		if winPoints == 0 {
			winPoints = 1
		} else {
			winPoints *= 2
		}
	}
	build.Reset()
	return winPoints, winCount, build
}

func findNumberOfCards(input []string) int {
	num := make(map[int]int)
	var count int

	for i, line := range input {
		cardNumber := i + 1
		// find how many times the current card is processed
		var times int
		if _, ok := num[cardNumber]; ok {
			times = num[cardNumber]
		} else {
			num[cardNumber] = 1
			times = 1
		}

		// add current number of cards in total count
		count += times

		// find number of winning numbers in the current card
		_, point := findWinPointAndCount(line)
		if point == 0 {
			continue
		}

		// add number of times in successive cards
		for j := 1; j <= point; j++ {
			if _, ok := num[cardNumber+j]; ok {
				num[cardNumber+j] += times
			} else {
				num[cardNumber+j] = times + 1
			}
		}
	}

	return count
}
