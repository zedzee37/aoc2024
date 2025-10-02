package main

import (
	"fmt"
	"os"
)

// func isPatternPossible(availablePatterns []string, pattern string) bool {

// }

func parseInput(fp string) ([]string, []string, error) {
	dat, err := os.ReadFile(fp)
	if err != nil {
		return nil, nil, err
	}

	fileContents := string(dat)

	availablePatterns := make([]string, 0)
	start := 0
	for i, ch := range fileContents {
		if ch == '\n' {
			availablePatterns = append(availablePatterns, fileContents[start:i])
			start = i
			break
		} else if ch == ' ' {
			start = i + 1
		} else if ch == ',' {
			availablePatterns = append(availablePatterns, fileContents[start:i-1])
			start = i
		}
	}

	endSlice := fileContents[start:]
	targetPatterns := make([]string, 0)
	start = 0
	for i, ch := range endSlice {
		if ch != '\n' {
			continue
		}

		if i-1 <= 0 {
			continue
		}
		subStr := endSlice[start : i-1]

		start = i + 1
		if len(subStr) <= 0 {
			continue
		}

		targetPatterns = append(targetPatterns, subStr)
	}

	return availablePatterns, targetPatterns, nil
}

func main() {
	_, targetPatterns, err := parseInput("input.txt")

	if err != nil {
		panic(err.Error())
	}

	// for _, pattern := range availablePatterns {
	// 	fmt.Println(pattern)
	// }

	fmt.Println()
	for _, pattern := range targetPatterns {
		fmt.Println(pattern)
	}
}
