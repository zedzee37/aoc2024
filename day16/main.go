package main

import (
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

func (v Vec2) neg() Vec2 {
	return v.mul(-1)
}

func parseInput(input string) []string {
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

}
