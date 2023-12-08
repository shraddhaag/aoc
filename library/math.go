package aoc

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers []int) int {
	if len(integers) == 0 {
		return 0
	}

	if len(integers) == 1 {
		return integers[0]
	}

	result := integers[0] * integers[1] / GCD(integers[1], integers[0])

	for i := 2; i < len(integers); i++ {
		result = LCM([]int{result, integers[i]})
	}

	return result
}
