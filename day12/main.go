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

func (v Vec2) neg() Vec2 {
	return v.mul(-1)
}

func (v Vec2) right() Vec2 {
	return Vec2{v[1], -v[0]}
}

func (v Vec2) left() Vec2 {
	return Vec2{-v[1], v[0]}
}

var cardinals [4]Vec2 = [4]Vec2{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

var ordinals [4]Vec2 = [4]Vec2{
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
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

func printEdges(size int, edges map[Vec2]bool) {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			pos := Vec2{x, y}

			if edges[pos] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func isValid(grid []string, pos Vec2, ch byte) bool {
	size := len(grid)
	return !isOffGrid(size, pos) && get(grid, pos) == ch
}

func isHorizontal(direction Vec2) bool {
	return direction == Vec2{1, 0} || direction == Vec2{-1, 0}
}

func countCorners(grid []string, ch byte, pos Vec2) int {
	count := 0
	cardinalCount := 0

	for _, dir := range cardinals {
		perpendicular := dir.left()
		perpendicularPos := pos.add(perpendicular)

		dirPos := pos.add(dir)
		diagonalInner := pos.add(dir.add(perpendicular))

		if !isValid(grid, dirPos, ch) {
			continue
		}

		cardinalCount++

		if !isValid(grid, perpendicularPos, ch) {
			continue
		}

		invDirPos := pos.add(dir.neg())
		invPerpenPos := pos.add(perpendicular.neg())

		if !isValid(grid, invDirPos, ch) && !isValid(grid, invPerpenPos, ch) {
			count++
		}

		isInner := isValid(grid, diagonalInner, ch)

		if !isInner {
			count++
		}
	}

	if cardinalCount == 1 {
		count += 2
	}

	fmt.Printf("There is %d corners at %d, %d, for the letter: %c\n", count, pos[0], pos[1], ch)

	return count
}

func countSides(grid []string, edges map[Vec2]bool, ch byte) int {
	if len(edges) == 1 {
		return 4
	}

	count := 0
	for edge := range edges {
		count += countCorners(grid, ch, edge)
	}
	return count
}

func getBulkFencePrice(grid []string) int {
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

			count := 1

			ch := get(grid, current[0])
			visitedEdges := make(map[Vec2]bool)
			sideCount := 0

			for len(current) > 0 {
				newCurrent := make([]Vec2, 0)

				for _, pos := range current {
					sideCount += countCorners(grid, ch, pos)

					for _, dir := range cardinals {
						neighbor := pos.add(dir)

						if isOffGrid(size, neighbor) || get(grid, neighbor) != ch {
							visitedEdges[pos] = true
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

			if count == 1 {
				sideCount = 4
			}
			fmt.Printf("There is %d for the letter: %c\n", sideCount, ch)
			fmt.Println()
			fmt.Println(count)
			price += sideCount * count
		}
	}

	return price
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
	price := getBulkFencePrice(grid)
	fmt.Println(price)
}
