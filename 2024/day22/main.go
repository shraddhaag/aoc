package main

import (
	"fmt"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")

	ans1, ans2 := calculateSecretNums(input)
	fmt.Println("answer to part 1: ", ans1)
	fmt.Println("answer to part 1: ", ans2)
}

func findSecretNumber(input int64) int64 {
	input = prune(mix(int64(64)*input, input))
	input = prune(mix(input/32, input))
	input = prune(mix(input*2048, input))
	return input
}

func mix(value, secretNum int64) int64 {
	return value ^ secretNum
}

func prune(value int64) int64 {
	return value % 16777216
}

type bananaPrice struct {
	num    int
	change int
}

func getNumAndChange(input int64, previous int) bananaPrice {
	return bananaPrice{num: int(input % 10), change: int(input%10) - previous}
}

func calculateSecretNums(input []string) (int64, int) {
	var sum int64
	b := make([][]bananaPrice, len(input))
	for i, line := range input {
		b[i] = make([]bananaPrice, 2000)
		num := int64(aoc.FetchNumFromStringIgnoringNonNumeric(line))
		prev := int(num % 10)
		for j := range 2000 {
			num = findSecretNumber(num)
			b[i][j] = getNumAndChange(num, prev)
			prev = int(num % 10)
		}
		sum += num
	}

	return sum, findMaxNumberOfBananas(b)
}

type seq struct {
	a, b, c, d int
}

func findMaxNumberOfBananas(b [][]bananaPrice) int {
	seqMap := make(map[seq][]int)
	for i, _ := range b {
		s := []int{b[i][0].change, b[i][1].change, b[i][2].change}
		for j := 3; j < len(b[i]); j++ {
			s = append(s, b[i][j].change)

			if _, ok := seqMap[seq{s[0], s[1], s[2], s[3]}]; !ok {
				seqMap[seq{s[0], s[1], s[2], s[3]}] = make([]int, len(b))
			}

			if seqMap[seq{s[0], s[1], s[2], s[3]}][i] == 0 {
				seqMap[seq{s[0], s[1], s[2], s[3]}][i] = b[i][j].num
			}

			s = s[1:]
		}
	}
	max := 0
	for _, n := range seqMap {
		sum := 0
		for _, r := range n {
			sum += r
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
