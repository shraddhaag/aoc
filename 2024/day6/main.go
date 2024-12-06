package main

import (
	"fmt"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.Get2DGrid(aoc.ReadFileLineByLine("input.txt"))
	ans1, path := findPath(input, findStartingPoint(input))
	ans2 := findNewObstacleCount(input, findStartingPoint(input), path)

	fmt.Println("answer for part 1: ", ans1)
	fmt.Println("answer for part 2: ", ans2)
}

func findStartingPoint(input [][]string) point {
	for i, row := range input {
		for j, char := range row {
			switch char {
			case "^":
				return point{j, i, up}
			case "<":
				return point{j, i, left}

			case ">":
				return point{j, i, right}

			case "v":
				return point{j, i, down}

			}
		}
	}
	return point{-1, -1, up}
}

type point struct {
	x         int
	y         int
	direction int
}

type coordinates struct {
	x int
	y int
}

const (
	up    = 0
	down  = 1
	right = 2
	left  = 3
)

func findPath(input [][]string, start point) (int, map[coordinates]int) {
	path := make(map[coordinates]int)
	count := 0
	current := start
	for {
		if _, ok := path[coordinates{current.x, current.y}]; !ok {
			count++
			path[coordinates{current.x, current.y}] = current.direction
		}

		isValid, newCurrent := findNextStep(input, current)
		if !isValid {
			return count, path
		}

		current = newCurrent
	}
	return count, path
}

// main thing to note: we enounter a loop whenever we come
// across the same coordinates + direction.
func isLoop(input [][]string, start point) bool {
	path := make(map[point]struct{})
	path2 := make(map[coordinates]struct{})
	current := start
	for {
		// update path
		if _, ok := path[current]; !ok {
			path[current] = struct{}{}
		} else {
			return true
		}

		if _, ok := path2[coordinates{current.x, current.y}]; !ok {
			path2[coordinates{current.x, current.y}] = struct{}{}
		}

		valid, newCurrent := findNextStep(input, current)
		if !valid {
			return false
		}

		current = newCurrent
	}
	return false
}

func findNextStep(input [][]string, current point) (bool, point) {
	valid, possibleNext := getNextStepWithDirectionPreserved(input, current)
	if !valid {
		return false, possibleNext
	}

	switch input[possibleNext.y][possibleNext.x] {
	case "#":
		// this is really subtle, consider the below case
		// where > represents your current position + direction:
		// ....#.....
		// ........>#
		// ........#.
		// When you turn right and step, you again encounter
		// an obstacle (a '#'). This is a valid case and you
		// can not exit the loop at this point.
		// Instead, you turn right once MORE, an effective turn of
		// 180 degrees this time, and then continue forward.
		return findNextStep(input, turn90(input, current))
	case ".":
		return true, possibleNext
	case "^":
		return true, possibleNext
	}
	return false, possibleNext
}

// for the guard to be stuck in a loop, the new obstacle has to
// be placed on the guard's existing path (ie path figured out
// in the first part).
// obstacle placed at any other place will not change the guard's path.
func findNewObstacleCount(input [][]string, start point, path map[coordinates]int) int {
	count := 0
	obstanceMap := make(map[coordinates]struct{})
	for step, _ := range path {
		if step.x == start.x && step.y == start.y {
			continue
		}
		if input[step.y][step.x] == "." {

			input[step.y][step.x] = "#"
			if isLoop(input, start) {
				if _, ok := obstanceMap[step]; !ok {
					count++
					obstanceMap[step] = struct{}{}
				}
			}
			input[step.y][step.x] = "."
		}
	}
	return count
}

func getNextStepWithDirectionPreserved(input [][]string, current point) (bool, point) {
	switch current.direction {
	case up:
		current.y -= 1
	case down:
		current.y += 1
	case right:
		current.x += 1
	case left:
		current.x -= 1
	}
	if current.x < 0 || current.y < 0 || current.x >= len(input[0]) || current.y >= len(input) {
		return false, current
	}
	return true, current
}

func turn90(input [][]string, current point) point {
	switch current.direction {
	case up:
		current.direction = right
	case down:
		current.direction = left
	case right:
		current.direction = down
	case left:
		current.direction = up
	}
	return current
}
