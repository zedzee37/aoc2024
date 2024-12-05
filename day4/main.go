package main

import (
	"fmt"
	"os"
	"strings"
)

const TARGET_WORD = "XMAS"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkInDir(
	grid []string,
	startX int, startY int,
	dx int, dy int,
) bool {
	for i := 0; i < len(TARGET_WORD); i++ {
		x := startX + i*dx
		y := startY + i*dy
		target := TARGET_WORD[i]

		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid) {
			return false
		}

		actual := grid[x][y]

		if target != actual {
			return false
		}
	}

	return true
}

func matchesAtPos(grid []string, x int, y int) uint {
	var count uint = 0

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if checkInDir(grid, x, y, dx, dy) {
				count++
			}
		}
	}

	return count
}

func accumulateMatches(grid []string) uint {
	var count uint = 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			count += matchesAtPos(grid, x, y)
		}
	}

	return count
}

func part1(text string) {
	grid := strings.Split(text, "\n")
	grid = grid[0 : len(grid)-1]

	matchCount := accumulateMatches(grid)
	fmt.Println(matchCount)
}

const MAS_STRING = "MAS"

func isOnEdge(grid []string, x int, y int) bool {
	l := len(grid)
	return x > 0 && y > 0 && x < l-1 && y < l-1
}

func hasX(grid []string, x int, y int) bool {
	if grid[x][y] != MAS_STRING[1] || !isOnEdge(grid, x, y) {
		return false
	}

	for dy := -1; dy <= 1; dy += 2 {
		for dx := -1; dx <= 1; dx += 2 {
			newX := x + dx
			newY := y + dy

			ch := grid[newX][newY]

			if ch == MAS_STRING[0] || ch == MAS_STRING[2] {
				toMatch := MAS_STRING[0]

				if ch == MAS_STRING[0] {
					toMatch = MAS_STRING[2]
				}

				newX = x - dx
				newY = y - dy

				if grid[newX][newY] != toMatch {
					return false
				}
			} else {
				return false
			}
		}
	}

	return true
}

func part2(text string) {
	grid := strings.Split(text, "\n")
	grid = grid[0 : len(grid)-1]

	count := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			if hasX(grid, x, y) {
				count++
			}
		}
	}

	fmt.Println(count)
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	text := string(data)

	part2(text)
}
