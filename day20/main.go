package main

import (
	"fmt"
	"os"
	"strings"
)

type Grid map[Vec2]*PathNode

type Vec2 struct {
	x int
	y int
}

func (v1 Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{
		x: v1.x + v2.x,
		y: v1.y + v2.y,
	}
}

type PathNode struct {
	DistanceToEnd int
	next          *PathNode
}

type MapInfo struct {
	min   Vec2
	max   Vec2
	start Vec2
	end   Vec2
}

func NewPathNode(GCost int) *PathNode {
	return &PathNode{
		DistanceToEnd: 0,
		next:          nil,
	}
}

func isInBounds(pos Vec2, mapInfo *MapInfo) bool {
	return pos.x >= mapInfo.min.x && pos.y >= mapInfo.min.y && pos.x <= mapInfo.max.x && pos.y <= mapInfo.max.y
}

func parseInput(fp string) (Grid, *MapInfo, error) {
	fileContents, err := os.ReadFile(fp)
	if err != nil {
		return nil, &MapInfo{}, err
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

	return grid, &mapInfo, nil
}

var directions [4]Vec2 = [4]Vec2{
	Vec2{x: 1, y: 0},
	Vec2{x: -1, y: 0},
	Vec2{x: 0, y: 1},
	Vec2{x: 0, y: -1},
}

func tracePath(grid Grid, mapInfo *MapInfo) map[Vec2]bool {
	currentPos := mapInfo.end

	wallPositions := make(map[Vec2]bool, 0)
	visited := make(map[Vec2]bool, 0)

	currentCost := 0
	for currentPos != mapInfo.start {
		p := Vec2{}
		for _, direction := range directions {
			newPos := currentPos.Add(direction)

			_, hasVisited := visited[newPos]
			if hasVisited {
				continue
			}

			cell, exists := grid[newPos]
			if exists {
				visited[currentPos] = true
				p = newPos
				currentCost++
				cell.DistanceToEnd = currentCost
			} else {
				wallPositions[newPos] = true
			}
		}
		currentPos = p
	}

	return wallPositions
}

func printGrid(grid Grid, mapInfo *MapInfo, walls map[Vec2]bool) {
	const Padding = 3

	for y := 0; y < mapInfo.max.y; y++ {
		for x := 0; x < mapInfo.max.x; x++ {
			pos := Vec2{x: x, y: y}

			cell, exists := grid[pos]
			_, wallExists := walls[pos]
			if exists {
				fmt.Printf("%*d", Padding, cell.DistanceToEnd)
			} else if wallExists {
				fmt.Printf("%*s", Padding, "@")
			} else {
				fmt.Printf("%*s", Padding, "#")
			}
		}
		fmt.Println()
	}
}

func main() {
	grid, mapInfo, err := parseInput("input.txt")
	if err != nil {
		panic(err.Error())
	}
	walls := tracePath(grid, mapInfo)
	printGrid(grid, mapInfo, walls)
}
