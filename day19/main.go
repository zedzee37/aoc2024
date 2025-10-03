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
	if !ok {
		nextNode = NewTrieNode()
		node.children[ch] = nextNode
	}
	nextNode.Insert(str[1:])
}

func isPatternPossible(availablePatterns *TrieNode, pattern string) bool {
	visited := make(map[string]bool, 0)
	return isPatternPossibleHelper(availablePatterns, pattern, visited)
}

func isPatternPossibleHelper(availablePatterns *TrieNode, pattern string, visited map[string]bool) bool {
	if len(pattern) == 0 {
		return true
	}

	_, exists := visited[pattern]
	if exists {
		return false
	}

	visited[pattern] = true

	node := availablePatterns
	i := 0
	for i < len(pattern) {
		ch := pattern[i]

		newNode, exists := node.children[ch]
		if !exists {
			return false
		}

		if newNode.Exists && isPatternPossibleHelper(availablePatterns, pattern[i+1:], visited) {
			return true
		}

		node = newNode
		i++
	}
	return false
}

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
			start = i + 1
			break
		} else if ch == ' ' {
			start = i + 1
		} else if ch == ',' {
			patterns.Insert(fileContents[start:i])
			start = i + 1
		}
	}

	endSlice := fileContents[start:]
	targetPatterns := make([]string, 0)
	start = 0
	for i, ch := range endSlice {
		if ch != '\n' {
			continue
		}

		if i <= 0 {
			continue
		}
		subStr := endSlice[start:i]

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
	patterns, targetPatterns, err := parseInput("input.txt")

	if err != nil {
		panic(err.Error())
	}

	count := 0
	for _, pattern := range targetPatterns {
		if isPatternPossible(patterns, pattern) {
			count++
		}
	}

	println(count)
}
