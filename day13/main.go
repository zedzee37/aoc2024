package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Mat2x2i [2][2]int

type Mat2x1i [2]int

func (m Mat2x2i) det() int {
	return (m[0][0]*m[1][1]) - (m[0][1]*m[1][0])
}

func (m Mat2x2i) inv() Mat2x2i {
	return Mat2x2i{
		{ m[1][1], -m[0][1] },
		{ -m[1][0], m[0][0] },
	}
}

func (m Mat2x2i) mul(n int) Mat2x2i {
	return Mat2x2i{
		{ m[0][0]*n, m[0][1]*n },
		{ m[1][0]*n, m[1][1]*n },
	}
}

func (m Mat2x1i) add(other Mat2x1i) Mat2x1i {
	return Mat2x1i{ m[0] + other[0], m[1] + other[1] }
}

func (m Mat2x1i) mul(n int) Mat2x1i {
	return Mat2x1i{ m[0]*n, m[1]*n }
}

func (m Mat2x1i) cross(other Mat2x2i) Mat2x1i {
	mat1 := Mat2x1i(other[0]).mul(m[0])
	mat2 := Mat2x1i(other[1]).mul(m[1])

	return mat1.add(mat2)
}

type Game struct {
	xA int
	yA int
	xB int
	yB int
	winX int
	winY int
}

const X_COST = 3
const Y_COST = 1

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func parseLine(line string, regex regexp.Regexp) ([]int, error) {
	matches := make([]int, 0)
	matchedLine := regex.FindAllString(line, -1)
	
	for _, num :=  range matchedLine {
		n, err := strconv.Atoi(num)

		if err != nil {
			return nil, err
		}

		matches = append(matches, n)
	}
	return matches, nil
}

func parseInput(contents string) ([]Game, error) {
	games := make([]Game, 0)
	gamesStr := strings.Split(contents, "\n\n")
	numRegex, err := regexp.Compile("(\\d+)+")

	if err != nil {
		return nil, err
	}

	for _, gameStr := range gamesStr {
		lines := strings.Split(gameStr, "\n") 
		
		lineA, err := parseLine(lines[0], *numRegex)			
		if err != nil {
			return nil, err
		}

		lineB, err := parseLine(lines[1], *numRegex)			
		if err != nil {
			return nil, err
		}

		finalLine, err := parseLine(lines[2], *numRegex)			
		if err != nil {
			return nil, err
		}

		games = append(games, Game{
			xA: lineA[0],
			yA: lineA[1],
			xB: lineB[0],
			yB: lineB[1],
			winX: finalLine[0],
			winY: finalLine[1],
		})
	}

	return games, nil
}

func greatestCommonDivisor(a int, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func leastCommonMultiple(a int, b int) int {
	return a*b / greatestCommonDivisor(a, b)
}

func cheapestToWin(game Game) int {
	lcm := leastCommonMultiple(game.xA, game.yA)
	mulX := lcm / game.xA
	mulY := lcm / game.yA

	realBCount := float64(game.winX*mulX - game.winY*mulY) / float64(game.xB*mulX - game.yB*mulY)

	if math.Round(realBCount) != realBCount {
		return -1
	}
	
	bCount := int(realBCount)	
	realACount := float64(game.winX - game.xB*bCount) / float64(game.xA)

	if math.Round(realACount) != realACount {
		return -1
	}

	return int(realACount)*3 + bCount
}

func main() {
	fileContents, err := os.ReadFile("input.txt")
	check(err)

	contents := string(fileContents)	
	games, err := parseInput(contents)
	check(err)

	sum := 0
	for _, game := range games {
		res := cheapestToWin(game)
		if res != -1 {
			fmt.Println(res)
			sum += res
		}
	}
	fmt.Println(sum)
}
