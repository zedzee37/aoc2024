package main

import (
	"fmt"
	"os"
)

type PatternNode struct {
	Leaves   []string
	Children map[byte]*PatternNode
	Parent   *PatternNode
}

func NewPatternNode() *PatternNode {
	patternNode := new(PatternNode)

	patternNode.Children = make(map[byte]*PatternNode, 0)
	patternNode.Leaves = make([]string, 0)
	patternNode.Parent = nil

	return patternNode
}

func (patternNode *PatternNode) InsertPattern(pattern string) {
	patternNode.insertPattern(pattern, 0)
}

func (patternNode *PatternNode) insertPattern(pattern string, depth int) {
	if len(patternNode.Leaves) == 0 && len(patternNode.Children) == 0 {
		patternNode.Leaves = append(patternNode.Leaves, pattern)
		return
	}

	if len(patternNode.Children) != 0 {
		for key, node := range patternNode.Children {
			if pattern[depth] == key {
				node.insertPattern(pattern, depth+1)
				return
			}
		}
	}

	for i, leaf := range patternNode.Leaves {
		if leaf[depth] == pattern[depth] {
			patternNode.Leaves = append(patternNode.Leaves[0:i], patternNode.Leaves[i+1:len(patternNode.Leaves)]...)

			newNode := NewPatternNode()
			newNode.Leaves = append(newNode.Leaves, leaf, pattern)
			patternNode.Children[leaf[depth]] = newNode

			return
		}
	}

	patternNode.Leaves = append(patternNode.Leaves, pattern)
}

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

func printPatternNode(patternNode *PatternNode) {
	println("Leaves: ")
	for _, leaf := range patternNode.Leaves {
		println(leaf)
	}
	println()

	for key, child := range patternNode.Children {
		fmt.Println("Printing leaf with key: {} --", string(key))
		println()
		printPatternNode(child)
	}
}

func main() {
	// _, _, err := parseInput("input.txt")

	// if err != nil {
	// 	panic(err.Error())
	// }

	headNode := NewPatternNode()
	headNode.InsertPattern("gug")
	headNode.InsertPattern("gork")
	headNode.InsertPattern("gork")
	headNode.InsertPattern("gag")
	headNode.InsertPattern("arg")

	printPatternNode(headNode)
}
