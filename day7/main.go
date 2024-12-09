package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	target int
	leaves []Leaf
}

type Leaf struct {
	left  int
	right int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseLines(lines []string) ([]Equation, error) {
	equations := make([]Equation, 0)

	for _, line := range lines {
		numbers := strings.Split(line, " ")
		equation := make([]Leaf, 0)

		target_str := numbers[0][:len(numbers[0])-1]
		target, err := strconv.Atoi(target_str)
		if err != nil {
			return nil, err
		}

		prev := -1
		for _, num := range numbers[1:] {
			n, err := strconv.Atoi(num)

			if err != nil {
				return nil, err
			}

			if prev != -1 {
				equation = append(equation, Leaf{left: prev, right: n})
				prev = -1
			} else {
				prev = n
			}
		}

		if prev != -1 {
			equation = append(equation, Leaf{left: prev, right: -1})
		}

		eq := Equation{target: target, leaves: equation}
		equations = append(equations, eq)
	}

	return equations, nil
}

func solveEquations(equations []Equation) int {
	
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	text := string(data)
	lines := strings.Split(text, "\n")

	equations, err := parseLines(lines)
	check(err)

	count := solveEquations(equations)
	fmt.Println(count)
}
