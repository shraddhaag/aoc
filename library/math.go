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

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// solveLinearEquation accepts the coeffecients of two linear equations:
//
//	a1x + b1y = c1
//	a2x + b2y = c2
//
// the function retruns two values:
//
// boolen - to indicate if the value of X and Y are whole numbers or not.
// coordinates - actual values for X and Y
func SolveLinearEquation(a, b, c Coordinates) (bool, Coordinates) {
	x := ((b.X * (-c.Y)) - (b.Y * (-c.X))) / ((a.X * b.Y) - (a.Y * b.X))
	y := (((-c.X) * a.Y) - ((-c.Y) * a.X)) / ((a.X * b.Y) - (a.Y * b.X))
	if (((b.X*(-c.Y))-(b.Y*(-c.X)))%((a.X*b.Y)-(a.Y*b.X)) == 0) &&
		((((-c.X)*a.Y)-((-c.Y)*a.X))%((a.X*b.Y)-(a.Y*b.X)) == 0) {
		return true, Coordinates{x, y}
	}
	return false, Coordinates{x, y}
}
