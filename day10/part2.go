package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vec2 [2]int

func (v Vec2) add(other Vec2) Vec2 {
	return Vec2{v[0] + other[0], v[1] + other[1]}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findStartLocations(grid []string) []Vec2 {
	locations := make([]Vec2, 0)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '0' {
				locations = append(locations, Vec2{x, y})
			}
		}
	}
	return locations
}

func isOffGrid(grid []string, loc Vec2) bool {
	return loc[0] < 0 || loc[1] < 0 || loc[1] >= len(grid) || loc[0] >= len(grid[loc[1]])
}

func get(grid []string, loc Vec2) byte {
	return grid[loc[1]][loc[0]]
}

func surroundingPositions(grid []string, loc Vec2) []Vec2 {
	directions := []Vec2{
		{0, -1}, // North
		{0, 1},  // South
		{-1, 0}, // West
		{1, 0},  // East
	}
	surrounding := make([]Vec2, 0)
	for _, dir := range directions {
		pos := loc.add(dir)
		if !isOffGrid(grid, pos) {
			surrounding = append(surrounding, pos)
		}
	}
	return surrounding
}

func countPaths(grid []string, loc Vec2, i int, visited map[Vec2]bool) int {
	if i > 9 {
		return 1
	}

	sum := 0
	for _, pos := range surroundingPositions(grid, loc) {
		if visited[pos] {
			continue
		}
		numRune := get(grid, pos)
		num, err := strconv.Atoi(string(numRune))
		if err == nil && num == i {
			visited[pos] = true
			sum += countPaths(grid, pos, i+1, visited)
			visited[pos] = false
		}
	}
	return sum
}

func pathCount(grid []string, startLocations []Vec2) int {
	count := 0
	for _, loc := range startLocations {
		visited := make(map[Vec2]bool)
		visited[loc] = true
		count += countPaths(grid, loc, 1, visited)
	}
	return count
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	text := string(data)
	grid := strings.Split(strings.TrimSpace(text), "\n")

	startLocations := findStartLocations(grid)
	count := pathCount(grid, startLocations)

	fmt.Println(count)
}
