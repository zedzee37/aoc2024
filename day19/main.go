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
	Leaves   LinkedListNode[string]
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

	poppedLeaves := make([]string, 0)

	for i := len(patternNode.Leaves) - 1; i >= 0; i-- {
		leaf := patternNode.Leaves[i]
		if len(leaf) > depth && leaf[depth] == pattern[depth] {
			poppedLeaves = append(poppedLeaves, leaf)

			startSlice := patternNode.Leaves[0:i]
			endSlice := patternNode.Leaves[i+1]
			patternNode.Leaves = append(startSlice, endSlice)
		}
	}

	if len(poppedLeaves) > 0 {
		newNode := NewPatternNode()
		newNode.Leaves = append(newNode.Leaves, pattern)
		newNode.Leaves = append(newNode.Leaves, poppedLeaves...)
		patternNode.Children[pattern[depth]] = newNode
		return
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
	headNode.InsertPattern("arg")
	headNode.InsertPattern("ar")

	printPatternNode(headNode)
}
