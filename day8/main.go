package main

import (
	"fmt"
	"os"
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

func isOffGrid(size int, pos Vec2) bool {
	return pos[0] < 0 || pos[1] < 0 || pos[0] >= size || pos[1] >= size
}

func getAntennas(grid []string) map[byte][]Vec2 {
	antennaLocations := make(map[byte][]Vec2, 0)

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			if grid[y][x] == '.' {
				continue
			}

			freq := grid[y][x]

			frequencyPositions, ok := antennaLocations[freq]
			if !ok {
				frequencyPositions = make([]Vec2, 0)
				antennaLocations[freq] = frequencyPositions
			}

			pos := Vec2{x, y}
			frequencyPositions = append(frequencyPositions, pos)
			antennaLocations[freq] = frequencyPositions
		}
	}

	return antennaLocations
}

func partOne(grid []string) int {
	antennaLocations := getAntennas(grid)
	size := len(grid)
	antinodeLocations := make(map[Vec2]bool)

	for key := range antennaLocations {
		frequencyPositions := antennaLocations[key]

		for i, pos := range frequencyPositions {
			if i == len(frequencyPositions)-1 {
				continue
			}
			requiredAntennas := frequencyPositions[i+1:]

			for _, otherPos := range requiredAntennas {
				dirFrom := pos.sub(otherPos)
				dirTo := otherPos.sub(pos)

				antinode := pos.add(dirFrom)
				otherAntinode := otherPos.add(dirTo)

				if !isOffGrid(size, antinode) {
					antinodeLocations[antinode] = true
				}

				if !isOffGrid(size, otherAntinode) {
					antinodeLocations[otherAntinode] = true
				}
			}
		}
	}

	return len(antinodeLocations)
}

func partTwo(grid []string) int {
	antennaLocations := getAntennas(grid)
	size := len(grid)
	antinodeLocations := make(map[Vec2]bool)

	for key := range antennaLocations {
		frequencyPositions := antennaLocations[key]

		for i, pos := range frequencyPositions {
			if i == len(frequencyPositions)-1 {
				continue
			}
			requiredAntennas := frequencyPositions[i+1:]

			for _, otherPos := range requiredAntennas {
				antinodeLocations[pos] = true
				antinodeLocations[otherPos] = true
				dirFrom := pos.sub(otherPos)
				dirTo := otherPos.sub(pos)

				antinode := pos.add(dirFrom)
				otherAntinode := otherPos.add(dirTo)

				for !isOffGrid(size, antinode) {
					antinodeLocations[antinode] = true
					antinode = antinode.add(dirFrom)
				}

				for !isOffGrid(size, otherAntinode) {
					antinodeLocations[otherAntinode] = true
					otherAntinode = otherAntinode.add(dirTo)
				}
			}
		}
	}

	return len(antinodeLocations)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	text := string(data)
	grid := strings.Split(text, "\n")

	count := partTwo(grid)
	fmt.Println(count)
}
