package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

func parseInput(contents string) ([]Game, err) {
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

		lineB, err := parseLine(lines[0], *numRegex)			
		if err != nil {
			return nil, err
		}

		finalLine, err := parseLine(lines[0], *numRegex)			
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

func cheapestToWin(game Game) int {
	count := 0
	return count
}

func main() {
	fileContents, err := os.ReadFile("input.txt")
	check(err)

	contents := string(fileContents)	
	games, err := parseInput(contents)
	check(err)

	sum := 0
	for _, game := range games {
		sum += cheapestToWin(game)
	}
	fmt.Println(sum)
}
