package main

import (
	"fmt"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.Get2DGrid(aoc.ReadFileLineByLine("input.txt"))
	ans1, ans2 := alternativeSolution(input)
	fmt.Println("answer for part 1: ", ans1)
	fmt.Println("answer for part 2: ", ans2)
}

type p struct {
	x, y int
}

func checkAll4(input [][]string, current p) []p {
	sameAround := []p{}

	// can check left
	if current.x > 0 && input[current.y][current.x-1] == input[current.y][current.x] {
		sameAround = append(sameAround, p{current.x - 1, current.y})
	}

	// can check right
	if current.x < len(input[0])-1 && input[current.y][current.x+1] == input[current.y][current.x] {
		sameAround = append(sameAround, p{current.x + 1, current.y})
	}

	// can check up
	if current.y > 0 && input[current.y-1][current.x] == input[current.y][current.x] {
		sameAround = append(sameAround, p{current.x, current.y - 1})
	}

	// can check down
	if current.y < len(input)-1 && input[current.y+1][current.x] == input[current.y][current.x] {
		sameAround = append(sameAround, p{current.x, current.y + 1})
	}
	return sameAround
}

func ans(input [][]string) (int, int) {
	cost, cost2 := 0, 0
	visited := make(map[p]struct{})
	for j, row := range input {
		for i, _ := range row {
			if _, ok := visited[p{i, j}]; ok {
				continue
			}
			shape := findAllGardensRecursive(input, p{i, j}, polynomial{}, visited)
			cost += shape.area * shape.perimeter
			cost2 += shape.area * shape.sides
		}
	}
	return cost, cost2
}

type polynomial struct {
	area      int
	perimeter int
	sides     int
}

func findAllGardensRecursive(input [][]string, current p, shape polynomial, visited map[p]struct{}) polynomial {
	if _, ok := visited[current]; ok {
		return shape
	}

	checkNext := checkAll4(input, current)

	// none surrounding are same garden
	if len(checkNext) == 0 {
		if shape.area == 0 {
			shape.area = 1
			shape.perimeter = 4
			visited[current] = struct{}{}
			shape.sides = checkCorners(input, current)
			return shape
		}
		return shape
	}

	shape.perimeter += 4 - len(checkNext)
	shape.area += 1
	visited[current] = struct{}{}
	shape.sides += checkCorners(input, current)

	for _, next := range checkNext {
		shape = findAllGardensRecursive(input, next, shape, visited)
	}
	return shape
}

// checkCorners checks if there are any corners on the
// current coordinates. It checks for both, outside
// and inside corners
func checkCorners(input [][]string, current p) int {
	count := 0
	gardenType := input[current.y][current.x]
	x, y := current.x, current.y

	if x == 0 && y == 0 {
		count += 1
	}

	if x == 0 && y == len(input)-1 {
		count += 1
	}

	if x == len(input[0])-1 && y == len(input)-1 {
		count += 1
	}

	if x == len(input[0])-1 && y == 0 {
		count += 1
	}

	// top left outside corner
	// ##   __   |#
	// #O   #O   |O
	if (x > 0 && y > 0 && input[y][x-1] != gardenType && input[y-1][x] != gardenType) ||
		(x > 0 && y == 0 && input[y][x-1] != gardenType) || (x == 0 && y > 0 && input[y-1][x] != gardenType) {
		count += 1
	}

	// top left inside corner
	// OO
	// O#
	if x < len(input[0])-1 && y < len(input)-1 && input[y][x+1] == gardenType && input[y+1][x] == gardenType && input[y+1][x+1] != gardenType {
		count += 1
	}

	// top right outside corner
	// ##   __    #|
	// O#   O#    O|
	if (x < len(input[0])-1 && y > 0 && input[y][x+1] != gardenType && input[y-1][x] != gardenType) ||
		(x < len(input[0])-1 && y == 0 && input[y][x+1] != gardenType) || (x == len(input[0])-1 && y > 0 && input[y-1][x] != gardenType) {
		count += 1
	}

	// top right inside corner
	// OO
	// #O
	if x > 0 && y < len(input)-1 && input[y][x-1] == gardenType && input[y+1][x] == gardenType && input[y+1][x-1] != gardenType {
		count += 1
	}

	// bottom left outside corner
	// #O   #O    |O
	// ##   --    |#
	if (x > 0 && y < len(input)-1 && input[y][x-1] != gardenType && input[y+1][x] != gardenType) ||
		(x > 0 && y == len(input)-1 && input[y][x-1] != gardenType) || (x == 0 && y < len(input)-1 && input[y+1][x] != gardenType) {
		count += 1
	}

	// bottom left inside corner
	// O#
	// OO
	if x < len(input[0])-1 && y > 0 && input[y][x+1] == gardenType && input[y-1][x] == gardenType && input[y-1][x+1] != gardenType {
		count += 1
	}

	// bottom right outside corner
	// O#   O#    O|
	// ##   --    #|
	if (x < len(input[0])-1 && y < len(input)-1 && input[y][x+1] != gardenType && input[y+1][x] != gardenType) ||
		(x < len(input[0])-1 && y == len(input)-1 && input[y][x+1] != gardenType) || (x == len(input[0])-1 && y < len(input)-1 && input[y+1][x] != gardenType) {
		count += 1
	}

	// bottom right inside corner
	// #O
	// OO
	if x > 0 && y > 0 && input[y][x-1] == gardenType && input[y-1][x] == gardenType && input[y-1][x-1] != gardenType {
		count += 1
	}
	return count
}

func alternativeSolution(input [][]string) (int, int) {
	cost1, cost2 := 0, 0
	visitedCoordinates := make(map[p]struct{})

	for j, _ := range input {
		for i, _ := range input[j] {
			if _, ok := visitedCoordinates[p{i, j}]; !ok {
				next := []p{{i, j}}
				shape := polynomial{}
				for len(next) != 0 {
					newShape, traverseNext := findAllGardensNonRecursively(input, next[0], shape, visitedCoordinates)
					shape = newShape
					next = append(next, traverseNext...)
					next = next[1:]
				}
				cost1 += shape.area * shape.perimeter
				cost2 += shape.area * shape.sides
			}
		}
	}
	return cost1, cost2
}

func findAllGardensNonRecursively(input [][]string, current p, shape polynomial, visited map[p]struct{}) (polynomial, []p) {
	if _, ok := visited[current]; ok {
		return shape, []p{}
	}

	checkNext := checkAll4(input, current)

	// none surrounding are same garden
	if len(checkNext) == 0 {
		if shape.area == 0 {
			visited[current] = struct{}{}
			shape = polynomial{
				area: 1, perimeter: 4, sides: 4,
			}
			return shape, []p{}
		}
		return shape, []p{}
	}

	shape.perimeter += 4 - len(checkNext)
	shape.area += 1
	visited[current] = struct{}{}
	shape.sides += checkCorners(input, current)

	return shape, checkNext
}
