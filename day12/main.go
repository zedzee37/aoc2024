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
	cells     []Vec2
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

func getIdx(size int, pos Vec2) int {
	return pos[0]*size + pos[1]
}

func getFencePrice(grid []string) int {
	size := len(grid)

	visited := make(map[Vec2]bool)
	price := 0

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			curPos := Vec2{x, y}
			if visited[curPos] {
				continue
			}

			visited[curPos] = true
			current := make([]Vec2, 1)
			current[0] = Vec2{x, y}	

			perimiter := 0
			count := 1

			ch := get(grid, current[0])
			
			for len(current) > 0 {
				newCurrent := make([]Vec2, 0)

				for _, pos := range current {
					neighbors := surroundingPositions(pos)

					for _, neighbor := range neighbors {
						if isOffGrid(size, neighbor) || get(grid, neighbor) != ch {
							perimiter++
							continue
						}

						if !visited[neighbor] {
							visited[neighbor] = true
							newCurrent = append(newCurrent, neighbor)
							count++
						}
					}
				}

				current = newCurrent
			}

			price += perimiter * count
		}
	}

	return price
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	contents := string(data)
	grid := strings.Split(contents, "\n")
	grid = grid[:len(grid)-1]
	price := getFencePrice(grid)
	fmt.Println(price)
}
