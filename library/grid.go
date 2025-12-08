package aoc

import "math"

type Coordinates struct {
	X, Y int
}

type Coordinates3D struct {
	X, Y, Z int
}

var (
	Up    = Coordinates{0, -1}
	Down  = Coordinates{0, 1}
	Left  = Coordinates{-1, 0}
	Right = Coordinates{1, 0}
)

// Entropy calculates the entropy of a 2D grid
// using Shannon's Entropy.
// https://en.wikipedia.org/wiki/Entropy_(information_theory)
func Entropy(img [][]int) float64 {
	// Flatten the image into a single slice
	var flatImg []int
	for _, row := range img {
		flatImg = append(flatImg, row...)
	}

	// Calculate histogram with 256 bins
	bins := 256
	histogram := make([]float64, bins)
	for _, value := range flatImg {
		if value >= 0 && value < bins {
			histogram[value]++
		}
	}

	// Normalize histogram
	totalPixels := float64(len(flatImg))
	for i := range histogram {
		histogram[i] /= totalPixels
	}

	// Filter out zero probabilities
	var marg []float64
	for _, prob := range histogram {
		if prob > 0 {
			marg = append(marg, prob)
		}
	}

	// Calculate entropy
	entropy := 0.0
	for _, p := range marg {
		entropy -= p * math.Log2(p)
	}

	return entropy
}

func CopyGrid(input [][]string) [][]string {
	newGrid := make([][]string, 0)

	for _, row := range input {
		newRow := make([]string, len(row))
		copy(newRow, row)
		newGrid = append(newGrid, newRow)
	}
	return newGrid
}
