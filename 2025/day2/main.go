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
	idRanges := getIDRanges(input[0])

	ans1, ans2 := 0, 0

	for _, r := range idRanges {
		i1, i2 := getInvalidID(r)
		ans1 += i1
		ans2 += i2
	}
	fmt.Println("answer for part 1: ", ans1)
	fmt.Println("answer for part 2: ", ans2)
}

func getIDRanges(input string) [][]int {
	idRanges := make([][]int, 0)

	ranges := strings.Split(input, ",")

	for _, r := range ranges {
		nums := strings.Split(r, "-")
		idRanges = append(idRanges, []int{aoc.FetchNumFromStringIgnoringNonNumeric(nums[0]), aoc.FetchNumFromStringIgnoringNonNumeric(nums[1])})
	}

	return idRanges
}

func getInvalidID(idRange []int) (int, int) {
	invalidIDs1 := 0
	invalidIDs2 := 0

	iterator := idRange[0]
	for iterator <= idRange[1] {
		if !isValidID(iterator) {
			invalidIDs1 += iterator
		}

		if !isValidID2(iterator) {
			invalidIDs2 += iterator
		}
		iterator++
	}
	return invalidIDs1, invalidIDs2
}

func isValidID(num int) bool {

	length := len(strconv.Itoa(num))

	if length%2 != 0 {
		return true
	}

	if (num % int(math.Pow(10, float64(length/2)))) == (num / int(math.Pow(10, float64(length/2)))) {
		return false
	}

	return true
}

/*
since all nums fit in int32, max value possible
is 2,147,483,647, ie. maximum number of digits
is 10. So the maximum length of repeating sequence
can only be 5. So we can check if sequences of
length [1,5] exist.
*/
func isValidID2(num int) bool {
	numLength := len(strconv.Itoa(num))

	switch {
	case numLength%5 == 0:
		if !isValidIDWithSeqLength(num, 5) {
			return false
		}
		fallthrough
	case numLength%4 == 0:
		if !isValidIDWithSeqLength(num, 4) {
			return false
		}
		fallthrough
	case numLength%3 == 0:
		if !isValidIDWithSeqLength(num, 3) {
			return false
		}
		fallthrough
	case numLength%2 == 0:
		if !isValidIDWithSeqLength(num, 2) {
			return false
		}
		fallthrough
	default:
		if !isValidIDWithSeqLength(num, 1) {
			return false
		}
	}
	return true
}

func isValidIDWithSeqLength(num, seqLength int) bool {
	seqs := make([]int, 0)

	for start := num; start > 0; {
		seqs = append(seqs, start%int(math.Pow(10, float64(seqLength))))
		start = start / int(math.Pow(10, float64(seqLength)))
	}

	if len(seqs) == 1 {
		return true
	}

	count := 0
	for _, s := range seqs {
		if s == seqs[0] {
			count++
		}
	}
	return count != len(seqs)
}
