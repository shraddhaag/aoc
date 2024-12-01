package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")

	fmt.Println("answer for part 1: ", evaluateAllHands(input, processHand))
	fmt.Println("answer for part 2: ", evaluateAllHands(input, processHand2))
}

type info struct {
	amount int
	hand   string
	rank   int
}

func evaluateAllHands(input []string, processHand func(string) info) int {
	rankMap := map[int][]info{}
	handMap := map[string]info{}
	for _, line := range input {
		hand := processHand(line)
		rankMap[hand.rank] = append(rankMap[hand.rank], hand)
		handMap[hand.hand] = hand
	}

	return fetchTotalWinnings(rankMap, handMap)
}

func processHand(input string) info {

	i := info{
		hand:   transformHand(input[:5]),
		amount: aoc.FetchNumFromStringIgnoringNonNumeric(input[5:]),
	}

	tempMap := map[rune]int{}

	for _, char := range input[:5] {
		tempMap[char] += 1
	}

	i.rank = calculateTypeOfHandWhenNoJoker(tempMap)
	return i
}

func calculateTypeOfHandWhenNoJoker(countChars map[rune]int) int {
	count2 := 0
	count3 := 0
	count1 := 0

	for _, value := range countChars {
		if value == 5 {
			return 7
		}

		if value == 4 {
			return 6
		}

		if value == 3 {
			count3++
		}

		if value == 2 {
			count2++
		}

		if value == 1 {
			count1++
		}
	}

	if count3 == 1 && count2 == 1 {
		return 5
	}
	if count3 == 1 && count2 == 0 && count1 == 2 {
		return 4
	}

	if count2 == 2 && count1 == 1 {
		return 3
	}

	if count2 == 1 && count1 == 3 {
		return 2
	}

	if count1 == 5 {
		return 1
	}

	return 0
}

func processHand2(input string) info {

	i := info{
		hand:   transformHand2(input[:5]),
		amount: aoc.FetchNumFromStringIgnoringNonNumeric(input[5:]),
	}

	tempMap := map[rune]int{}
	joker := 0
	for _, char := range input[:5] {
		if char == 'J' {
			joker++
			continue
		}
		tempMap[char] += 1
	}

	if joker == 0 {
		i.rank = calculateTypeOfHandWhenNoJoker(tempMap)
		return i
	}

	count2 := 0
	count3 := 0
	count1 := 0

	for char, value := range tempMap {
		if char == 'J' {
			continue
		}
		if value == 5 {
			i.rank = 7
			return i
		}

		if value == 4 {
			i.rank = 6
		}

		if value == 3 {
			count3++
		}

		if value == 2 {
			count2++
		}

		if value == 1 {
			count1++
		}
	}
	if (i.rank == 6 && joker == 1) || (joker >= 4) {
		i.rank = 7
		return i
	}

	if count3 == 1 && joker >= 1 {
		i.rank = 5 + joker
		return i
	}

	if count2 == 2 && joker == 1 {
		i.rank = 5
		return i
	}

	if count2 == 1 && joker == 1 {
		i.rank = 4
		return i
	}

	if count2 == 1 && joker >= 2 {
		i.rank = 4 + joker
		return i
	}

	if count2 == 0 && count3 == 0 {
		if joker == 3 {
			i.rank = 6
			return i
		}

		if joker == 2 {
			i.rank = 4
			return i
		}

		if joker == 1 {
			i.rank = 2
			return i
		}
	}

	i.rank = 0
	return i
}

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, 2
// T, Q, K, J, A, 9, 8, 7, 6, 5, 4, 3, 2
func transformHand(input string) string {
	var output strings.Builder

	for _, char := range input {
		if unicode.IsDigit(char) {
			output.WriteRune(char)
		}

		if char == 'A' {
			output.WriteRune('T')
		}

		if char == 'K' {
			output.WriteRune('Q')
		}

		if char == 'Q' {
			output.WriteRune('K')
		}
		if char == 'J' {
			output.WriteRune('J')
		}
		if char == 'T' {
			output.WriteRune('A')
		}
	}
	return output.String()
}

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, 2
// T, Q, K, A, 9, 8, 7, 6, 5, 4, 3, 2, 0
func transformHand2(input string) string {
	var output strings.Builder

	for _, char := range input {
		if unicode.IsDigit(char) {
			output.WriteRune(char)
		}

		if char == 'A' {
			output.WriteRune('T')
		}

		if char == 'K' {
			output.WriteRune('Q')
		}

		if char == 'Q' {
			output.WriteRune('K')
		}
		if char == 'J' {
			output.WriteRune('0')
		}
		if char == 'T' {
			output.WriteRune('A')
		}
	}
	return output.String()
}

func fetchTotalWinnings(rankMap map[int][]info, handMap map[string]info) int {
	rank := 1
	sum := 0
	for i := 0; i <= 7; i++ {
		if _, ok := rankMap[i]; !ok {
			continue
		}

		sliceHand := []string{}
		for _, value := range rankMap[i] {
			sliceHand = append(sliceHand, value.hand)
		}

		sort.Strings(sliceHand)

		for _, j := range sliceHand {
			mul := (rank * handMap[j].amount)
			sum += mul
			rank++
		}
	}

	return sum
}
