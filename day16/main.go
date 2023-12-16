package main

import (
	"fmt"
	"math"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")

	fmt.Println("answer for part 1: ", part1(input))
	fmt.Println("answer for part 2: ", part2(input))
}

func part1(input []string) int {
	start := findStart(0, 0, input)
	return countHash(getFinalMatrixState(input, start[0]))
}

func countHash(output [][]string) int {
	count := 0
	for _, line := range output {
		for _, char := range line {
			if char == "#" {
				count++
			}
		}
	}
	return count
}

func part2(input []string) int {
	max := 0
	for i, line := range input {

		// if the current line is the first or last, shine light from each cell in the line
		// and calculate energise number.
		if i == 0 || i == len(input)-1 {
			for j, _ := range line {
				starts := findStart(i, j, input)
				for _, s := range starts {
					max = int(math.Max(float64(max), float64(countHash(getFinalMatrixState(input, s)))))
				}
			}
			continue
		}

		// if the current line is not first or last, only shine light from first and last cell of the line
		// and calculate energise number for each
		starts := findStart(i, 0, input)
		for _, s := range starts {
			max = int(math.Max(float64(max), float64(countHash(getFinalMatrixState(input, s)))))
		}

		starts = findStart(i, len(input[0])-1, input)
		for _, s := range starts {
			max = int(math.Max(float64(max), float64(countHash(getFinalMatrixState(input, s)))))
		}
	}
	return max
}

func getIntialisedOutput(input []string) [][]string {
	output := make([][]string, len(input))
	for i, _ := range output {
		output[i] = make([]string, len(input[0]))
		for j, _ := range input[0] {
			output[i][j] = "."
		}
	}
	return output
}

func getFinalMatrixState(input []string, start []step) [][]string {
	output := getIntialisedOutput(input)

	// evaluated is used to store starting points of ray of light that we have already traversed.
	// while traversing the ray path, if we encounter a location + direction
	// that has been traversed before as a starting point, we simply break.
	// this helps prevent infinite loops.
	evaluated := map[step]bool{}

	for len(start) != 0 {
		s := start[0]
		output[s.lastx][s.lasty] = "#"
		for isValid(s, input, evaluated) {
			o := evaluateSquare(input, output, s)

			if len(o) == 0 {
				break
			}

			s = o[0]

			// if more than one location can be the next step, add all others to the start queue
			// so they can be evaluated later.
			if len(o) > 1 {
				start = append(start, o[1:]...)
			}
		}
		evaluated[start[0]] = true
		start = start[1:]
	}
	return output
}

type step struct {
	x     int
	y     int
	lastx int
	lasty int
}

// isValid checks if the current coordinates are within the grid and
// ensures light from this direction has not been evaluated before.
func isValid(s step, input []string, evaluated map[step]bool) bool {
	if _, ok := evaluated[s]; ok {
		return false
	}
	if s.x < 0 || s.y < 0 || s.x >= len(input) || s.y >= len(input[0]) {
		return false
	}
	return true
}

// evaluateSquare is responsible for returning all next steps an incoming
// ray of light can take when encountered with the mirror direction at the
// specified cell. To accurately determine this, we need 3 things:
//   - coordinates of current cell
//   - coordinates of the previous cell (to determine the direction of light)
//   - symbol (ie angle of mirror) at the current cell
func evaluateSquare(input []string, output [][]string, s step) []step {
	if s.x >= len(input) || s.x < 0 || s.y < 0 || s.y >= len(input[0]) {
		return []step{}
	}

	output[s.x][s.y] = "#"

	dir := getDirection(s)

	switch input[s.x][s.y] {
	case '.':
		switch dir {
		case up:
			s.x--
			s.lastx--
		case down:
			s.x++
			s.lastx++
		case left:
			s.y--
			s.lasty--
		case right:
			s.y++
			s.lasty++
		}
		return []step{s}
	case '-':
		switch dir {
		case up, down:
			return []step{{s.x, s.y + 1, s.x, s.y}, {s.x, s.y - 1, s.x, s.y}}
		case left:
			s.y--
			s.lasty--
			return []step{s}
		case right:
			s.y++
			s.lasty++
			return []step{s}
		}
	case '|':
		switch dir {
		case up:
			s.x--
			s.lastx--
			return []step{s}
		case down:
			s.x++
			s.lastx++
			return []step{s}
		case left, right:
			return []step{{s.x + 1, s.y, s.x, s.y}, {s.x - 1, s.y, s.x, s.y}}
		}
	case '/':
		switch dir {
		case up:
			s.y++
			s.lastx--
		case down:
			s.y--
			s.lastx++
		case left:
			s.x++
			s.lasty--
		case right:
			s.x--
			s.lasty++
		}
		return []step{s}
	case 92:
		switch dir {
		case up:
			s.y--
			s.lastx--
		case down:
			s.y++
			s.lastx++
		case left:
			s.x--
			s.lasty--
		case right:
			s.x++
			s.lasty++
		}
		return []step{s}
	}

	panic("unhandled character found")
}

// findStart is responsible for finding all directions in which a ray can go
// when the given (x,y) lie at the border of the input with no previous location.
// note: the corners can have 2 possible directions for incoming ray of light.
// We should consider all possible income directions and return next step/s for each.
func findStart(i, j int, input []string) [][]step {
	output := [][]step{}
	if input[i][j] == '.' {
		if i >= 0 && i < len(input) {
			if j == 0 {
				output = append(output, []step{{i, j + 1, i, j}})
			} else if j == len(input[0])-1 {
				output = append(output, []step{{i, j - 1, i, j}})
			}
		}
		if j >= 0 && j < len(input[0]) {
			if i == 0 {
				output = append(output, []step{{i + 1, j, i, j}})
			} else if i == len(input)-1 {
				output = append(output, []step{{i - 1, j, i, j}})
			}
		}
	}

	if input[i][j] == '|' {
		if j >= 0 && j < len(input[0]) {
			if i == 0 {
				output = append(output, []step{{i + 1, j, i, j}})
			} else if i == len(input)-1 {
				output = append(output, []step{{i - 1, j, i, j}})
			}
		}
		if i >= 0 && i < len(input) {
			if j == 0 || j == len(input[0])-1 {
				output = append(output, []step{{i + 1, j, i, j}, {i - 1, j, i, j}})
			}
		}
	}

	if input[i][j] == '-' {
		if j >= 0 && j < len(input[0]) {
			if i == 0 || i == len(input[0])-1 {
				output = append(output, []step{{i, j + 1, i, j}, {i, j - 1, i, j}})
			}
		}
		if i >= 0 && i < len(input) {
			if j == 0 {
				output = append(output, []step{{i, j + 1, i, j}})
			} else if j == len(input[0])-1 {
				output = append(output, []step{{i, j - 1, i, j}})
			}
		}
	}

	if input[i][j] == 92 {
		if i >= 0 && i < len(input) {
			if j == 0 {
				output = append(output, []step{{i + 1, j, i, j}})
			} else if j == len(input[0])-1 {
				output = append(output, []step{{i - 1, j, i, j}})
			}
		}
		if j >= 0 && j < len(input[0]) {
			if i == 0 {
				output = append(output, []step{{i, j + 1, i, j}})
			} else if i == len(input)-1 {
				output = append(output, []step{{i, j - 1, i, j}})
			}
		}
	}

	if input[i][j] == '/' {
		if i >= 0 && i < len(input) {
			if j == 0 {
				output = append(output, []step{{i - 1, j, i, j}})
			} else if j == len(input[0])-1 {
				output = append(output, []step{{i + 1, j, i, j}})
			}
		}
		if j >= 0 && j < len(input[0]) {
			if i == 0 {
				output = append(output, []step{{i, j - 1, i, j}})
			} else if i == len(input)-1 {
				output = append(output, []step{{i, j + 1, i, j}})
			}
		}
	}

	return output
}

type direction int

const (
	up direction = iota
	down
	left
	right
)

func getDirection(s step) direction {
	if s.x == s.lastx && s.y < s.lasty {
		return left
	}

	if s.x == s.lastx && s.y > s.lasty {
		return right
	}

	if s.y == s.lasty && s.x > s.lastx {
		return down
	}

	if s.y == s.lasty && s.x < s.lastx {
		return up
	}

	panic("invalid step encountered")
}
