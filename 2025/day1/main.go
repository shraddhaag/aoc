package main

import (
	"fmt"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")

	ans1, ans2 := getPassword(input)
	fmt.Println("answer for part 1: ", ans1)
	fmt.Println("answer for part 2: ", ans2)
}

func getPassword(input []string) (int, int) {
	password := 0
	zeroCross := 0
	currentSeq := 50
	isZero := 0

	for _, seq := range input {
		currentSeq, isZero = getNumber(currentSeq, seq)
		if currentSeq == 0 {
			password++
		}
		zeroCross += isZero
	}
	return password, password + zeroCross
}

func getNumber(initial int, seq string) (int, int) {
	distance := aoc.FetchNumFromStringIgnoringNonNumeric(seq)
	currentSeq, zeroCrossed := rotate(string(seq[0]), distance%100, initial)
	return currentSeq, zeroCrossed + distance/100
}

func rotate(direction string, distance, initial int) (int, int) {
	final := 0
	zeroCrossed := 0
	switch direction {
	case "L":
		if distance > initial {
			final = 100 - (distance - initial)
			if initial != 0 {
				zeroCrossed++
			}
		} else {
			final = initial - distance
		}
	case "R":
		if distance+initial > 99 {
			final = (distance + initial) - 100

			if final != 0 {
				zeroCrossed++
			}
		} else {
			final = distance + initial
		}
	}
	return final, zeroCrossed
}
