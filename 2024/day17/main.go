package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("test1.txt")
	output := doCalc(processInput(input))

	strSlice := make([]string, len(output))
	for i, v := range output {
		strSlice[i] = strconv.Itoa(int(v))
	}
	ans1 := strings.Join(strSlice, ",")
	fmt.Println("answer for part 1: ", ans1)
	fmt.Println("answer for part 2: ", findValueOfA(processInput(input)))
}

func getOperandValue(input int64, registerA, registerB, registerC int64) int64 {
	switch input {
	case 0, 1, 2, 3:
		return input
	case 4:
		return registerA
	case 5:
		return registerB
	case 6:
		return registerC
	case 7:
		return 7
	}
	return input
}

func processOpcode(operand int64, opcode int64, registerA, registerB, registerC int64, ip int64) (int64, int64, int64, int64, int64) {
	operandCalculated := getOperandValue(operand, registerA, registerB, registerC)
	var output int64
	output = -1

	switch opcode {
	case 0:
		registerA = registerA / int64(math.Pow(2, float64(operandCalculated)))
	case 1:
		registerB = registerB ^ operandCalculated
	case 2:
		registerB = (operandCalculated % 8) % 1000
	case 3:
		if registerA != 0 {
			ip = operand
		}
	case 4:
		registerB = registerB ^ registerC
	case 5:
		output = int64(operandCalculated % 8)
	case 6:
		registerB = registerA / int64(math.Pow(2, float64(operandCalculated)))
	case 7:
		registerC = registerA / int64(math.Pow(2, float64(operandCalculated)))
	}
	return registerA, registerB, registerC, ip, output
}

func processInput(input []string) (int64, int64, int64, []int64) {
	var a, b, c int64
	inst := []int64{}
	for i, row := range input {
		nums := aoc.FetchSliceOfIntsInString(row)
		switch i {
		case 0:
			a = int64(nums[0])
		case 1:
			b = int64(nums[0])
		case 2:
			c = int64(nums[0])
		case 4:
			for _, n := range nums {
				inst = append(inst, int64(n))
			}
		}
	}
	return a, b, c, inst
}

func doCalc(a, b, c int64, inst []int64) []int64 {
	var ip int64
	var output []int64
	for ip < int64(len(inst)-1) {
		var newIP, o int64
		newIP = ip
		o = -1
		a, b, c, newIP, o = processOpcode(inst[ip+1], inst[ip+0], a, b, c, ip)
		if newIP == ip {
			ip += 2
		} else {
			ip = newIP
		}
		if o != -1 {
			output = append(output, o)
		}
	}
	return output
}

func findValueOfA(a, b, c int64, inst []int64) int64 {
	output := []int64{}
	a = 1
	continueLoop := true

	for continueLoop {
		output = doCalc(a, b, c, inst)

		if reflect.DeepEqual(output, inst) {
			return a
		}

		if len(inst) > len(output) {
			a *= 2
			continue
		}

		if len(inst) == len(output) {
			for j := len(inst) - 1; j >= 0; j-- {
				if inst[j] != output[j] {
					// Key Insight: every nth digit increments at every 8^n th step.
					// https://www.reddit.com/r/adventofcode/comments/1hg38ah/comment/m2gkd6m/
					a += int64(math.Pow(8, float64(j)))
					break
				}
			}
		}

		if len(inst) < len(output) {
			a /= 2
		}
	}
	return a
}
