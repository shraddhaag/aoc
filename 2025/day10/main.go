package main

import (
	"fmt"
	"strings"

	"github.com/draffensperger/golp"
	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	machines := parseInput(input)

	ans1, ans2 := ans(machines)
	fmt.Println("answer for part 1: ", ans1)
	fmt.Println("answer for part 2: ", ans2)
}

type machine struct {
	pattern    string
	intPattern []int
	buttons    [][]int
	joltage    []int
}

func parseInput(input []string) []machine {
	output := make([]machine, 0)
	for _, line := range input {
		m := machine{
			pattern: line[1:strings.Index(line, "]")],
			joltage: aoc.FetchSliceOfIntsInString(line[strings.Index(line, "{"):strings.Index(line, "}")]),
		}

		buttons := strings.Split(line[strings.Index(line, "]")+1:strings.Index(line, "{")+1], " ")
		for _, b := range buttons {
			button := aoc.FetchSliceOfIntsInString(b)
			if len(button) != 0 {
				m.buttons = append(m.buttons, button)
			}
		}

		intPattern := []int{}
		for _, c := range m.pattern {
			switch c {
			case '.':
				intPattern = append(intPattern, 0)
			case '#':
				intPattern = append(intPattern, 1)
			}
		}
		m.intPattern = intPattern

		output = append(output, m)
	}
	return output
}

// findLeastButtonPresses uses BFS to find the
// least number of button presses to get to the
// desired state.
//
// Some key insights which helped reach me this
// solution:
// 1. If a button is pressed once, it has a measurable
// impact on the state. If it is pressed twice, it reverts
// what its impact on the state. Extrapolating that, it
// does not matter how many times the same button is pressed.
// It only matters if the button was pressed or not.
// 2. Order of pressing buttons is irrelevant.
//
// Based on the above, my solution does the following:
// 1. Start with a single button press and see if we reach
// the desired state. If not, start pressing 2 and so on.
// 2. Exit criteria: since we are starting from the least
// number of presses, the first result we encounter will be the solution.
// 3. Since measurable difference is only made when a button is
// pressed once, when pressing 2 or more buttons, I ensure
// all the buttons are unique. (this is done while adding
// button sequences to press to queue)
// 4. Visited map uses button indexes joined together in a string
// as key, as buttons can have more than 1 switches.
func findLeastButtonPresses(m machine) int {
	// use button indexes joined together in a string as key
	// eg - pressing the 1st and 3rd button will have key: "1,3".
	visited := make(map[string]interface{})
	// use button indexes in queue. So, if 1st and 3rd button
	// needs to be presssed togthere, that entry will be []int{1,3}
	queue := make([][]int, 0)

	// we are starting off with single button presses,
	// thus appending all buttons in queue individually.
	for i := range m.buttons {
		queue = append(queue, []int{i})
	}

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		currentStringRep := aoc.ConvertIntSliceToString(current, ",")

		if _, ok := visited[currentStringRep]; ok {
			continue
		} else {
			visited[currentStringRep] = struct{}{}
		}

		buttonPresses := 0
		state := []rune{}
		for _ = range len(m.pattern) {
			state = append(state, '.')
		}

		// currentButtonMap stroes what buttons are pressed
		// in the current sequence. This is used later
		// while adding next possible sequences to the queue
		// so we only add other unique buttons to the sequences.
		currentButtonMap := make(map[int]interface{})

		// get state after the current button sequence is pressed
		for _, c := range current {
			currentButtonMap[c] = struct{}{}
			button := m.buttons[c]
			for _, i := range button {
				switch state[i] {
				case '.':
					state[i] = '#'
				case '#':
					state[i] = '.'
				}
			}
			buttonPresses++
		}

		if string(state) == m.pattern {
			// we are already iterating from the lowest
			return buttonPresses
		} else {
			// add the next possile sequences in the queue
			for i := range m.buttons {
				if _, ok := currentButtonMap[i]; ok {
					continue
				}
				queue = append(queue, append([]int{i}, current...))
			}
		}
	}
	// this is incorrect, but we need a return statement
	return -1
}

func ans(input []machine) (int, int) {
	ans1, ans2 := 0, 0
	for _, m := range input {
		ans1 += findLeastButtonPresses(m)
		ans2 += findLeastButtonPressesUsingMIP(m, false, m.joltage)
	}
	return ans1, ans2
}

// findLeastButtonPressesUsingMIP uses Mixed Integer Programming.
// This results in integer solutions to Linear Programming.
//
// How is Part 2 a Linear Programming Question?
// Lets take the third example from the test case, say we have buttons
// B1=(0,1,2,3,4), B2=(0,3,4), B3=(0,1,2,4,5) and B4=(1,2), the desired
// joltage is J={10,11,11,5,10,5}, where we need to find number of presses
// for each button.
// If we look closely, this is a system of equations. Lets supposed each
// button is pressed X1, X2, X3 and X4 times. So the problem can be presented as:
// X1*B1 + X2*B2 + X3*B3 + X4*B4 = J, where X1+X2+X3+X4 needs to be minimum
// Expanding this a lil more:
// X1 X2 X3 X4
// |1  1  1 0| = |10|
// |1  0  1 1| = |11|
// |1  0  1 1| = |11|
// |1  1  0 0| = |5 |
// |1  1  1 0| = |10|
// |0  0  1 0| = |5 |
func findLeastButtonPressesUsingMIP(m machine, isBinary bool, target []int) int {
	lp := golp.NewLP(0, len(m.buttons))
	matrix := make([][]int, 0)
	objectiveRow := []float64{}
	for _ = range len(target) {
		matrix = append(matrix, make([]int, len(m.buttons)))
	}
	for i, b := range m.buttons {
		for _, j := range b {
			matrix[j][i] = 1
		}
		objectiveRow = append(objectiveRow, 1)
		if isBinary {
			lp.SetBinary(i, true)
		} else {
			lp.SetInt(i, true)
		}
	}

	for i, row := range matrix {
		entry := make([]golp.Entry, 0)
		for j, v := range row {
			entry = append(entry, golp.Entry{j, float64(v)})
		}
		lp.AddConstraintSparse(entry, golp.EQ, float64(target[i]))
	}

	lp.SetObjFn(objectiveRow)
	lp.Solve()

	count := lp.Variables()
	ans := 0
	for _, c := range count {
		ans += int(c)
	}
	return ans

}
