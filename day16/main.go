package main

import (
	"fmt"
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

type AStarNode struct {
	pos Vec2
	g int
	f int
	from *AStarNode
}

func (node AStarNode) getPathCost() int {
	cost := 1
	current := node.from
	prevDir := emptyVec

	for current != nil {
		next := current.from
	
		if next == nil {
			break
		}

		dir := next.pos.sub(current.pos)

		cost++
		current = next
		if prevDir == emptyVec {
			prevDir = dir
		} else if prevDir != dir {
			cost += 1000
		}
	}

	return cost
}

func (node AStarNode) getPath() map[Vec2]bool {
	path := make(map[Vec2]bool)
	current := &node

	for current != nil {
		path[current.pos] = true
		current = current.from
	}

	return path
}

type Grid []string

func (grid Grid) printPath(path AStarNode) {
	a := path.getPath()	

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			pos := Vec2{x, y}
			
			ch := grid.get(pos)

			if a[pos] {
				fmt.Print("P")
			} else {
				fmt.Print(string(ch))
			}
		}
		fmt.Println()
	}
}

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

func (grid Grid) findShortestPath() (AStarNode, error) {
	startPos, endPos, err := grid.findEndPoints()
	
	if err != nil {
		return AStarNode{}, err
	}

	visited := make(map[Vec2]bool)
	
	current := make([]AStarNode, 0)
	startDst := startPos.manhattanDst(endPos)
	current = append(current, AStarNode{
		pos: startPos,
		g: 0,
		f: startDst,
		from: nil,
	})

	for len(current) > 0 {
		lowestCostIdx := findLowestCost(current)	
		lowestCostCell := current[lowestCostIdx]
		current = slices.Delete(current, lowestCostIdx, lowestCostIdx+1)

		if visited[lowestCostCell.pos] {
			continue
		}
		
		visited[lowestCostCell.pos] = true
		surrounding := grid.getSurrounding(lowestCostCell.pos)
		
		for _, neighbor := range surrounding {
			if visited[neighbor] {
				continue
			}

			gCost := lowestCostCell.g + 1
			dst := neighbor.manhattanDst(endPos)
			cell := AStarNode{
				pos: neighbor,
				g: gCost,
				f: dst + gCost,
				from: &lowestCostCell,
			}

			if neighbor == endPos {
				return cell, nil
			}

			current = append(current, cell)
		}
	}

	return AStarNode{}, fmt.Errorf("Could not find viable path to end!")
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
	shortestPath, err := grid.findShortestPath()
	check(err)

	grid.printPath(shortestPath)

	cost := shortestPath.getPathCost()
	fmt.Println(cost)
}
