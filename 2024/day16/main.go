package main

import (
	"fmt"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.Get2DGrid(aoc.ReadFileLineByLine("input.txt"))
	score, path := dijktras(input)
	fmt.Println("answer for part 1: ", score)
	fmt.Println("answer for part 2: ", getUniqueCount(input, path))
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

func dijktras(matrix [][]string) (int, map[aoc.Coordinates]int) {

	priorityQueue := pq.NewWith(func(a, b interface{}) int {
		return a.(step).score - b.(step).score
	})

	priorityQueue.Enqueue(step{aoc.Coordinates{1, len(matrix) - 2}, aoc.Right, 0, make(map[aoc.Coordinates]int)})

	visited := make(map[point]struct{})

	for !priorityQueue.Empty() {
		element, _ := priorityQueue.Dequeue()

		currentNode := element.(step)

		if _, ok := visited[point{currentNode.co, currentNode.lastDir}]; ok {
			continue
		}

		currentNode.path[currentNode.co] = currentNode.score

		if matrix[currentNode.co.Y][currentNode.co.X] == "E" {
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
		return []aoc.Coordinates{aoc.Up, aoc.Left, aoc.Right}
	case aoc.Down:
		return []aoc.Coordinates{aoc.Down, aoc.Left, aoc.Right}
	case aoc.Left:
		return []aoc.Coordinates{aoc.Up, aoc.Left, aoc.Down}
	case aoc.Right:
		return []aoc.Coordinates{aoc.Up, aoc.Down, aoc.Right}
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
		if dir != current.lastDir {
			score += 1000
		}

		possibleNext = append(possibleNext, step{
			co:      newPosition,
			lastDir: dir,
			score:   score,
			path:    copyMap(current.path),
		})
	}
	return possibleNext
}

func getUniqueCount(matrix [][]string, path map[aoc.Coordinates]int) int {

	priorityQueue := pq.NewWith(func(a, b interface{}) int {
		return a.(step).score - b.(step).score
	})

	priorityQueue.Enqueue(step{aoc.Coordinates{1, len(matrix) - 2}, aoc.Right, 0, make(map[aoc.Coordinates]int)})

	visited := make(map[point]struct{})
	newSafeCoordinates := make(map[aoc.Coordinates]struct{})

	for !priorityQueue.Empty() {
		element, _ := priorityQueue.Dequeue()

		currentNode := element.(step)

		if score, ok := path[currentNode.co]; ok && score == currentNode.score {
			for point, _ := range currentNode.path {
				if _, ok := path[point]; !ok {
					newSafeCoordinates[point] = struct{}{}
				}
			}
		}

		if _, ok := visited[point{currentNode.co, currentNode.lastDir}]; ok {
			continue
		}

		currentNode.path[currentNode.co] = currentNode.score

		if matrix[currentNode.co.Y][currentNode.co.X] == "E" {
			continue
		}

		nextSteps := getNextSteps(currentNode, matrix, visited)
		for _, n := range nextSteps {
			priorityQueue.Enqueue(n)
		}

		visited[point{currentNode.co, currentNode.lastDir}] = struct{}{}
	}
	return len(path) + len(newSafeCoordinates)
}
