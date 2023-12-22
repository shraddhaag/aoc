package main

import (
	"fmt"
	"slices"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	sortedBlocks := getBlocks(input)
	getNumberOfBlocksShiftedAfterSettling(sortedBlocks)
	ans1, ans2 := findBlocksToDisintegrate(sortedBlocks)
	fmt.Println("answer for part 1: ", ans1)
	fmt.Println("answer for part 2: ", ans2)
}

type block struct {
	start point
	end   point
}

type point struct {
	x int
	y int
	z int
}

func getBlocks(input []string) []block {
	var output []block
	for _, line := range input {
		digits := aoc.FetchSliceOfIntsInString(line)
		output = append(output, block{
			start: point{digits[0], digits[1], digits[2]},
			end:   point{digits[3], digits[4], digits[5]},
		})
	}

	slices.SortFunc(output, func(a, b block) int {
		return a.start.z - b.start.z
	})
	return output
}

func findNewZCoordiantes(xyPlane [][]int, p1, p2 point) (int, int) {
	z1, z2 := 0, 0

	// if x and z coordinates are same, block elongates in y dir
	// note: this condition will also handle if all coordinates are same, ie 1 cubic meter
	if p1.x == p2.x && p1.z == p2.z {
		for i := p1.y; i <= p2.y; i++ {
			if xyPlane[p1.x][i] > z1 {
				z1 = xyPlane[p1.x][i]
			}
		}

		// max value only needs to be incremeanted by one
		// as block has only width 1 in z dir
		z1, z2 = z1+1, z1+1

		for i := p1.y; i <= p2.y; i++ {
			xyPlane[p1.x][i] = z1
		}

		// if y and z coordinates are same, block elongates in x dir
	} else if p1.y == p2.y && p1.z == p2.z {
		for i := p1.x; i <= p2.x; i++ {
			if xyPlane[i][p1.y] > z1 {
				z1 = xyPlane[i][p1.y]
			}
		}

		// max value only needs to be incremeanted by one
		// as block has only width 1 in z dir
		z1, z2 = z1+1, z1+1

		// value only needs to be updated by one
		// as block has only width 1 in z dir
		for i := p1.x; i <= p2.x; i++ {
			xyPlane[i][p1.y] = z1
		}

		// if x and y coordinates are same, block elongates in z dir
	} else if p1.x == p2.x && p1.y == p2.y {
		// as the block extends in Z dir, only a single block in
		// the count plane will need to be updated with latest Z coordinate
		z1 = xyPlane[p1.x][p1.y] + 1
		z2 = z1 + p2.z - p1.z
		xyPlane[p1.x][p1.y] = z2
	}

	return z1, z2
}

func findBlocksToDisintegrate(blocks []block) (int, int) {
	allowed := 0
	sum := 0

	for i, _ := range blocks {

		newBlocks := make([]block, len(blocks))
		copy(newBlocks, blocks)
		if i == 0 {
			newBlocks = newBlocks[1:]
		} else if i == len(newBlocks)-1 {
			newBlocks = blocks[:len(newBlocks)-1]
		} else {
			newBlocks = append(newBlocks[:i], newBlocks[i+1:]...)
		}

		changes := getNumberOfBlocksShiftedAfterSettling(newBlocks)
		if changes == 0 {
			allowed++
		}
		sum += changes

	}
	return allowed, sum
}

func getNumberOfBlocksShiftedAfterSettling(blocks []block) int {
	count := 0

	xyPlane := make([][]int, 10)
	for i, _ := range xyPlane {
		xyPlane[i] = make([]int, 10)
	}

	for i, b := range blocks {

		b.start.z, b.end.z = findNewZCoordiantes(xyPlane, b.start, b.end)

		if b.start != blocks[i].start || b.end != blocks[i].end {
			blocks[i] = block{b.start, b.end}
			count++
		}
	}
	return count
}
