package main

import (
	"fmt"
	"os"
)

type TrieNode struct {
	Exists   bool
	children map[byte]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Exists:   false,
		children: make(map[byte]*TrieNode, 0),
	}
}

func (node *TrieNode) Insert(str string) {
	if len(str) == 0 {
		node.Exists = true
		return
	}

	ch := str[0]
	nextNode, ok := node.children[ch]
	if ok {
		nextNode.Insert(str[1:])
		return
	}

	newNode := NewTrieNode()
	newNode.Insert(str[1:])
	node.children[ch] = newNode
}

// func isPatternPossible(availablePatterns []string, pattern string) bool {

// }

func parseInput(fp string) (*TrieNode, []string, error) {
	dat, err := os.ReadFile(fp)
	if err != nil {
		return nil, nil, err
	}

	fileContents := string(dat)

	patterns := NewTrieNode()
	start := 0
	for i, ch := range fileContents {
		if ch == '\n' {
			patterns.Insert(fileContents[start:i])
			start = i
			break
		} else if ch == ' ' {
			start = i + 1
		} else if ch == ',' {
			patterns.Insert(fileContents[start : i-1])
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

func printPatternNode(patternNode *TrieNode, depth int) {
	for key, child := range patternNode.children {
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
	// patterns, _, err := parseInput("input.txt")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// printPatternNode(patterns, 0)

	headNode := NewTrieNode()
	headNode.Insert("gu")
	headNode.Insert("gug")
	headNode.Insert("gugg")
	headNode.Insert("arg")
	headNode.Insert("ar")
	headNode.Insert("avg")
	headNode.Insert("arvv")
	headNode.Insert("b")
	printPatternNode(headNode, 0)
}
