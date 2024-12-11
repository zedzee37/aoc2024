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

func (v Vec2) sub(other Vec2) Vec2 {
	return Vec2{v[0] - other[0], v[1] - other[1]}
}

func (v Vec2) mul(n int) Vec2 {
	return Vec2{v[0] * n, v[1] * n}
}

func (v Vec2) div(n int) Vec2 {
	return Vec2{v[0] / n, v[1] / n}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findStartLocations(grid []string) []Vec2 {
	locations := make([]Vec2, 0)

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			if grid[y][x] == '0' {
				locations = append(locations, Vec2{x, y})
			}
		}
	}

	return locations
}

func isOffGrid(grid []string, loc Vec2) bool {
	l := len(grid)
	return loc[0] < 0 || loc[1] < 0 || loc[0] >= l || loc[1] >= l
}

func surroundingPositions(grid []string, loc Vec2) []Vec2 {
	surrounding := make([]Vec2, 0)
	directions := []Vec2{
		{0, -1}, // North
		{0, 1},  // South
		{-1, 0}, // West
		{1, 0},  // East
	}
	for _, dir := range directions {
		pos := loc.add(dir)

		if !isOffGrid(grid, pos) {
			surrounding = append(surrounding, pos)
		}
	}
	return surrounding
}

func print(size int, pos Vec2, final bool) {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if y == pos[1] && x == pos[0] {
				if final {
					fmt.Print("F")
				} else {
					fmt.Print("X")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func countPaths(grid []string, from Vec2, i int, endPositions []Vec2) ([]Vec2, error) {
	for _, pos := range surroundingPositions(grid, from) {
		atPos := grid[pos[1]][pos[0]]
		numAtPos, err := strconv.Atoi(string(atPos))
		if err != nil {
			return []Vec2{}, err
		}

		if numAtPos == i {
			if i == 9 {
				print(len(grid), pos, true)
				return []Vec2{pos}, nil
			} else {
				print(len(grid), pos, false)
			}

			positions, err := countPaths(grid, pos, i+1, endPositions)
			endPositions = append(endPositions, positions...)

			if err != nil {
				return []Vec2{}, err
			}
		}
	}

	return endPositions, nil
}

func pathCount(grid []string, startLocations []Vec2) (int, error) {
	sum := 0

	for _, startLoc := range startLocations {
		uniqueEnds := make(map[Vec2]bool)
		positions, err := countPaths(grid, startLoc, 1, []Vec2{})

		for _, pos := range positions {
			if !uniqueEnds[pos] {
				uniqueEnds[pos] = true
			}
		}

		if err != nil {
			return -1, err
		}

		sum += len(uniqueEnds)
	}

	return sum, nil
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	text := string(data)
	grid := strings.Split(text, "\n")

	startLocations := findStartLocations(grid)
	count, err := pathCount(grid, startLocations)
	check(err)

	fmt.Println(count)
}
