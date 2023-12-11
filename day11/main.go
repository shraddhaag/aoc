package main

import (
	"fmt"
	"math"
	"regexp"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	fmt.Println("answer for part 1: ", ans(input, 2))
	fmt.Println("answer for part 2: ", ans(input, 1000000))
}

func ans(input []string, times int) int {
	h, v, p := expandGalaxy(input)
	return findGalaxies(h, v, p, times)
}

func expandGalaxy(input []string) (map[int]interface{}, map[int]interface{}, []position) {
	hashtag := regexp.MustCompile("#")

	vertical := map[int]interface{}{}
	horizontal := map[int]interface{}{}
	hashes := []position{}

	for i := 0; i < len(input[0]); i++ {
		matches := hashtag.FindAllStringIndex(input[i], -1)
		if len(matches) == 0 {
			horizontal[i] = struct{}{}
		}

		for _, m := range matches {
			hashes = append(hashes, position{
				x: i,
				y: m[0],
			})
		}

		countY := 0
		for j := 0; j < len(input); j++ {
			if input[j][i] == '#' {
				countY++
			}
		}
		if countY == 0 {
			vertical[i] = struct{}{}
		}
	}

	return horizontal, vertical, hashes

}

type position struct {
	x int
	y int
}

func findGalaxies(horizontal, vertical map[int]interface{}, hashes []position, times int) int {

	modifiedHashes := []position{}

	for _, hash := range hashes {
		x1, y1 := hash.x, hash.y
		for x, _ := range horizontal {
			if x < hash.x {
				x1 += times - 1
			}
		}

		for y, _ := range vertical {
			if y < hash.y {
				y1 += times - 1
			}
		}
		modifiedHashes = append(modifiedHashes, position{x1, y1})
	}

	count := 0
	for _, h1 := range modifiedHashes {
		for _, h2 := range modifiedHashes {
			if h1 == h2 {
				continue
			}
			x := int(math.Abs(float64(h2.x - h1.x)))
			y := int(math.Abs(float64(h2.y - h1.y)))
			count += x + y
		}
	}
	return int(count / 2)
}
