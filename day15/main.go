package main

import (
	"fmt"
	"os"
	"strings"
)

type Vec2 [2]int

func (v Vec2) add(other Vec2) Vec2 {
	return Vec2{v[0] + other[0], v[1] + other[1]}
}

func (v Vec2) sub(other Vec2) Vec2 {
	return Vec2{v[0] - other[0], v[1] - other[1]}
}

func (v Vec2) mul(n int) Vec2 {
	return Vec2{v[0] * n, v[1] * n}
}

func (v Vec2) div(n int) Vec2 {
	return Vec2{v[0] / n, v[1] / n}
}

func (v Vec2) neg() Vec2 {
	return v.mul(-1)
}

func (v Vec2) right() Vec2 {
	return Vec2{v[1], -v[0]}
}

func (v Vec2) left() Vec2 {
	return Vec2{-v[1], v[0]}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (v Vec2) manhattanDistance(other Vec2) int {
	return abs(v[0]-other[0]) + abs(v[1]-other[1])
}

var cardinals [4]Vec2 = [4]Vec2{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func findChar(grid []string, ch byte) (Vec2, error) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			if grid[y][x] == ch {
				return Vec2{x, y}, nil
			}
		}
	}

	return Vec2{-1, -1}, fmt.Errorf("Could not find character: '%c'.", ch)
}

func findShortestPath(grid []string) ([]Vec2, error) {
	startPos, err := findChar(grid, 'S')
	if err != nil {
		return []Vec2{}, err
	}

	endPos, err := findChar(grid, 'E')
	if err != nil {
		return []Vec2{}, err
	}

	path := make([]Vec2, 0)
	return path, nil
}

func getPathCost(path []Vec2) int {
	cost := 0
	prevCell := path[0]
	prevDir := Vec2{1, 0}
	for _, cell := range path[1:] {
		dir := prevCell.sub(cell)

		if dir != prevDir {
			cost += 1000
		}
		cost++
	}
	return cost
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fileContents, err := os.ReadFile("input.txt")
	check(err)

	grid := parseInput(string(fileContents))
	shortestPath, err := findShortestPath(grid)
	check(err)

	cost := getPathCost(shortestPath)
	fmt.Println(cost)
}
