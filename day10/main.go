package main

import (
	"fmt"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	history, count := findCount(input)
	numberOfInsideElements, visual := findArea(history, input)
	fmt.Println("answer for part 1: ", count)
	fmt.Println("answer for part 2: ", numberOfInsideElements)

	for i := 0; i < len(input); i++ {
		if _, ok := visual[i]; ok {
			fmt.Println(visual[i].v, visual[i].count)
		}
	}
}

func findS(input []string) (int, int) {
	for i, line := range input {
		for j, char := range line {
			if char == 'S' {
				return i, j
			}
		}
	}
	return 0, 0
}

type position struct {
	x int
	y int
}

func findCount(input []string) ([]position, int) {
	history := []position{}

	x, y := findS(input)
	history = append(history, position{
		x: x,
		y: y,
	})

	x, y = getStartingPosition(x, y, input)

	for input[x][y] != 'S' {
		lastP := history[len(history)-1]
		x1, y1 := findNext(x, y, input, lastP.x, lastP.y)
		history = append(history, position{
			x: x,
			y: y,
		})
		x, y = x1, y1
	}
	return history, len(history) / 2
}

func getStartingPosition(x, y int, input []string) (int, int) {

	if y != (len(input[x])-1) && (input[x][y+1] == '-' || input[x][y+1] == 'J' || input[x][y+1] == '7') {
		y++
		return x, y
	}

	if x != (len(input)-1) && (input[x+1][y] == '|' || input[x+1][y] == 'J' || input[x+1][y] == 'L') {
		x++
		return x, y
	}

	if y != 0 && (input[x][y-1] == 'F' || input[x][y-1] == '-' || input[x][y-1] == 'L') {
		y--
		return x, y
	}

	if x != 0 && (input[x-1][y] == '|' || input[x-1][y] == 'L' || input[x-1][y] == '7') {
		x--
		return x, y
	}
	return x, y
}

func findNext(x, y int, input []string, lastx, lasty int) (int, int) {
	if input[x][y] == '-' {
		if y != 0 && lasty == (y-1) {
			y++
			return x, y
		}

		if y != (len(input[x])-1) && lasty == (y+1) {
			y--
			return x, y
		}
	}

	if input[x][y] == 'J' {
		if x != 0 && y != 0 && lastx == (x-1) {
			y--
			return x, y
		}

		if x != 0 && y != 0 && lasty == (y-1) {
			x--
			return x, y
		}
	}

	if input[x][y] == '|' {
		if x != 0 && x != len(input)-1 && lastx == (x-1) {
			x++
			return x, y
		}

		if x != 0 && x != len(input)-1 && lastx == (x+1) {
			x--
			return x, y
		}
	}

	if input[x][y] == 'L' {
		if x != 0 && y != len(input[x])-1 && lastx == (x-1) {
			y++
			return x, y
		}

		if x != 0 && y != len(input[x])-1 && lasty == (y+1) {
			x--
			return x, y
		}
	}

	if input[x][y] == '7' {
		if y != 0 && x != (len(input)-1) && lasty == (y-1) {
			x++
			return x, y
		}

		if x != (len(input)-1) && y != 0 && lastx == (x+1) {
			y--
			return x, y
		}
	}

	if input[x][y] == 'F' {
		if x != (len(input)-1) && y != (len(input[x])-1) && lasty == (y+1) {
			x++
			return x, y
		}

		if x != (len(input)-1) && y != (len(input[x])-1) && lastx == (x+1) {
			y++
			return x, y
		}
	}

	return x, y
}

type zoo struct {
	v     string
	count int
}

func findArea(path []position, input []string) (int, map[int]zoo) {
	replaceWith := map[string]string{}
	replaceWith["J"] = "┘"
	replaceWith["L"] = "└"
	replaceWith["7"] = "┐"
	replaceWith["F"] = "┌"
	replaceWith["|"] = "│"
	replaceWith["-"] = "─"

	mapPosition := make(map[int][]int)
	resultMap := make(map[int]zoo)
	sum := 0

	for _, p := range path {
		mapPosition[p.x] = append(mapPosition[p.x], p.y)
	}

	for k, v := range mapPosition {
		a := strings.Split(input[k], "")
		for _, j := range v {
			if a[j] != "S" {
				a[j] = replaceWith[a[j]]
			} else {
				a[j] = replaceWith[string(replaceSWith(path[0], path[1], path[len(path)-1], input))]
			}
		}
		// clean edges
		left, right := false, false
		for i, j := range a {
			if i == (len(a) - 1 - i) {
				break
			}
			r := a[len(a)-i-1]
			if j == "┘" || j == "└" || j == "┐" || j == "┌" || j == "│" || j == "─" {
				left = true
			}
			if r == "┘" || r == "└" || r == "┐" || r == "┌" || r == "│" || r == "─" {
				right = true
			}

			if !left {
				a[i] = " "
			}
			if !right {
				a[len(a)-1-i] = " "
			}
		}
		count := 0
		isInside := 0
		last := "-"
		// this is the crux of the solution:
		// we count the number of | encountered in pipe path when traversing each  cleaned input line
		// if it is odd, that means the character encountered is inside the pipe path, else outside
		// corner cases:
		// 		- (F is followed by a J) or (L is followed by a 7) then its equivalent to a single |
		// 		- '-' does not affect count
		// 		- J, L, 7, F do not affect count individually, unless 1st case is encountered
		for i, char := range a {
			if char == "│" {
				isInside++
				continue
			}
			if char == "─" {
				continue
			}

			if last == "-" && (char == "┘" || char == "└" || char == "┐" || char == "┌") {
				if char == "┘" || char == "└" || char == "┐" || char == "┌" {
					last = char
					continue
				}
			} else if last != "-" && (char == "┘" || char == "└" || char == "┐" || char == "┌") {
				if last == "└" && char == "┐" {
					isInside++
				}

				if last == "┌" && char == "┘" {
					isInside++
				}
				last = "-"
				continue
			}

			if isInside%2 == 0 {
				a[i] = " "
			}

			if isInside%2 != 0 {
				a[i] = "\033[0;32m█\033[0m"
				count++
			}
		}
		resultMap[k] = zoo{
			v:     strings.Join(a, ""),
			count: count,
		}
		sum += count
	}

	return sum, resultMap
}

func replaceSWith(s, first, last position, input []string) byte {

	for _, char := range "-|JL7F" {
		duplicateInput := input
		a := strings.Split(duplicateInput[s.x], "")
		a[s.y] = string(char)
		duplicateInput[s.x] = strings.Join(a, "")
		x, y := findNext(s.x, s.y, duplicateInput, last.x, last.y)
		if x == first.x && y == first.y {
			return byte(char)
		}
	}

	return 'S'
}
