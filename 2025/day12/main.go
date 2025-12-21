package main

import (
	"fmt"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	shapes, trees := getShapesAndSizes(input)
	fmt.Println("answer for part 1: ", ans1(shapes, trees))
	// fmt.Println("answer for part 2: ")
}

type tree struct {
	width    int
	height   int
	presents []int
}

func getShapesAndSizes(input []string) ([][]string, []tree) {
	shapes := [][]string{}
	output := []tree{}
	currentShape := []string{}
	isShape := true
	for _, line := range input {

		if strings.Index(line, "x") != -1 {
			isShape = false
		}

		switch isShape {
		case true:
			if strings.Index(line, ":") != -1 {
				continue
			}
			if len(line) == 0 {
				shapes = append(shapes, currentShape)
				currentShape = []string{}
				continue
			}
			currentShape = append(currentShape, line)
		case false:
			output = append(output, tree{
				width:    aoc.FetchNumFromStringIgnoringNonNumeric(line[:strings.Index(line, "x")]),
				height:   aoc.FetchNumFromStringIgnoringNonNumeric(line[strings.Index(line, "x"):strings.Index(line, ":")]),
				presents: aoc.FetchSliceOfIntsInString(line[strings.Index(line, ":")+1:]),
			})
		}
	}
	return shapes, output
}

func getAreaOfEachShape(shapes [][]string) map[int]int {
	area := make(map[int]int)
	for i, s := range shapes {
		a := 0
		for _, line := range s {
			a += strings.Count(line, "#")
		}
		area[i] = a
	}
	fmt.Println(area)
	return area
}

func canTotalPresentAreaFitInSpace(input tree, area map[int]int) bool {
	totalArea := 0
	for i, p := range input.presents {
		totalArea += area[i] * p
	}
	fmt.Println(totalArea, input.height*input.width)
	return totalArea <= input.height*input.width
}

func ans1(shapes [][]string, trees []tree) int {
	area := getAreaOfEachShape(shapes)
	count := 0
	for _, t := range trees {
		if canTotalPresentAreaFitInSpace(t, area) {
			count++
		}
	}
	return count
}
