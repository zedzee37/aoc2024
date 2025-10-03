package main

import (
	"fmt"
	"os"
)

type LinkedListNode[T any] struct {
	value T
	next  *LinkedListNode[T]
}

type PatternNode struct {
	Previous *PatternNode
	Exists   bool
	Children map[byte]*PatternNode
	Parent   *PatternNode
}

func NewPatternNode() *PatternNode {
	patternNode := new(PatternNode)

	patternNode.Children = make(map[byte]*PatternNode, 0)
	patternNode.Exists = false
	patternNode.Parent = nil
	patternNode.Previous = nil

	return patternNode
}

func (patternNode *PatternNode) InsertPattern(pattern string) {
	if len(pattern) == 0 {
		patternNode.Exists = true
		return
	}

	for key, node := range patternNode.Children {
		if pattern[0] == key {
			node.InsertPattern(pattern[1:])
			return
		}
	}

	newNode := NewPatternNode()
	newNode.InsertPattern(pattern[1:])
	newNode.Previous = patternNode
	patternNode.Children[pattern[0]] = newNode
}

// func isPatternPossible(availablePatterns []string, pattern string) bool {

// }

func parseInput(fp string) (*PatternNode, []string, error) {
	dat, err := os.ReadFile(fp)
	if err != nil {
		return nil, nil, err
	}

	fileContents := string(dat)

	patterns := NewPatternNode()
	start := 0
	for i, ch := range fileContents {
		if ch == '\n' {
			patterns.InsertPattern(fileContents[start:i])
			start = i
			break
		} else if ch == ' ' {
			start = i + 1
		} else if ch == ',' {
			patterns.InsertPattern(fileContents[start : i-1])
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

	return patterns, targetPatterns, nil
}

func printPatternNode(patternNode *PatternNode, depth int) {
	for key, child := range patternNode.Children {
		padding := ""
		for _ = range depth {
			padding += "  >  "
		}

		fmt.Printf("%s Printing leaf with key: %s --\n", padding, string(key))

		existsString := "false"
		if child.Exists {
			existsString = "true"
		}

		fmt.Printf("%s Exists: %s\n\n", padding, existsString)
		printPatternNode(child, depth+1)
	}
}

func main() {
	patterns, _, err := parseInput("input.txt")

	if err != nil {
		panic(err.Error())
	}

	printPatternNode(patterns, 0)
}
