package main

import (
	"fmt"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	connections := getConnections(input)

	fmt.Println("answer for part 1: ", getUniquePathUsingDFS("you", make(map[string]int), "out", connections))

	// there are two possible ways to go from svr to out:
	// svr -> dac -> fft -> out
	// svr -> fft -> dac -> out
	a := getUniquePathUsingDFS("svr", make(map[string]int), "dac", connections)
	b := getUniquePathUsingDFS("svr", make(map[string]int), "fft", connections)
	c := getUniquePathUsingDFS("dac", make(map[string]int), "fft", connections)
	d := getUniquePathUsingDFS("fft", make(map[string]int), "dac", connections)
	e := getUniquePathUsingDFS("fft", make(map[string]int), "out", connections)
	f := getUniquePathUsingDFS("dac", make(map[string]int), "out", connections)
	fmt.Println("answer for part 2: ", (a*c*e)+(b*d*f))
}

func getConnections(input []string) map[string][]string {
	output := make(map[string][]string)

	for _, line := range input {
		output[line[:strings.Index(line, ":")]] = strings.Fields(line[strings.Index(line, ":")+1:])
	}
	return output
}

func getUniquePathUsingDFS(currentNode string, visited map[string]int, end string, connections map[string][]string) int {
	if currentNode == end {
		return 1
	}

	if c, ok := visited[currentNode]; ok {
		return c
	}

	count := 0
	if next, ok := connections[currentNode]; ok {
		for _, node := range next {
			count += getUniquePathUsingDFS(node, visited, end, connections)
		}
	}
	visited[currentNode] = count
	return count
}
