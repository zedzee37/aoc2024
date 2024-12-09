package main

import (
	"fmt"
	"os"
	"strings"
)

type Grid []string

type Vec2 struct {
	x int
	y int
}

type State struct {
	pos Vec2
	dir Vec2
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isOffGrid(size int, pos Vec2) bool {
	return pos.x < 0 || pos.y < 0 || pos.x >= size || pos.y >= size
}

func right(dir Vec2) Vec2 {
	return Vec2{x: dir.y, y: -dir.x}
}

func add(a Vec2, b Vec2) Vec2 {
	return Vec2{x: a.x + b.x, y: a.y - b.y}
}

func findGuard(grid Grid) Vec2 {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '^' {
				return Vec2{x: x, y: y}
			}
		}
	}
	return Vec2{x: -1, y: -1}
}

func printGrid(grid Grid) {
	for _, str := range grid {
		fmt.Println(str)
	}
	fmt.Println()
}

func tracePath(grid Grid, startPos Vec2, obstruction Vec2) (bool, []State) {
	pos := startPos
	dir := Vec2{x: 0, y: 1}
	path := make([]State, 0)
	size := len(grid)
	seen := make(map[State]bool)

	for true {
		newPos := add(pos, dir)

		if isOffGrid(size, newPos) {
			break
		}

		for newPos == obstruction || grid[newPos.y][newPos.x] == '#' {
			dir = right(dir)
			newPos = add(pos, dir)
		}

		if isOffGrid(size, newPos) {
			break
		}

		state := State{
			pos: newPos,
			dir: dir,
		}
		if seen[state] {
			// for y := range size {
			// 	for x := range size {
			// 		if y == pos.y && x == pos.x {
			// 			fmt.Print("^")
			// 		} else if y == obstruction.y && x == obstruction.x {
			// 			fmt.Print("O")
			// 		} else if grid[y][x] == '#' {
			// 			fmt.Print("#")
			// 		} else {
			// 			fmt.Print(".")
			// 		}
			// 	}
			// 	fmt.Println()
			// }
			// fmt.Println()
			return true, path
		}

		path = append(path, state)
		seen[state] = true
		pos = newPos
	}

	return false, path
}

func uniquePath(path []State) []Vec2 {
	seen := make(map[Vec2]bool)
	unique := make([]Vec2, 1)

	for _, state := range path {
		if seen[state.pos] {
			continue
		}

		seen[state.pos] = true
		unique = append(unique, state.pos)
	}

	return unique
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	text := string(data)
	grid := strings.Split(text, "\n")
	guardPos := findGuard(grid)

	// Check if the initial path already has a loop
	_, initialPath := tracePath(grid, guardPos, Vec2{x: -1, y: -1})
	unique := uniquePath(initialPath)

	count := 0
	for _, pos := range unique[1:] {
		looped, _ := tracePath(grid, guardPos, pos)
		if looped {
			count++
		}
	}

	fmt.Println(count)
}
