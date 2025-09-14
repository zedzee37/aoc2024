package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BinaryNode[T any] struct {
	value T
	left  *BinaryNode[T]
	right *BinaryNode[T]
}

type BinaryHeap[T comparable] struct {
	head BinaryNode[T]

	// returns if a should be above b
	compareFunc func(a T, b T) bool
}

type AStarCell struct {
	IsWall   bool
	GCost    float64
	HCost    float64
	CameFrom *AStarCell
}

func (cell AStarCell) fCost() float64 {
	return cell.GCost + cell.HCost
}

func to1D(x int, y int) int {
	return (y * 71) + x
}

func isOnGrid(idx int, gridSize int) bool {
	return idx >= 0 && idx < gridSize
}

func neighbors(idx int, grid []AStarCell) []int {
	neighbors := make([]int, 0)

	directions := []int{
		// left
		to1D(-1, 0),

		// right
		to1D(1, 0),

		// up
		to1D(0, -1),

		// down
		to1D(0, 1),
	}

	// iterate through every direction
	for _, direction := range directions {
		neighbor := idx + direction

		if isOnGrid(neighbor, len(grid)) {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func findPath(grid []AStarCell) *AStarCell {
	target := len(grid) - 1
	current := 0

	for current != target {
		neighbors := neighbors(current, grid)

		if len(neighbors) == 0 {
			continue
		}
	}

	return &grid[target]
}

func parseInput(fileContents string) ([]AStarCell, error) {
	grid := make([]AStarCell, 71*71)

	// lines := strings.SplitSeq(fileContents, "\n")
	lines := strings.Split(fileContents, "\n")
	for i := range 1024 {
		line := lines[i]
		if line == "" {
			continue
		}
		nums := strings.Split(line, ",")

		x, err := strconv.Atoi(nums[0])
		if err != nil {
			return nil, err
		}

		y, err := strconv.Atoi(nums[1])
		if err != nil {
			return nil, err
		}

		idx := to1D(x, y)
		grid[idx].IsWall = true
	}

	return grid, nil
}

func printGrid(grid []AStarCell) {
	for y := range 71 {
		for x := range 71 {
			idx := to1D(x, y)

			if grid[idx].IsWall {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	fileContents, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	grid, err := parseInput(string(fileContents))
	if err != nil {
		panic(err)
	}

	path := findPath(grid)
	fmt.Println(path.GCost)
}
