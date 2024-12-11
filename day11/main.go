package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput(input string) ([]int, error) {
	numbers := make([]int, 0)

	numberStrings := strings.Split(input, " ")
	for _, str := range numberStrings {
		converted, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, converted)
	}

	return numbers, nil
}

func splitNumber(n int) (int, int, error) {
	// Convert number to string
	numStr := strconv.Itoa(n)
	length := len(numStr)

	// Handle odd-length numbers
	if length%2 != 0 {
		return 0, 0, fmt.Errorf("number has an odd number of digits")
	}

	// Split into two halves
	half := length / 2
	firstHalf, err1 := strconv.Atoi(numStr[:half])
	secondHalf, err2 := strconv.Atoi(numStr[half:])

	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("failed to convert split parts to integers")
	}

	return firstHalf, secondHalf, nil
}

func blink(stones []int, memo *map[int][]int) []int {
	new_stones := []int{}
	for _, stone := range stones {
		nums, ok := (*memo)[stone]

		if ok {
			for _, num := range nums {
				new_stones = append(new_stones, num)
			}
		} else {
			if stone == 0 {
				new_stones = append(new_stones, 1)
				(*memo)[0] = []int{1}
			} else {
				num1, num2, err := splitNumber(stone)

				if err != nil {
					num := stone * 2024
					(*memo)[stone] = []int{num}
					new_stones = append(new_stones, num)
				} else {
					(*memo)[stone] = []int{num1, num2}
					new_stones = append(new_stones, num1)
					new_stones = append(new_stones, num2)
				}
			}
		}
	}
	return new_stones
}

func stoneCount(number int, iterations int, memo *map[[2]int]int) int {
	ct, ok := (*memo)[[2]int{number, iterations}]

	if ok {
		return ct
	}

	ret := 0
	if iterations == 0 {
		ret = 1
	} else if number == 0 {
		ret = stoneCount(1, iterations-1, memo)
	} else {
		num1, num2, err := splitNumber(number)

		if err != nil {
			ret = stoneCount(number*2024, iterations-1, memo)
		} else {
			ret = stoneCount(num1, iterations-1, memo) + stoneCount(num2, iterations-1, memo)
		}
	}

	(*memo)[[2]int{number, iterations}] = ret
	return ret
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	text := string(data)
	numbers, err := parseInput(text)
	check(err)

	memo := make(map[[2]int]int)
	sum := 0
	for _, i := range numbers {
		sum += stoneCount(i, 75, &memo)
	}
	fmt.Println(sum)
}
