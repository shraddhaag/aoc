package main

import (
	"fmt"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	grid, movements, start := getGridAndMovements(input)
	expandedGrid, otherStart := expandGrid(grid)
	fmt.Println("answer for part 1: ", getFinalGrid(grid, movements, start))
	fmt.Println("answer for part 2: ", getFinalGrid(expandedGrid, movements, otherStart))
}

func getGridAndMovements(input []string) (grid [][]string, movements string, start aoc.Coordinates) {
	isMovements := false
	for index, row := range input {
		if len(row) == 0 {
			grid = aoc.Get2DGrid(input[:index+1])
			isMovements = true
		}

		if isMovements && len(row) != 0 {
			movements = strings.Join([]string{movements, row}, "")
		}

		if strings.Contains(row, "@") {
			start = aoc.Coordinates{strings.Index(row, "@"), index}
		}
	}
	return
}

func expandGrid(input [][]string) (finalGrid [][]string, start aoc.Coordinates) {
	for j, row := range input {
		finalRow := []string{}
		for i, char := range row {
			switch char {
			case ".":
				finalRow = append(finalRow, []string{".", "."}...)
			case "#":
				finalRow = append(finalRow, []string{"#", "#"}...)
			case "O":
				finalRow = append(finalRow, []string{"[", "]"}...)
			case "@":
				start.X = i * 2
				start.Y = j
				finalRow = append(finalRow, []string{"@", "."}...)
			}

		}
		finalGrid = append(finalGrid, finalRow)
	}
	return finalGrid, start
}

func processStep(start aoc.Coordinates, grid [][]string, step string) ([][]string, aoc.Coordinates) {
	var direction aoc.Coordinates
	switch step {
	case "^":
		direction = aoc.Coordinates{0, -1}
	case "<":
		direction = aoc.Coordinates{-1, 0}
	case ">":
		direction = aoc.Coordinates{1, 0}

	case "v":
		direction = aoc.Coordinates{0, 1}
	}

	newX, newY := start.X+direction.X, start.Y+direction.Y

	if !isValidStep(aoc.Coordinates{newX, newY}, grid) {
		return grid, start
	}

	switch grid[newY][newX] {
	case ".":
		grid[start.Y][start.X] = "."
		grid[newY][newX] = "@"
		return grid, aoc.Coordinates{newX, newY}
	case "#":
		return grid, start
	case "O":
		for isValidStep(aoc.Coordinates{newX, newY}, grid) {
			switch grid[newY][newX] {
			case ".":
				grid[start.Y][start.X] = "."
				grid[newY][newX] = "O"
				grid[start.Y+direction.Y][start.X+direction.X] = "@"
				return grid, aoc.Coordinates{start.X + direction.X, start.Y + direction.Y}
			case "#":
				return grid, start
			}
			newX += direction.X
			newY += direction.Y
		}
	case "[", "]":
		isPossible := moveBoxes(aoc.Coordinates{newX, newY}, direction, grid)
		if !isPossible {
			return grid, start
		}
		grid[newY][newX] = "@"
		grid[start.Y][start.X] = "."
		return grid, aoc.Coordinates{newX, newY}
	}
	return grid, start
}

func isValidStep(current aoc.Coordinates, input [][]string) bool {
	if current.X < 0 || current.Y < 0 || current.X >= len(input[0]) || current.Y >= len(input) {
		return false
	}
	return true
}

func getFinalGrid(grid [][]string, movements string, start aoc.Coordinates) int {
	for _, char := range movements {
		grid, start = processStep(start, grid, string(char))
	}

	count := 0
	for j, row := range grid {
		for i, char := range row {
			switch char {
			case "O", "[":
				count += (i) + (j)*100
			}
		}
	}
	return count
}

func moveBoxes(start, direction aoc.Coordinates, grid [][]string) bool {
	if direction.Y == 0 && (direction.X == 1 || direction.X == -1) {
		return moveBoxesHorizontally(start, direction, grid)
	}

	if direction.X == 0 && (direction.Y == 1 || direction.Y == -1) {
		return moveBoxVertically(start, direction, grid)

	}
	return false
}

// moveBoxesHorizontally does a DFS to look for the first empty slot where all the
// blocks can be shifted.
func moveBoxesHorizontally(start, direction aoc.Coordinates, grid [][]string) bool {
	newX, newY := start.X+direction.X, start.Y+direction.Y

	switch grid[newY][newX] {
	case "#":
		return false
	case ".":
		grid[newY][newX], grid[start.Y][start.X] = grid[start.Y][start.X], grid[newY][newX]
		return true
	case "]", "[":
		isPossible := moveBoxes(aoc.Coordinates{newX, newY}, direction, grid)
		if !isPossible {
			return false
		}
		grid[start.Y][start.X], grid[newY][newX] = grid[newY][newX], grid[start.Y][start.X]
	}
	return true
}

// moveBoxVertically does a BFS to check if all child nodes end with a free space
// such that boxes can be moved. If true, we move the boxes vertically in the given
// direction.
func moveBoxVertically(start, direction aoc.Coordinates, grid [][]string) bool {
	next := []aoc.Coordinates{start}
	if grid[start.Y][start.X] == "]" {
		next = append(next, aoc.Coordinates{start.X - 1, start.Y})
	} else {
		next = append(next, aoc.Coordinates{start.X + 1, start.Y})
	}

	visited := make(map[aoc.Coordinates]struct{})
	visitedSlice := []aoc.Coordinates{}

	for len(next) != 0 {
		process := next[0]
		next = next[1:]

		if _, ok := visited[process]; ok {
			continue
		}

		visited[process] = struct{}{}
		visitedSlice = append(visitedSlice, process)

		newX, newY := process.X+direction.X, process.Y+direction.Y
		switch grid[newY][newX] {
		case ".":
			continue
		case "#":
			return false
		case "]":
			next = append(next, aoc.Coordinates{newX, newY})
			next = append(next, aoc.Coordinates{newX - 1, newY})
		case "[":
			next = append(next, aoc.Coordinates{newX, newY})
			next = append(next, aoc.Coordinates{newX + 1, newY})
		}
	}

	// When traversing the grid to perfrom BFS, we also store all the cells
	// that we are processing, as these will be the same cells that will need to be moved.
	// We start shifting cells from the end of the traversed path
	// so that we first move the top-most/bottom-most cell first.
	// Each cell is shifted in the desired direction and the current cell is marked empty.
	// Moving blocks this way saves us from a lot of corner cases.
	for i := len(visitedSlice) - 1; i >= 0; i-- {
		x, y := visitedSlice[i].X+direction.X, visitedSlice[i].Y+direction.Y
		grid[y][x] = grid[visitedSlice[i].Y][visitedSlice[i].X]
		grid[visitedSlice[i].Y][visitedSlice[i].X] = "."
	}

	return true
}
