package aoc

type Coordinates struct {
	X, Y int
}

var (
	Up    = Coordinates{0, -1}
	Down  = Coordinates{0, 1}
	Left  = Coordinates{-1, 0}
	Right = Coordinates{1, 0}
)
