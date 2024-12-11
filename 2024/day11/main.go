package main

import (
	"fmt"
	"strconv"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.FetchSliceOfIntsInString(aoc.ReadFileLineByLine("input.txt")[0])
	fmt.Println("answer for part 1: ", getStoneCountAfterBlinking(input, 25))
	fmt.Println("answer for part 2: ", getStoneCountAfterBlinking(input, 75))
}

func getStonesAfterBlink(stone int) []int {
	resultStones := []int{}

	switch {
	case stone == 0:
		resultStones = append(resultStones, 1)
	case len(strconv.Itoa(stone))%2 == 0:
		s1, s2 := splitStone(stone)
		resultStones = append(resultStones, s1, s2)
	default:
		resultStones = append(resultStones, stone*2024)
	}

	return resultStones
}

func splitStone(stone int) (int, int) {
	stoneString := strconv.Itoa(stone)
	stone1, stone2 := stoneString[:len(stoneString)/2], stoneString[len(stoneString)/2:]
	return aoc.FetchNumFromStringIgnoringNonNumeric(stone1), aoc.FetchNumFromStringIgnoringNonNumeric(stone2)
}

func getCountAfterBlinks(stone int, cache map[int][]int, blinkCount int) int {
	if _, ok := cache[stone]; ok {
		if cache[stone][blinkCount-1] != 0 {
			return cache[stone][blinkCount-1]
		}
	} else {
		cache[stone] = make([]int, 75)
	}

	if blinkCount == 1 {
		cache[stone][blinkCount-1] = len(getStonesAfterBlink(stone))
		return len(getStonesAfterBlink(stone))
	}

	sum := 0

	for _, stone := range getStonesAfterBlink(stone) {
		sum += getCountAfterBlinks(stone, cache, blinkCount-1)
	}

	cache[stone][blinkCount-1] = sum
	return sum
}

func getStoneCountAfterBlinking(input []int, timesBlink int) int {
	sum := 0
	cache := make(map[int][]int)
	for _, stone := range input {
		sum += getCountAfterBlinks(stone, cache, timesBlink)
	}
	return sum
}
