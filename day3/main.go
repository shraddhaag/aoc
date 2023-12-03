package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readFileLineByLine("input.txt")
	fmt.Println("answer for part 1: ", getEngineNumbers(input))
	fmt.Println("answer for part 2: ", getSlowEngineNumbers(input))
}

func readFileLineByLine(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var output []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}

type engineNum struct {
	start int
	end   int
	num   int64
}

func fetchNumsInInput(input []string) [][]engineNum {
	numsInInput := [][]engineNum{}
	for _, line := range input {
		numsInLine := []engineNum{}
		var build strings.Builder
		var localNum engineNum
		var err error
		for j, char := range line {
			if isDigit(byte(char)) {
				build.WriteRune(char)
			}

			// if current character is not a digit and the previous character is a digit
			// then it is a valid engine number and add it in sum
			if !isDigit(byte(char)) && build.Len() != 0 {
				// get the num
				numS := build.String()
				localNum.num, err = strconv.ParseInt(numS, 10, 12)
				if err != nil {
					panic(err)
				}

				localNum.end = j - 1
				localNum.start = localNum.end - build.Len() + 1
				build.Reset()
				numsInLine = append(numsInLine, localNum)
				localNum = engineNum{}
			}
		}

		if build.Len() != 0 {
			numS := build.String()
			localNum.num, err = strconv.ParseInt(numS, 10, 64)
			if err != nil {
				panic(err)
			}

			localNum.end = len(line) - 1
			localNum.start = localNum.end - build.Len() + 1

			numsInLine = append(numsInLine, localNum)
		}

		numsInInput = append(numsInInput, numsInLine)
	}
	return numsInInput
}

func getEngineNumbers(input []string) int64 {
	var sum int64

	for i, value := range fetchNumsInInput(input) {
		for _, localNum := range value {
			// if char before number is symbol
			if localNum.start != 0 && input[i][localNum.start-1] != '.' && !isDigit(input[i][localNum.start-1]) {
				sum += localNum.num
				continue
			}

			// if char after number is a symbol
			if localNum.end < len(input[i])-1 && input[i][localNum.end+1] != '.' && !isDigit(input[i][localNum.end+1]) {
				sum += localNum.num
				continue
			}

			// if char in the line above or below have a symbol
			k := 0
			if localNum.start != 0 {
				k = localNum.start - 1
			}
			end := localNum.end
			if localNum.end < len(input[i])-1 {
				end += 1
			}
			for ; k <= end; k++ {
				if i != 0 {
					if input[i-1][k] != '.' && !isDigit(input[i-1][k]) {
						sum += localNum.num
						break
					}
				}

				if i != len(input)-1 {
					if input[i+1][k] != '.' && !isDigit(input[i+1][k]) {
						sum += localNum.num
						break
					}
				}
			}
		}
	}
	return sum
}

func getSlowEngineNumbers(input []string) int64 {
	var sum int64
	numsInInput := fetchNumsInInput(input)

	for i, line := range input {
		for j, char := range line {
			var count int
			var mul int64
			if char == '*' {

				// has a digit before in the same line
				if j != 0 && isDigit(line[j-1]) {
					count++

					for _, localNum := range numsInInput[i] {
						if localNum.start <= j-1 && localNum.end >= j-1 {
							if mul == 0 {
								mul = localNum.num
							} else {
								mul *= localNum.num
							}
							break
						}
					}
				}

				// has a digit after in the same line
				if j != len(line)-1 && isDigit(line[j+1]) {
					count++

					for _, localNum := range numsInInput[i] {
						if localNum.start <= j+1 && localNum.end >= j+1 {

							if mul == 0 {
								mul = localNum.num
							} else {
								mul *= localNum.num
							}
							break
						}
					}
				}

				// has number/s in the line above or below
				start := 0
				if j != 0 {
					start = j - 1
				}
				end := j
				if j < len(line)-1 {
					end += 1
				}

				// if line above has digit/s
				if i != 0 {

					for _, localNum := range numsInInput[i-1] {
						// this is the tricky bit. Nums can exist in three different cases:
						// note: '(' denotes beginning digit of number and ')' denotes the last digit
						// note: '|' denote the start and beginning of the range where digit will be considered adjacent.
						// 1. ----(-------)---- ie a single num spans the entire 3 'adjacent' indices
						//    ------|-*-|------
						// 2. ----(---)-------- ie a num end in one of the allowed indices.
						//    ------|-*-|------
						// 3. ---------(----)-- ie a num begins in one of the allowed indices.
						//    ------|-*-|------
						// considering cases 2 & 3, we can have 2 numbers adjacent to a '*' in the lines above or below, eg:
						// ---123-456----
						// ------*-------

						// case 1
						if (localNum.start <= start && localNum.end >= end) ||
							// case 2
							(localNum.start <= start && localNum.end <= end && localNum.end >= start) ||
							// case 3
							(localNum.start >= start && localNum.start <= end && localNum.end >= end) {
							count++

							if mul == 0 {
								mul = localNum.num
							} else {
								mul *= localNum.num
							}

						}
					}

				}

				// if line below has digit/s
				if i != len(input)-1 {

					// same logic as above applied here, this time for lines below '*'
					for _, localNum := range numsInInput[i+1] {
						if (localNum.start <= start && localNum.end >= end) || (localNum.start <= start && localNum.end <= end && localNum.end >= start) || (localNum.start >= start && localNum.start <= end && localNum.end >= end) {
							count++
							if mul == 0 {
								mul = localNum.num
							} else {
								mul *= localNum.num
							}

						}
					}
				}

			}

			if count == 2 {
				sum += mul
			}
		}
	}
	return sum
}

func isDigit(char byte) bool {
	if char >= 48 && char <= 57 {
		return true
	}
	return false
}
