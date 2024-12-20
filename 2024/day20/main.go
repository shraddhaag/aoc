package main

import (
	"fmt"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.Get2DGrid(aoc.ReadFileLineByLine("input.txt"))
	start := findStart(input)
	_, path := dijktras(start, input)

	fmt.Println("answer to part 1: ", ans1(path, input, 100))
	fmt.Println("answer to part 2: ", ans2(path, input, 100))
}

func findStart(grid [][]string) aoc.Coordinates {
	for j, _ := range grid {
		for i, char := range grid[j] {
			if char == "S" {
				return aoc.Coordinates{i, j}
			}
		}
	}
	return aoc.Coordinates{}
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

func dijktras(start aoc.Coordinates, matrix [][]string) (int, map[aoc.Coordinates]int) {

	priorityQueue := pq.NewWith(func(a, b interface{}) int {
		return a.(step).score - b.(step).score
	})

	priorityQueue.Enqueue(step{start, aoc.Up, 0, make(map[aoc.Coordinates]int)})

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

		possibleNext = append(possibleNext, step{
			co:      newPosition,
			lastDir: dir,
			score:   score,
			path:    copyMap(current.path),
		})
	}
	return possibleNext
}

type cheat struct {
	s aoc.Coordinates
	e aoc.Coordinates
}

func ans1(path map[aoc.Coordinates]int, grid [][]string, leastGain int) int {
	dirs := []aoc.Coordinates{aoc.Up, aoc.Down, aoc.Right, aoc.Left}
	uniqueCheats := make(map[cheat]int)
	count := 0

	for current, currentPathCount := range path {
		for _, dir1 := range dirs {
			cheat1 := aoc.Coordinates{current.X + dir1.X, current.Y + dir1.Y}
			if !isValidStep(cheat1, grid) {
				continue
			}
			if grid[cheat1.Y][cheat1.X] != "#" {
				continue
			}

			for _, dir2 := range dirs {
				cheat2 := aoc.Coordinates{cheat1.X + dir2.X, cheat1.Y + dir2.Y}
				if !isValidStep(cheat2, grid) {
					continue
				}
				if val2, ok := path[cheat2]; ok && val2-currentPathCount-2 > 0 {
					uniqueCheats[cheat{s: cheat1, e: cheat2}] = val2 - currentPathCount - 2
					if val2-currentPathCount-2 >= leastGain {
						count++
					}
				}
			}
		}
	}
	return count
}

func ans2(path map[aoc.Coordinates]int, grid [][]string, leastGain int) int {
	uniqueCheats := make(map[cheat]int)
	count := 0

	for p1, d1 := range path {
		for p2, d2 := range path {
			if d2-d1-aoc.ManhattanDistance(p1, p2) >= leastGain && aoc.ManhattanDistance(p1, p2) <= 20 {
				uniqueCheats[cheat{s: p1, e: p2}] = d2 - d1
				count++
			}
		}
	}
	return count
}
