package main

import (
	"os"
)

func parseInput(fp string) ([]string, []string, error) {
	dat, err := os.ReadFile(fp)
	if err != nil {
		return nil, nil, err
	}

	fileContents := string(dat)

	availablePatterns := make([]string, 0)
	start := 0
	for i, ch := range fileContents {
		switch ch {
		case '\n':
			availablePatterns = append(availablePatterns, fileContents[start:i-1])
			start = i
			break
		case ' ':
			availablePatterns = append(availablePatterns, fileContents[start:i-1])
			start = i
		case ',':
			start = i
		default:
			continue
		}
	}

	endSlice := fileContents[start:]
	targetPatterns := make([]string, 0)
	start = 0
	for i, ch := range endSlice {
		if ch != '\n' {
			continue
		}

		subStr := endSlice[start : i-1]

		start = 0
		if len(subStr) <= 0 {
			continue
		}

		targetPatterns = append(targetPatterns, subStr)
	}

	return availablePatterns, targetPatterns, nil
}

func main() {

}
