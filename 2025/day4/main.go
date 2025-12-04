package main

import (
	"fmt"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	grid := aoc.Get2DGrid(input)

	ans1, _ := removeRollsOfPaper(grid)
	ans2 := getMaxPaperRollsRemoved(grid)
	fmt.Println("answer for part 1: ", ans1)
	fmt.Println("answer for part 2: ", ans2)
}

func removeRollsOfPaper(grid [][]string) (int, [][]string) {
	validPositions := 0
	newGrid := aoc.CopyGrid(grid)
	for i, row := range grid {
		for j, val := range row {
			if val != "@" {
				continue
			}

			adjacentRolls := 0

			if i > 0 && grid[i-1][j] == "@" {
				adjacentRolls++
			}
			if i > 0 && j > 0 && grid[i-1][j-1] == "@" {
				adjacentRolls++

			}
			if i > 0 && j < len(row)-1 && grid[i-1][j+1] == "@" {
				adjacentRolls++
			}
			if j > 0 && grid[i][j-1] == "@" {
				adjacentRolls++
			}
			if j < len(row)-1 && grid[i][j+1] == "@" {
				adjacentRolls++
			}
			if i < len(grid)-1 && grid[i+1][j] == "@" {
				adjacentRolls++
			}
			if i < len(grid)-1 && j > 0 && grid[i+1][j-1] == "@" {
				adjacentRolls++
			}
			if i < len(grid)-1 && j < len(row)-1 && grid[i+1][j+1] == "@" {
				adjacentRolls++
			}

			if adjacentRolls < 4 {
				validPositions++
				newGrid[i][j] = "."
			}
		}
	}
	return validPositions, newGrid

}

func getMaxPaperRollsRemoved(grid [][]string) int {
	maxRollsRemoved, rollsRemoved := 0, 0

	for {
		rollsRemoved, grid = removeRollsOfPaper(grid)
		if rollsRemoved == 0 {
			return maxRollsRemoved
		}
		maxRollsRemoved += rollsRemoved
	}
}
