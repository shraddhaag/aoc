package main

import (
	"fmt"
	"strings"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	steps := getSteps(input[0])

	sum := 0
	for _, line := range steps {
		sum += hashAlgo(line)
	}

	fmt.Println("answer for part 1: ", sum)
	// answer for part 2 using slices
	fmt.Println("answer for part 2: ", calculateFocusingPower(getFinalStateOfBoxes(steps)))
	// answer for part2 using linked hashmap
	fmt.Println("answer for part 2: ", calculateFocusingPower2(getFinalStateOfBoxes2(steps)))

}

func getSteps(input string) []string {
	return strings.Split(input, ",")
}

func hashAlgo(input string) int {
	current := 0
	for _, char := range input {
		current += int(char)
		current *= 17
		current %= 256
	}
	return current
}

func parseSequence(input string) (label string, symbol string, boxNum, num int) {
	symbolIndex := strings.Index(input, "-")
	symbol = "-"
	if symbolIndex == -1 {
		symbolIndex = strings.Index(input, "=")
		symbol = "="
	}
	label = input[:symbolIndex]
	boxNum = hashAlgo(label)
	if symbol == "=" {
		num = aoc.FetchNumFromStringIgnoringNonNumeric(input[symbolIndex+1:])
	}
	return
}

type lens struct {
	label string
	num   int
}

func getFinalStateOfBoxes(input []string) map[int][]lens {
	boxes := make(map[int][]lens, 256)
	for _, line := range input {
		label, symbol, boxNum, num := parseSequence(line)
		switch symbol {
		case "-":
			handleDashSymbol(boxes, boxNum, label)
		case "=":
			handleEqualSymbol(boxes, boxNum, num, label)
		default:
			panic("unhandled symbol encountered")
		}
	}
	return boxes
}

func getFinalStateOfBoxes2(input []string) map[int]*aoc.Linkedhashmap[string, int] {
	boxes := make(map[int]*aoc.Linkedhashmap[string, int], 256)
	for _, line := range input {
		label, symbol, boxNum, num := parseSequence(line)
		switch symbol {
		case "-":
			handleDashSymbol2(boxes, boxNum, label)
		case "=":
			handleEqualSymbol2(boxes, boxNum, num, label)
		default:
			panic("unhandled symbol encountered")
		}
	}

	return boxes
}

func calculateFocusingPower(boxes map[int][]lens) int {
	sum := 0
	for boxNum, content := range boxes {
		for i, l := range content {
			c := (boxNum + 1) * (i + 1) * l.num
			sum += c
		}
	}
	return sum
}

func calculateFocusingPower2(boxes map[int]*aoc.Linkedhashmap[string, int]) int {
	sum := 0
	for boxNum, content := range boxes {
		i := 0
		for _, num := range content.GetAllValues() {
			c := (boxNum + 1) * (i + 1) * num
			sum += c
			i++
		}
	}
	return sum
}

func handleDashSymbol(boxes map[int][]lens, boxNum int, label string) {
	if _, ok := boxes[boxNum]; !ok {
		return
	}
	for i, l := range boxes[boxNum] {
		if l.label == label {
			replaceWith := boxes[boxNum][:i]
			if i+1 < len(boxes[boxNum]) {
				replaceWith = append(replaceWith, boxes[boxNum][i+1:]...)
			}
			boxes[boxNum] = replaceWith
		}
	}
}

func handleEqualSymbol2(boxes map[int]*aoc.Linkedhashmap[string, int], boxNum, num int, label string) {
	if _, ok := boxes[boxNum]; !ok {
		boxes[boxNum] = &aoc.Linkedhashmap[string, int]{}
	}
	boxes[boxNum].Put(label, num)
}

func handleDashSymbol2(boxes map[int]*aoc.Linkedhashmap[string, int], boxNum int, label string) {
	if _, ok := boxes[boxNum]; !ok {
		return
	}
	boxes[boxNum].Delete(label)
}

func handleEqualSymbol(boxes map[int][]lens, boxNum, num int, label string) {
	if _, ok := boxes[boxNum]; !ok {
		boxes[boxNum] = []lens{
			{
				label: label,
				num:   num,
			},
		}
		return
	}

	for i, l := range boxes[boxNum] {
		if l.label == label {
			new := boxes[boxNum][:i]
			new = append(new, lens{label: l.label, num: num})
			if i+1 < len(boxes[boxNum]) {
				new = append(new, boxes[boxNum][i+1:]...)
			}
			boxes[boxNum] = new
			return
		}
	}

	boxes[boxNum] = append(boxes[boxNum], lens{
		label: label,
		num:   num,
	})
}
