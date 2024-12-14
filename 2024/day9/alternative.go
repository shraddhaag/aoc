package main

import (
	"container/heap"
	"math"

	aoc "github.com/shraddhaag/aoc/library"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type fileSegment struct {
	startIndex int
	length     int
	fileNum    int
}

func getHeapsAndFiles(input string) ([]IntHeap, []fileSegment) {
	heaps := make([]IntHeap, 9)
	for i, _ := range heaps {
		heap.Init(&heaps[i])
	}

	files := []fileSegment{}
	index := 0

	for i, char := range input {
		num := aoc.FetchNumFromStringIgnoringNonNumeric(string(char))
		if num == 0 {
			continue
		}

		switch {
		case i%2 != 0:
			heap.Push(&heaps[num-1], index)
		default:
			files = append(files, fileSegment{
				startIndex: index,
				length:     num,
				fileNum:    i / 2,
			})
		}

		index += num
	}
	return heaps, files
}

func performFileCompaction(heaps []IntHeap, files []fileSegment) []fileSegment {
	for index := len(files) - 1; index >= 0; index-- {
		file, heapIndex, emptySpaceIndex := files[index], -1, math.MaxInt

		for i := file.length; i <= 9; i++ {
			if heaps[i-1].Len() == 0 {
				continue
			}

			newEmptySpaceIndex := heap.Pop(&heaps[i-1]).(int)

			// if new empty space's index is greater than the file index, we can not use it
			// if new empty space's index is smaller than the previous empty space index, we will use that instead
			if !(newEmptySpaceIndex < file.startIndex && newEmptySpaceIndex < emptySpaceIndex) {
				heap.Push(&heaps[i-1], newEmptySpaceIndex)
				continue
			}

			if heapIndex != -1 {
				heap.Push(&heaps[heapIndex], emptySpaceIndex)
			}

			emptySpaceIndex = newEmptySpaceIndex
			heapIndex = i - 1
		}

		// we did not find an empty slot for the current file
		if heapIndex == -1 {
			continue
		}

		// move file to empty space
		files[index].startIndex = emptySpaceIndex

		// if empty space length is greater than the file length,
		// push the remaining space back in the correct heap
		if heapIndex > file.length-1 {
			newHeapIndex := heapIndex - file.length
			heap.Push(&heaps[newHeapIndex], emptySpaceIndex+file.length)
		}

	}

	return files
}

func calculateUpdatedCheckSum(files []fileSegment) int {
	sum := 0
	for _, file := range files {
		for i := range file.length {
			sum += (file.startIndex + i) * file.fileNum
		}
	}
	return sum
}
