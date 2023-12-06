package main

import (
	"fmt"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input2.txt")

	fmt.Println("answer for part 1: ", ans1(input))
	fmt.Println("answer for part 2: ", ans2(input))
}

func ans1(input []string) int {
	answer := 1
	time := aoc.FetchSliceOfIntsInString(input[0])
	distance := aoc.FetchSliceOfIntsInString(input[1])

	for i, t := range time {
		answer *= findNoOfWaysToWin(t, distance[i])
	}
	return answer
}

func ans2(input []string) int {
	time := aoc.FetchNumFromStringIgnoringNonNumeric(input[0])
	distance := aoc.FetchNumFromStringIgnoringNonNumeric(input[1])

	return findNoOfWaysToWin(time, int(distance))
}

// this is the core of the problem.
// first we count number of ways we will NOT win when starting from hold time 0.
// in each iterations: check if current distance is greater than given distance:
//   - greater: break loop
//   - else: increase counter, increase hold time and continue
//
// we will get the same number of count if we were to start counting from hold time T.
//
// this means:
//   - total number of ways to not win: count * 2
//   - total number of ways: time + 1 (because we start from 0 instead of 1)
//
// so: total number of ways to win = time +1 - (count * 2)
func findNoOfWaysToWin(time, distance int) int {
	var count int
	var hold int

	for ; hold <= time; hold++ {
		dist := hold * (time - hold)
		if dist <= distance {
			count++
		} else {
			break
		}
	}
	return time - (count * 2) + 1
}
