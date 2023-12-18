package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")

	fmt.Println("answer for part 1: ", sol1(input))
	fmt.Println("answer for part 2: ", sol2(input))
}

type node struct {
	x int64
	y int64
}

func sol1(input []string) int64 {
	start := node{0, 0}
	vertices := []node{}
	prev := string(input[len(input)-1][0])
	var sum int64
	for _, line := range input {
		space := line[2 : 2+strings.Index(line[2:], " ")]
		num := int64(aoc.FetchNumFromStringIgnoringNonNumeric(space))
		newV := calculateVertices(start, int64(num), string(line[0]), prev)
		vertices = append(vertices, newV)
		start = newV
		prev = string(line[0])
		sum += num
	}
	return findArea(vertices) + (sum / 2) + 1
}

func sol2(input []string) int64 {

	dir := map[int]string{}
	dir[0] = "R"
	dir[1] = "D"
	dir[2] = "L"
	dir[3] = "U"

	start := node{0, 0}
	vertices := []node{}
	var sum int64
	prev := dir[aoc.FetchNumFromStringIgnoringNonNumeric(string(input[len(input)-1][strings.Index(input[len(input)-1], "#")+1:][5]))]
	for _, line := range input {
		real := line[strings.Index(line, "#")+1:]
		d := dir[aoc.FetchNumFromStringIgnoringNonNumeric(string(real[5]))]
		num, _ := strconv.ParseInt(real[0:5], 16, 64)
		newV := calculateVertices(start, num, d, prev)
		vertices = append(vertices, newV)
		start = newV
		prev = d
		sum += num
	}

	return findArea(vertices) + (sum / 2) + 1
}

// R - j++
// L - j--
// U - i--
// D - i++
func calculateVertices(current node, num int64, dir, prev string) node {
	final := current
	switch dir {
	case "R":
		final.y += num
	case "L":
		final.y -= num
	case "U":
		final.x -= num
	case "D":
		final.x += num
	}
	return final
}

func findArea(vertices []node) int64 {
	var sum int64

	for i := 0; i < len(vertices)-2; i++ {
		node1 := vertices[i]
		node2 := vertices[i+1]
		sum += (node1.x * node2.y) - (node1.y * node2.x)
	}

	return int64(math.Abs(float64(sum / 2)))
}
