package main

import (
	"fmt"
	"strings"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	grid := createGrid(input, 71, 71, 1024)
	count, path := dijktras(grid)
	fmt.Println("answer to part 1: ", count)
	fmt.Println("answer to part 2: ", getCount(input, 1024, 71, 71, path))
}

func createGrid(input []string, maxX, maxY int, linesToProcess int) [][]string {
	grid := make([][]string, maxY)
	for index, _ := range grid {
		resultString := strings.Repeat(".", maxX)

		grid[index] = strings.Split(resultString, "")
	}

	for _, row := range input[:linesToProcess] {
		nums := aoc.FetchSliceOfIntsInString(row)
		grid[nums[1]][nums[0]] = "#"
	}

	return grid
}

func dijktras(matrix [][]string) (int, map[aoc.Coordinates]int) {

	priorityQueue := pq.NewWith(func(a, b interface{}) int {
		return a.(step).score - b.(step).score
	})

	priorityQueue.Enqueue(step{aoc.Coordinates{0, 0}, aoc.Right, 0, make(map[aoc.Coordinates]int)})

	visited := make(map[point]struct{})

	for !priorityQueue.Empty() {
		element, _ := priorityQueue.Dequeue()

		currentNode := element.(step)

		if _, ok := visited[point{currentNode.co, currentNode.lastDir}]; ok {
			continue
		}

		currentNode.path[currentNode.co] = currentNode.score

		if currentNode.co.Y == len(matrix)-1 && currentNode.co.X == len(matrix[0])-1 {
			return currentNode.score, currentNode.path
		}

		nextSteps := getNextSteps(currentNode, matrix, visited)
		for _, n := range nextSteps {
			priorityQueue.Enqueue(n)
		}

		visited[point{currentNode.co, currentNode.lastDir}] = struct{}{}
	}
	return -1, make(map[aoc.Coordinates]int)
}

func isValidStep(current aoc.Coordinates, input [][]string) bool {
	if current.X < 0 || current.Y < 0 || current.X >= len(input[0]) || current.Y >= len(input) {
		return false
	}
	return true
}

func copyMap(path map[aoc.Coordinates]int) map[aoc.Coordinates]int {
	new := make(map[aoc.Coordinates]int, len(path))
	for key, value := range path {
		new[key] = value
	}
	return new
}

func getAllowedDirections(direction aoc.Coordinates) []aoc.Coordinates {
	switch direction {
	case aoc.Up:
		return []aoc.Coordinates{aoc.Up, aoc.Left, aoc.Right, aoc.Down}
	case aoc.Down:
		return []aoc.Coordinates{aoc.Down, aoc.Left, aoc.Right, aoc.Up}
	case aoc.Left:
		return []aoc.Coordinates{aoc.Up, aoc.Left, aoc.Down, aoc.Right}
	case aoc.Right:
		return []aoc.Coordinates{aoc.Up, aoc.Down, aoc.Right, aoc.Left}
	}
	return []aoc.Coordinates{}
}

func getNextSteps(current step, grid [][]string, visited map[point]struct{}) []step {
	possibleNext := []step{}
	for _, dir := range getAllowedDirections(current.lastDir) {
		newPosition := aoc.Coordinates{current.co.X + dir.X, current.co.Y + dir.Y}

		if !isValidStep(newPosition, grid) {
			continue
		}

		if grid[newPosition.Y][newPosition.X] == "#" {
			continue
		}

		if _, ok := visited[point{newPosition, dir}]; ok {
			continue
		}

		score := current.score + 1

		possibleNext = append(possibleNext, step{
			co:      newPosition,
			lastDir: dir,
			score:   score,
			path:    copyMap(current.path),
		})
	}
	return possibleNext
}

type step struct {
	co      aoc.Coordinates
	lastDir aoc.Coordinates
	score   int
	path    map[aoc.Coordinates]int
}

type point struct {
	co      aoc.Coordinates
	lastDir aoc.Coordinates
}

func getCount(input []string, start, maxX, maxY int, originalPath map[aoc.Coordinates]int) string {
	for i := start; i < len(input); i++ {

		// if the new block does not land on the previous path,
		// it will not block the path and hence, a valid path exists.
		newObstacleCo := aoc.Coordinates{aoc.FetchSliceOfIntsInString(input[i-1])[0], aoc.FetchSliceOfIntsInString(input[i-1])[1]}
		if _, ok := originalPath[newObstacleCo]; !ok {
			continue
		}

		grid := createGrid(input, maxX, maxY, i)
		count, path := dijktras(grid)
		if count == -1 {
			return input[i-1]
		}
		originalPath = path
	}
	return "none"
}
