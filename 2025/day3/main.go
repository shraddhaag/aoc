package main

import (
	"fmt"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")

	fmt.Println("answer for part 1: ", getTotalOutputVoltage(getBatteryBanks(input), 2))
	fmt.Println("answer for part 2: ", getTotalOutputVoltage(getBatteryBanks(input), 12))
}

func getBatteryBanks(input []string) [][]int {
	batteryBank := make([][]int, 0)

	for _, b := range input {
		a := strings.Split(b, "")
		bank := []int{}
		for _, c := range a {
			bank = append(bank, aoc.FetchNumFromStringIgnoringNonNumeric(c))
		}
		batteryBank = append(batteryBank, bank)
	}
	return batteryBank
}

func getTotalOutputVoltage(input [][]int, digits int) int {
	totalVoltage := 0

	for _, i := range input {
		totalVoltage += getLargestJoltageUsingGreedy(i, digits)
	}
	return totalVoltage
}

// func getLargestJoltage(bank []int) int {
// 	largestJoltage := 0

// 	if len(bank) >= 2 {
// 		largestJoltage = (bank[0] * 10) + bank[1]
// 	}

// 	for _, j := range bank[2:] {
// 		// take last LSB as MSB
// 		possibility1 := j + (largestJoltage%10)*10
// 		possibility2 := (largestJoltage/10)*10 + j
// 		largestJoltage = int(math.Max(float64(largestJoltage), math.Max(float64(possibility1), float64(possibility2))))
// 		// fmt.Println(possibility1, possibility2, largestJoltage)
// 	}
// 	return largestJoltage
// }

func getLargestJoltageUsingGreedy(bank []int, digits int) int {
	ans := []int{}

	for len(ans) < digits {
		if len(bank)+len(ans) == digits {
			ans = append(ans, bank...)
			continue
		}

		if len(bank)+len(ans) < digits {
			return -1
		}

		highest := -1
		index := -1
		for i, num := range bank[:len(bank)-digits+len(ans)+1] {
			if num > highest {
				highest = num
				index = i
			}

			if highest == 9 {
				break
			}
		}

		ans = append(ans, highest)
		bank = bank[index+1:]
	}

	return aoc.FetchNumFromStringIgnoringNonNumeric(aoc.ConvertIntSliceToString(ans, ""))
}
