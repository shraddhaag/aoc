package main

import (
	"fmt"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.Get2DIntGrid(aoc.ReadFileLineByLine("input.txt"))
	ans1, ans2 := findTrailHeads(input)
	fmt.Println("answer for part 1: ", ans1)
	fmt.Println("answer for part 2: ", ans2)
}

type point struct {
	x, y int
}

func findNext(input [][]int, current point) []point {
	validNextSteps := []point{}
	// can check left
	if current.x > 0 && input[current.y][current.x-1] == input[current.y][current.x]+1 {
		validNextSteps = append(validNextSteps, point{current.x - 1, current.y})
	}

	// can check right
	if current.x < len(input[0])-1 && input[current.y][current.x+1] == input[current.y][current.x]+1 {
		validNextSteps = append(validNextSteps, point{current.x + 1, current.y})
	}

	// can check up
	if current.y > 0 && input[current.y-1][current.x] == input[current.y][current.x]+1 {
		validNextSteps = append(validNextSteps, point{current.x, current.y - 1})
	}

	// can check down
	if current.y < len(input)-1 && input[current.y+1][current.x] == input[current.y][current.x]+1 {
		validNextSteps = append(validNextSteps, point{current.x, current.y + 1})
	}
	return validNextSteps
}

func findScore(input [][]int, start point, trailHeads map[point]struct{}, count int) (map[point]struct{}, int) {
	if input[start.y][start.x] == 9 {
		if _, ok := trailHeads[start]; !ok {
			trailHeads[start] = struct{}{}
		}
		return trailHeads, count + 1
	}
	nextSteps := findNext(input, start)
	if len(nextSteps) == 0 {
		return trailHeads, count
	}

	for _, step := range nextSteps {
		trailHeads, count = findScore(input, step, trailHeads, count)
	}
	return trailHeads, count
}

func findTrailHeads(input [][]int) (int, int) {
	countScore := 0
	countRating := 0
	for j, row := range input {
		for i, char := range row {
			if char == 0 {
				score, rating := findScore(input, point{i, j}, make(map[point]struct{}), 0)
				countScore += len(score)
				countRating += rating
			}
		}
	}
	return countScore, countRating
}
