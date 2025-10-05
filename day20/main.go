package main

import (
	"os"
	"strings"
)

type Grid map[Vec2]*PathNode

type Vec2 struct {
	x int
	y int
}

type PathNode struct {
	GCost int
	next  *PathNode
}

type MapInfo struct {
	min   Vec2
	max   Vec2
	start Vec2
	end   Vec2
}

func NewPathNode(GCost int) *PathNode {
	return &PathNode{
		GCost: 0,
		next:  nil,
	}
}

func isInBounds(pos Vec2, mapInfo MapInfo) bool {
	return pos.x >= mapInfo.min.x && pos.y >= mapInfo.min.y && pos.x <= mapInfo.max.x && pos.y <= mapInfo.max.y
}

func parseInput(fp string) (Grid, MapInfo, error) {
	fileContents, err := os.ReadFile(fp)
	if err != nil {
		return nil, MapInfo{}, err
	}

	lines := strings.Split(string(fileContents), "\n")

	mapInfo := MapInfo{}
	grid := make(Grid)

	mapInfo.min.x = 0
	mapInfo.min.y = 0
	mapInfo.max.y = len(lines) - 1

	for y, line := range lines {
		for x, char := range line {
			pos := Vec2{
				x: x, y: y,
			}

			if char == 'S' {
				mapInfo.start = pos
			} else if char == 'E' {
				mapInfo.end = pos
			} else if char == '#' {
				continue
			}

			grid[pos] = NewPathNode(0)
			mapInfo.max.x = max(pos.x, mapInfo.max.x)
		}
	}

	mapInfo.max.x += 2

	return grid, mapInfo, nil
}

func printGrid(grid Grid, mapInfo MapInfo) {
	for y := range mapInfo.max.y {
		for x := range mapInfo.max.x {
			pos := Vec2{x: x, y: y}

			_, exists := grid[pos]
			if exists {
				print(".")
			} else {
				print("#")
			}
		}
		println()
	}
}

func main() {
	grid, mapInfo, err := parseInput("input.txt")
	if err != nil {
		panic(err.Error())
	}
	printGrid(grid, mapInfo)
}
