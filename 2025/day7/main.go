package main

import (
	"fmt"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	grid := aoc.Get2DGrid(input)
	startingPoint := getStartingPoint(grid)

	fmt.Println("answer for part 1: ", getTotalBeamCount(grid, startingPoint, make(map[aoc.Coordinates]interface{})))
	fmt.Println("answer for part 2: ", getTotalUniquePaths(grid, startingPoint, make(map[aoc.Coordinates]int)))
}

func getStartingPoint(grid [][]string) aoc.Coordinates {
	for i := range grid {
		// fmt.Println(grid[i])
		for j := range grid[i] {
			if grid[i][j] == "S" {
				grid[i][j] = "."
				return aoc.Coordinates{i, j}
			}
		}
	}
	return aoc.Coordinates{0, 0}
}

func getTotalBeamCount(grid [][]string, curr aoc.Coordinates, processed map[aoc.Coordinates]interface{}) int {

	if _, ok := processed[curr]; ok {
		fmt.Println("already processed", curr)
		return 0
	}
	processed[curr] = struct{}{}

	for grid[curr.X][curr.Y] == "." {
		curr.X++

		if curr.X >= len(grid) {
			return 0
		}
	}

	count := 0
	next := []aoc.Coordinates{}
	if curr.Y > 0 {
		if _, ok := processed[aoc.Coordinates{curr.X, curr.Y - 1}]; !ok {
			next = append(next, aoc.Coordinates{curr.X, curr.Y - 1})
		}
	}
	if curr.Y < len(grid[0])-1 {
		if _, ok := processed[aoc.Coordinates{curr.X, curr.Y + 1}]; !ok {
			next = append(next, aoc.Coordinates{curr.X, curr.Y + 1})
		}
	}

	if len(next) != 0 {
		count++
	}

	for _, i := range next {
		count += getTotalBeamCount(grid, i, processed)
	}
	return count
}

func getTotalUniquePaths(grid [][]string, curr aoc.Coordinates, processed map[aoc.Coordinates]int) int {
	if c, ok := processed[curr]; ok {
		return c
	}

	for grid[curr.X][curr.Y] == "." {
		curr.X++

		if curr.X >= len(grid) {
			return 1
		}
	}

	count := 0
	next := []aoc.Coordinates{}
	if curr.Y > 0 {
		next = append(next, aoc.Coordinates{curr.X, curr.Y - 1})
	}
	if curr.Y < len(grid[0])-1 {
		next = append(next, aoc.Coordinates{curr.X, curr.Y + 1})
	}

	for _, i := range next {
		c := getTotalUniquePaths(grid, i, processed)
		processed[i] = c
		count += c
	}
	processed[curr] = count
	return count
}
