package main

import (
	"fmt"
	"sort"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	ccord := getCoordinates(input)
	distanceMap, distances := findDistancesBetweenAll(ccord)

	fmt.Println("answer for part 1: ", findCircuits(distanceMap, distances, 1000))
	fmt.Println("answer for part 2: ", findCircuits2(distanceMap, distances, ccord))
}

func getCoordinates(input []string) []aoc.Coordinates3D {
	coordinates := []aoc.Coordinates3D{}

	for _, line := range input {
		c := aoc.FetchSliceOfIntsInString(line)
		coordinates = append(coordinates, aoc.Coordinates3D{c[0], c[1], c[2]})
	}

	return coordinates
}

func distance(a, b aoc.Coordinates3D) int {
	return ((a.X - b.X) * (a.X - b.X)) + ((a.Y - b.Y) * (a.Y - b.Y)) + ((a.Z - b.Z) * (a.Z - b.Z))
}

func findDistancesBetweenAll(points []aoc.Coordinates3D) (map[int][]aoc.Coordinates3D, []int) {
	distances := []int{}
	distanceMap := map[int][]aoc.Coordinates3D{}

	for i, p1 := range points {
		for _, p2 := range points[i:] {
			if p1 == p2 {
				continue
			}
			d := distance(p1, p2)
			distanceMap[d] = []aoc.Coordinates3D{p1, p2}
			distances = append(distances, d)
		}
	}
	sort.Ints(distances)
	return distanceMap, distances
}

func findCircuits(distanceMap map[int][]aoc.Coordinates3D, distances []int, connectionCount int) int {
	circuits := [][]aoc.Coordinates3D{}
	pointsUsed := map[aoc.Coordinates3D]int{}
	count := 0
	for _, d := range distances {
		if count == connectionCount {
			break
		}

		points := distanceMap[d]
		count++
		existingCircuit1, existingCircuit2 := -1, -1
		if c, ok := pointsUsed[points[0]]; ok {
			existingCircuit1 = c
		}
		if c, ok := pointsUsed[points[1]]; ok {
			existingCircuit2 = c
		}

		// create new ciruit
		if existingCircuit1 == -1 && existingCircuit2 == -1 {
			circuits = append(circuits, points)
			pointsUsed[points[0]] = len(circuits) - 1
			pointsUsed[points[1]] = len(circuits) - 1
			continue
		}

		// both already exist in the same circuit
		if existingCircuit1 == existingCircuit2 {
			continue
		}

		// if any one point is not existing in the circuit
		if existingCircuit1 == -1 || existingCircuit2 == -1 {
			p := aoc.Coordinates3D{}
			c := -1

			if existingCircuit1 == -1 {
				p = points[0]
				c = existingCircuit2
			}
			if existingCircuit2 == -1 {
				p = points[1]
				c = existingCircuit1
			}
			pointsUsed[p] = c
			circuits[c] = append(circuits[c], p)
			continue
		}

		// if both are in different circuits
		if existingCircuit1 != existingCircuit2 {
			// lets move all of circuit 2 in 1
			circuits[existingCircuit1] = append(circuits[existingCircuit1], circuits[existingCircuit2]...)

			// fix numbers in map
			for _, p := range circuits[existingCircuit2] {
				pointsUsed[p] = existingCircuit1
			}

			circuits[existingCircuit2] = []aoc.Coordinates3D{}
		}
	}

	lenCircuits := []int{}

	for _, c := range circuits {
		lenCircuits = append(lenCircuits, len(c))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(lenCircuits)))
	return lenCircuits[0] * lenCircuits[1] * lenCircuits[2]
}

func findCircuits2(distanceMap map[int][]aoc.Coordinates3D, distances []int, allPoints []aoc.Coordinates3D) int {
	circuits := [][]aoc.Coordinates3D{}
	pointsUsed := map[aoc.Coordinates3D]int{}
	points := []aoc.Coordinates3D{}
	for _, d := range distances {

		// check if ciruits if only 1 now
		count, numPoints := 0, 0
		for _, i := range circuits {
			if len(i) == 0 {
				count++
			} else {
				numPoints = len(i)
			}
		}
		if count == len(circuits)-1 && numPoints == len(allPoints) {
			break
		}

		// fmt.Println(circuits)
		points = distanceMap[d]
		existingCircuit1, existingCircuit2 := -1, -1
		if c, ok := pointsUsed[points[0]]; ok {
			existingCircuit1 = c
		}
		if c, ok := pointsUsed[points[1]]; ok {
			existingCircuit2 = c
		}

		// create new ciruit
		if existingCircuit1 == -1 && existingCircuit2 == -1 {
			circuits = append(circuits, points)
			pointsUsed[points[0]] = len(circuits) - 1
			pointsUsed[points[1]] = len(circuits) - 1
			continue
		}

		// both already exist in the same circuit
		if existingCircuit1 == existingCircuit2 {
			continue
		}

		// if any one point is not existing in the circuit
		if existingCircuit1 == -1 || existingCircuit2 == -1 {
			p := aoc.Coordinates3D{}
			c := -1

			if existingCircuit1 == -1 {
				p = points[0]
				c = existingCircuit2
			}
			if existingCircuit2 == -1 {
				p = points[1]
				c = existingCircuit1
			}
			pointsUsed[p] = c
			circuits[c] = append(circuits[c], p)
			continue
		}

		// if both are in different circuits
		if existingCircuit1 != existingCircuit2 {
			// lets move all of circuit 2 in 1
			circuits[existingCircuit1] = append(circuits[existingCircuit1], circuits[existingCircuit2]...)

			// fix numbers in map
			for _, p := range circuits[existingCircuit2] {
				pointsUsed[p] = existingCircuit1
			}

			circuits[existingCircuit2] = []aoc.Coordinates3D{}
		}
	}

	return points[0].X * points[1].X
}
