package main

import (
	"fmt"
	"strconv"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	matrix := createMatrix(input)

	fmt.Println("answer for part 1: ", dijktras(matrix, 0, 3))
	fmt.Println("answer for part 2: ", dijktras(matrix, 4, 10))
}

func createMatrix(input []string) [][]int {
	output := make([][]int, len(input))
	for i, line := range input {
		output[i] = make([]int, len(line))
		for j, char := range line {
			num, _ := strconv.ParseInt(string(char), 10, 64)
			output[i][j] = int(num)
		}
	}
	return output
}

type step struct {
	x       int
	y       int
	lastDir direction
	count   int
}

type queueNode struct {
	step
	heatLoss int
}

type direction string

const (
	up    direction = "up"
	down  direction = "down"
	left  direction = "left"
	right direction = "right"
)

func dijktras(matrix [][]int, minX, maxX int) int {

	// priority queue is used to store the (state, heatloss) for future
	// possible direction
	// it starts with the first two possible directions enqueued and as
	// the least cost node is dequeued, we enqueue all possible states
	// that can be reached from the current node.
	priorityQueue := pq.NewWith(func(a, b interface{}) int {
		return a.(queueNode).heatLoss - b.(queueNode).heatLoss
	})

	// enqueue two possibility of starting points, going left and going down
	// note: this is really subtle and not made explicity clear in the problem
	// statement, but the starting point (0,0) does not count towards the direction count.
	// this means, if we start at (0,0) going right, we can travel upto (0, max), not (0, max-1).
	priorityQueue.Enqueue(queueNode{step{0, 1, right, 1}, matrix[0][1]})
	priorityQueue.Enqueue(queueNode{step{1, 0, down, 1}, matrix[1][0]})

	// visited is a cache used to prevent loops of states
	visited := make(map[step]int)

	for !priorityQueue.Empty() {
		element, _ := priorityQueue.Dequeue()

		// priority queue always gives the least heatloss node present in the queue
		currentNode := element.(queueNode)

		if _, ok := visited[currentNode.step]; ok {
			continue
		}

		// if the final node is reached, exit
		// the first occurance is the answer because we are using a priority queue
		// and it will dequeue the least heatloss state first.
		if currentNode.x == len(matrix)-1 && currentNode.y == len(matrix[0])-1 {
			if currentNode.count < minX {
				continue
			}
			return currentNode.heatLoss
		}

		// find all possible traversals from the current node and insert them in the queue
		dirs := fetchPossibleDirections(currentNode, matrix, minX, maxX)
		possibleNextNodes := getNodesFromDirection(dirs, currentNode, matrix)
		for _, p := range possibleNextNodes {
			if _, ok := visited[p.step]; ok {
				continue
			}
			priorityQueue.Enqueue(p)
		}

		visited[currentNode.step] = currentNode.heatLoss
	}

	return -1
}

type point struct {
	x int
	y int
}

var positionMap = map[direction]point{
	up:    {-1, 0},
	down:  {1, 0},
	left:  {0, -1},
	right: {0, 1},
}

func isValid(p point, input [][]int) bool {
	if p.x < 0 || p.x > len(input)-1 || p.y < 0 || p.y > len(input[0])-1 {
		return false
	}
	return true
}

func fetchPossibleDirections(node queueNode, matrix [][]int, min, max int) []direction {
	output := []direction{}
	switch node.lastDir {
	case up, down:
		if node.count < max {
			output = append(output, node.lastDir)
		}
		if node.count >= min {
			output = append(output, left)
			output = append(output, right)
		}
	case left, right:
		if node.count < max {
			output = append(output, node.lastDir)
		}
		if node.count >= min {
			output = append(output, up)
			output = append(output, down)
		}
	}
	return output
}

func getNodesFromDirection(dirs []direction, startNode queueNode, matrix [][]int) []queueNode {
	nodes := []queueNode{}
	for _, dir := range dirs {
		addPoint := positionMap[dir]
		newPoint := point{startNode.x + addPoint.x, startNode.y + addPoint.y}

		if !isValid(newPoint, matrix) {
			continue
		}

		count := 1
		if startNode.lastDir == dir {
			count = startNode.count + 1
		}

		nodes = append(nodes, queueNode{
			step{
				x:       startNode.x + addPoint.x,
				y:       startNode.y + addPoint.y,
				lastDir: dir,
				count:   count,
			},
			startNode.heatLoss + matrix[newPoint.x][newPoint.y],
		})
	}
	return nodes
}
