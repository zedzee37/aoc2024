package main

import (
	"fmt"
	"math"
	"os"
	"slices"
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

func (v Vec2) neg() Vec2 {
	return v.mul(-1)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (v Vec2) manhattanDst(other Vec2) int {
	return abs(v[0] - other[0]) + abs(v[1] - other[1])
}

var cardinals = [4]Vec2{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

var emptyVec = Vec2{}

type State struct {
	pos Vec2
	dir Vec2
}

type AStarNode struct {
	pos Vec2
	g int
	f int
	from Vec2
	cost int
}

type Grid []string

func (grid Grid) get(pos Vec2) byte {
	return grid[pos[1]][pos[0]]
}

func (grid Grid) isOffGrid(pos Vec2) bool {
	return pos[0] < 0 || pos[1] < 0 || pos[0] >= len(grid) || pos[1] >= len(grid)
}

func (grid Grid) getSurrounding(from Vec2) []Vec2 {
	surrounding := make([]Vec2, 0)

	for _, dir := range cardinals {
		pos := from.add(dir)
		
		if !grid.isOffGrid(pos) && grid.get(pos) != '#' {
			surrounding = append(surrounding, pos)	
		}
	}

	return surrounding
}

func (grid Grid) findEndPoints() (Vec2, Vec2, error) {
	startPoint := emptyVec
	endPoint := emptyVec

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			if startPoint != emptyVec && endPoint != emptyVec {
				return startPoint, endPoint, nil
			}

			pos := Vec2{x, y}
			ch := grid.get(pos)
			switch ch {
			case 'S':
				startPoint = pos
			case 'E':
				endPoint = pos
			}
		}
	}

	return startPoint, endPoint, fmt.Errorf("Failed to find start/end point of the path!")
}

func findLowestCost(nodes []AStarNode) int {
	lowestCost := nodes[0].f
	lowestCostIdx := 0
	for i, node := range nodes {
		if node.f < lowestCost {
			lowestCostIdx = i
			lowestCost = node.f
		}
	}
	return lowestCostIdx
}

func (grid Grid) findShortestPaths() ([]AStarNode, error) {
	startPos, endPos, err := grid.findEndPoints()
	paths := make([]AStarNode, 0)
	
	if err != nil {
		return []AStarNode{}, err
	}

	visited := make(map[State]bool)
	
	current := make([]AStarNode, 0)
	startDst := startPos.manhattanDst(endPos)
	current = append(current, AStarNode{
		pos: startPos,
		g: 0,
		f: startDst,
		from: Vec2{1, 0},
		cost: 0,
	})

	for len(current) > 0 {
		lowestCostIdx := findLowestCost(current)	
		lowestCostCell := current[lowestCostIdx]
		current = slices.Delete(current, lowestCostIdx, lowestCostIdx+1)
		lowestCostState := State{
			lowestCostCell.pos,
			lowestCostCell.from,
		}

		if visited[lowestCostState] {
			continue
		}
		
		visited[lowestCostState] = true
		surrounding := grid.getSurrounding(lowestCostCell.pos)
		
		for _, neighbor := range surrounding {
			dir := lowestCostCell.pos.sub(neighbor)
			state := State{
				pos: neighbor,
				dir: dir,
			}

			if visited[state] {
				continue
			}

			cost := lowestCostCell.cost+1
			if dir != lowestCostCell.from {
				cost += 1000
			}

			gCost := lowestCostCell.g + 1
			dst := neighbor.manhattanDst(endPos)
			cell := AStarNode{
				pos: neighbor,
				g: gCost,
				f: dst + gCost,
				from: dir,
				cost: cost,
			}

			if neighbor == endPos {
				paths = append(paths, cell)
			} else {
				current = append(current, cell)
			}
		}
	}

	return paths, nil
}

func findLowestCostPath(paths []AStarNode) int {
	lowestCost := math.MaxInt32
	for _, path := range paths {
		lowestCost = min(lowestCost, path.cost)
	}
	return lowestCost
}

func parseInput(input string) Grid {
	return strings.Split(input, "\n")
}

func check(err error) {
	if err != nil {
		check(err)
	}
}

func main() {
	fileContents, err := os.ReadFile("input.txt")
	check(err)

	grid := parseInput(string(fileContents))
	grid = grid[:len(grid)-1]
	shortestPaths, err := grid.findShortestPaths()
	check(err)

	lowestCost := findLowestCostPath(shortestPaths)
	fmt.Println(lowestCost)
}
