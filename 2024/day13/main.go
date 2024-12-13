package main

import (
	"fmt"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	ans1, ans2 := getTokenCount(getPrizes(input))
	fmt.Println("answer for part 2: ", ans1)
	fmt.Println("answer for part 2: ", ans2)
}

func getPrizes(input []string) [][]aoc.Coordinates {
	values := []aoc.Coordinates{}
	prizes := [][]aoc.Coordinates{}
	for _, line := range input {
		if len(line) == 0 {
			continue
		}
		nums := aoc.FetchSliceOfIntsInString(line)
		values = append(values, aoc.Coordinates{nums[0], nums[1]})

		if strings.Contains(line, "Prize") {
			prizes = append(prizes, values)
			values = []aoc.Coordinates{}
		}
	}
	return prizes
}

// getButtonPressesIfValid checks if there is a valid number
// of button presses to achieve the target.
//
// Consider A and B to be the final number of buttom presses for
// buttons A and B respectively.
//
// 1. Its only possible to acheive the target if A and B are integers.
// (We can't press a button in fraction).
//
// 2. To find values of A and B, we can get it by solving the linear
// equations (explained why below).
func getButtonPressesIfValid(a, b, final aoc.Coordinates, add int) (bool, aoc.Coordinates) {
	// Consider this:
	// On pressing button A: X moves by A1 and Y by A2
	// On pressing button B: X moves by B1 and Y by B2
	// The final location needed is (C1, C2).
	// If X and Y be the button presses for A and B respectively, then:
	// A1*X + B1*Y = C1
	// A2*X + B2*Y = C2
	// the above are just linear equations!
	// we can get the values of Y and Y easily using Crammer's Rule.
	final.X += add
	final.Y += add
	return aoc.SolveLinearEquation(a, b, final)
}

func getTokenCount(input [][]aoc.Coordinates) (int, int) {
	count1, count2 := 0, 0
	for _, prize := range input {
		isValid1, tokens1 := getButtonPressesIfValid(prize[0], prize[1], prize[2], 0)
		if isValid1 {
			count1 += tokens1.X*3 + tokens1.Y
		}

		isValid2, tokens2 := getButtonPressesIfValid(prize[0], prize[1], prize[2], 10000000000000)
		if isValid2 {
			count2 += tokens2.X*3 + tokens2.Y

		}
	}
	return count1, count2
}
