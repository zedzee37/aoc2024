package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	target  int
	numbers []int
}

func concatNumbers(a, b int) int {
	numDigits := int(math.Log10(float64(b))) + 1
	aShifted := a * int(math.Pow10(numDigits))
	result := aShifted + b

	return result
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
		equation := make([]int, 0)

		target_str := numbers[0][:len(numbers[0])-1]
		target, err := strconv.Atoi(target_str)
		if err != nil {
			return nil, err
		}

		for _, num := range numbers[1:] {
			n, err := strconv.Atoi(num)

			if err != nil {
				return nil, err
			}

			equation = append(equation, n)
		}

		eq := Equation{target: target, numbers: equation}
		equations = append(equations, eq)
	}

	return equations, nil
}

func solveEquations(equations []Equation) int {
	sum := 0
	for _, equation := range equations {
		if doesEquationWork(equation, 0, equation.numbers[0]) {
			sum += equation.target
		}
	}
	return sum
}

func doesEquationWork(equation Equation, i int, total int) bool {
	if i >= len(equation.numbers)-1 {
		return total == equation.target
	}

	return doesEquationWork(equation, i+1, total+equation.numbers[i+1]) ||
		doesEquationWork(equation, i+1, total*equation.numbers[i+1]) ||
		doesEquationWork(equation, i+1, concatNumbers(total, equation.numbers[i+1]))
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
