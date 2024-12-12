package main

import (
	"fmt"
	"unicode"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.Get2DGrid(aoc.ReadFileLineByLine("input.txt"))
	ans1, ans2 := countAntinode(input, findAllAntennas(input))
	fmt.Println("answer for part 1: ", ans1)
	fmt.Println("answer for part 2: ", ans2)
}

type c struct {
	x, y int
}

func findAllAntennas(input [][]string) map[string][]c {
	antennas := make(map[string][]c)
	for j, row := range input {
		for i, char := range row {
			if unicode.IsDigit(rune(char[0])) || unicode.IsLetter(rune(char[0])) {
				if _, ok := antennas[char]; !ok {
					antennas[char] = []c{{i, j}}
				} else {
					antennas[char] = append(antennas[char], c{i, j})
				}
			}
		}
	}
	return antennas
}

func countAntinode(input [][]string, antennas map[string][]c) (int, int) {
	maxX, maxY := len(input[0]), len(input)
	count1, count2 := 0, 0
	uniqueAntiNodes1 := make(map[c]struct{})
	uniqueAntiNodes2 := make(map[c]struct{})
	for _, antenna := range antennas {
		if len(antenna) == 1 {
			continue
		}
		for i := 0; i < len(antenna)-1; i++ {
			for j := i + 1; j < len(antenna); j++ {
				validAntinodes := calculateValidAntinode(antenna[i], antenna[j], maxX, maxY, input)
				for _, antinode := range validAntinodes {
					if _, ok := uniqueAntiNodes1[antinode]; !ok {
						uniqueAntiNodes1[antinode] = struct{}{}
						count1++
					}
				}

				validAntinodes = calculateValidAntinode2(antenna[i], antenna[j], maxX, maxY, input)
				for _, antinode := range validAntinodes {
					if _, ok := uniqueAntiNodes2[antinode]; !ok {
						uniqueAntiNodes2[antinode] = struct{}{}
						count2++
					}
				}
			}
		}
	}

	return count1, count2
}

func calculateValidAntinode(a1, a2 c, maxX, maxY int, input [][]string) []c {
	validAntinode := []c{}
	diffX, diffY := a2.x-a1.x, a2.y-a1.y

	c1 := c{a1.x - diffX, a1.y - diffY}
	c2 := c{a2.x + diffX, a2.y + diffY}

	if isValidPosition(c1, maxX, maxY) {
		validAntinode = append(validAntinode, c1)
	}

	if isValidPosition(c2, maxX, maxY) {
		validAntinode = append(validAntinode, c2)
	}

	return validAntinode
}

func calculateValidAntinode2(a1, a2 c, maxX, maxY int, input [][]string) []c {
	validAntinode := []c{a1, a2}
	diffX, diffY := a2.x-a1.x, a2.y-a1.y

	c1, c2 := a1, a2

	for isValidPosition(c1, maxX, maxY) || isValidPosition(c2, maxX, maxY) {
		if isValidPosition(c1, maxX, maxY) {
			validAntinode = append(validAntinode, c1)
		}

		if isValidPosition(c2, maxX, maxY) {
			validAntinode = append(validAntinode, c2)
		}

		c1 = c{c1.x - diffX, c1.y - diffY}
		c2 = c{c2.x + diffX, c2.y + diffY}
	}

	return validAntinode
}

func isValidPosition(c1 c, maxX, maxY int) bool {

	if !(c1.x >= maxX || c1.x < 0) && !(c1.y >= maxY || c1.y < 0) {
		return true
	}
	return false
}
