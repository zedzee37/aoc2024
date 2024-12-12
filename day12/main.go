package main

import (
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Region struct {
	cells []Vec2
	perimeter int
}

func (r Region) area() int {
	return len(r.cells)
}

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

var cardinals [4]Vec2 = [4]Vec2{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func isOffGrid(size int, pos Vec2) bool {
	return pos[0] < 0 || pos[1] < 0 || pos[0] >= size || pos[1] >= size
}

func surroundingPositions(pos Vec2) []Vec2 {
	surrounding := make([]Vec2, 0)
	for _, cardinal := range cardinals {
		surrounding = append(surrounding, pos.add(cardinal))
	}
	return surrounding
}

func get(grid []string, pos Vec2) byte {
	return grid[pos[1]][pos[0]]
}

func getRegions(grid []string) []*Region {
	regionMap := make(map[Vec2]*Region)
	regions := make([]*Region, 0)
	size := len(grid)
	
	y := 0
	x := 0
	for y < size {
		for x < size {
			x++
		}
		y++
	}

	return regions
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	contents := string(data)
	grid := strings.Split(contents, "\n")
	grid = grid[:len(grid) - 1]
	regions := getRegions(grid)

	sum := 0
	for _, region := range regions {
		sum += region.perimeter * region.area()
		fmt.Println(region.perimeter)
	}
	fmt.Println(sum)
}
