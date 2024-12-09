package main

import (
	"fmt"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	fmt.Println("answer for part 1: ", calculateChecksum(moveFileBlocks(generateFileBlock(input[0]))))
	fmt.Println("answer for part 2: ", calculateChecksum(moveFileBlocks2(generateFileBlock(input[0]))))
}

func generateFileBlock(input string) []int {
	fileBlock := []int{}
	for index, char := range input {
		for _ = range aoc.FetchNumFromStringIgnoringNonNumeric(string(char)) {

			// even positions are files
			if index%2 == 0 {
				fileBlock = append(fileBlock, index/2)
			} else {
				fileBlock = append(fileBlock, -1)
			}
		}
	}
	return fileBlock
}

func moveFileBlocks(block []int) []int {
	start, end := 0, len(block)-1
	for start < end {
		if block[start] == -1 && block[end] != -1 {
			block[start], block[end] = block[end], block[start]
			start++
			end--
			continue
		}
		if block[start] != -1 {
			start++
		}
		if block[end] == -1 {
			end--
		}
	}
	return block
}

func writeFile(block []int, fileNum int, fileLength int, index int) {
	for i := range fileLength {
		block[index+i] = fileNum
	}
}

func clearFile(block []int, length int, index int) {
	for i := range length {
		block[index+i] = -1
	}
}

func calculateChecksum(input []int) int {
	sum := 0
	for index, fileNumber := range input {
		if fileNumber == -1 {
			continue
		}
		sum += index * fileNumber
	}
	return sum
}

func moveFile(fileBlock []int, length int, originalStartIndex int) {
	freeSpaceCount := 0
	freeSpaceStartIndex := -1
	for i := 1; i <= originalStartIndex; i++ {
		if fileBlock[i-1] != -1 && fileBlock[i] == -1 {
			if freeSpaceCount >= length {
				writeFile(fileBlock, fileBlock[originalStartIndex], length, freeSpaceStartIndex)
				clearFile(fileBlock, length, originalStartIndex)
			}
			freeSpaceStartIndex = i
			freeSpaceCount = 1
			continue
		}

		if fileBlock[i-1] == -1 && fileBlock[i] != -1 {
			if freeSpaceCount >= length {
				writeFile(fileBlock, fileBlock[originalStartIndex], length, freeSpaceStartIndex)
				clearFile(fileBlock, length, originalStartIndex)
			}
			freeSpaceStartIndex = -1
			freeSpaceCount = 0
		}

		if fileBlock[i] == -1 {
			freeSpaceCount++
		}
	}
}

func moveFileBlocks2(fileBlock []int) []int {
	currentFile := -1
	currentFileLength := 0
	for i := len(fileBlock) - 1; i > 0; i-- {
		if fileBlock[i] != -1 {
			currentFileLength += 1
			currentFile = i
			if fileBlock[i] != fileBlock[i-1] {
				if currentFileLength != 0 {
					moveFile(fileBlock, currentFileLength, currentFile)
					currentFileLength = 0
				}
			}
		} else {
			if currentFileLength != 0 {
				moveFile(fileBlock, currentFileLength, currentFile)
				currentFileLength = 0
			}
		}
	}
	return fileBlock
}
