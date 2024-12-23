package main

import (
	"fmt"
	"math"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	fmt.Println("answer for part 1: ", getScore(getRobots(input), 101, 103))
	fmt.Println("answer for part 1: ", findFrameWithLeastEntropy(getRobots(input), 101, 103))
}

type robot struct {
	position aoc.Coordinates
	velocity aoc.Coordinates
}

func getRobots(input []string) []robot {
	robots := []robot{}
	for _, line := range input {
		nums := aoc.FetchSliceOfIntsInString(line)
		robots = append(robots, robot{
			position: aoc.Coordinates{nums[0], -nums[1]},
			velocity: aoc.Coordinates{nums[2], nums[3]},
		})
	}
	return robots
}

func getPositionAfterCountSeconds(r robot, count int, maxX, maxY int) aoc.Coordinates {
	x, y := r.position.X, r.position.Y

	a := (x + count*r.velocity.X) % maxX
	b := (y - count*r.velocity.Y) % maxY

	if a >= maxX {
		a = a - maxX
	}

	if a < 0 {
		a = maxX + a
	}

	if b <= -maxY {
		b = b + maxY
	}

	if b > 0 {
		b = -(maxY - b)
	}
	return aoc.Coordinates{a, b}
}

func getScore(r []robot, maxX, maxY int) int {
	quad := []int{0, 0, 0, 0}
	for _, ro := range r {
		newPosition := getPositionAfterCountSeconds(ro, 100, maxX, maxY)
		x, y := newPosition.X, newPosition.Y
		switch {
		case x < maxX/2 && y > -maxY/2:
			quad[0] += 1
		case x > maxX/2 && y > -maxY/2:
			quad[1] += 1
		case x < maxX/2 && y < -maxY/2:
			quad[2] += 1
		case x > maxX/2 && y < -maxY/2:
			quad[3] += 1
		}
	}
	return quad[0] * quad[1] * quad[2] * quad[3]
}

func writeGridAfterXSecond(r []robot, maxX, maxY int, count int) {
	grid := make([][]string, maxY)
	for i := 0; i < maxY; i++ {
		grid[i] = strings.Split(strings.Repeat(".", maxX), "")
	}
	for _, ro := range r {
		newLoc := getPositionAfterCountSeconds(ro, count, maxX, maxY)
		grid[-newLoc.Y][newLoc.X] = "#"
	}
	printGrid(grid, maxX, maxY)
	return
}

func printGrid(grid [][]string, maxX, maxY int) {
	for i := 0; i < maxY; i++ {
		aoc.WriteToFile("output.txt", fmt.Sprintf(strings.Join(grid[i], "")+"\n"))
	}
}

func printAllPossibleGridsToFile(r []robot, maxX, maxY int) {
	count := 0
	for count < 103*101 {
		writeGridAfterXSecond(r, maxX, maxY, count)
		aoc.WriteToFile("output.txt", fmt.Sprintf("output after %d seconds\n", count))
		count++
	}
}

func getGridAfterXSecond(r []robot, maxX, maxY int, count int) [][]int {
	grid := make([][]int, maxY)
	for i := 0; i < maxY; i++ {
		grid[i] = make([]int, maxX)
	}
	for _, ro := range r {
		newLoc := getPositionAfterCountSeconds(ro, count, maxX, maxY)
		grid[-newLoc.Y][newLoc.X] += 1
	}
	return grid
}

func findFrameWithLeastEntropy(r []robot, maxX, maxY int) int {
	count := 0
	entropy := math.MaxFloat64
	leastEntropySecond := 0
	for count < 103*101 {
		currentEntropy := aoc.Entropy(getGridAfterXSecond(r, maxX, maxY, count))
		if currentEntropy < entropy {
			entropy = currentEntropy
			leastEntropySecond = count
		}
		count++
	}
	writeGridAfterXSecond(r, maxX, maxY, leastEntropySecond)
	return leastEntropySecond
}
