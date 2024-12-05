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

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	text := string(data)
	fmt.Println(text)

	grid := strings.Split(text, "\n")
	grid = grid[0 : len(grid)-1]

	matchCount := accumulateMatches(grid)
	fmt.Println(matchCount)
}
