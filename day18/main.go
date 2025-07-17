package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AStarCell struct {
	IsWall bool
	GCost float32
	HCost float32
} 

func to1D(x int, y int) int {
	return (y * 71) + x
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

	printGrid(grid)
}

