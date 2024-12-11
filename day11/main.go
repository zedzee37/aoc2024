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

func blink(stones []int) []int {
	new_stones := []int{}
	for _, stone := range stones {
		if stone == 0 {
			new_stones = append(new_stones, 1)
		} else {
			num1, num2, err := splitNumber(stone)

			if err != nil {
				new_stones = append(new_stones, stone*2024)
			} else {
				new_stones = append(new_stones, num1)
				new_stones = append(new_stones, num2)
			}
		}
	}
	return new_stones
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	text := string(data)
	numbers, err := parseInput(text)
	check(err)

	for i := 0; i < 25; i++ {
		numbers = blink(numbers)
		fmt.Println(numbers)
	}

	fmt.Println(len(numbers))
}
