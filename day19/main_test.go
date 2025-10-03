package main

import (
	"testing"
)

// Helper to build Trie from a list of strings
func buildTrie(patterns []string) *TrieNode {
	root := NewTrieNode()
	for _, pattern := range patterns {
		root.Insert(pattern)
	}
	return root
}

func TestTrieNode_Insert(t *testing.T) {
	root := NewTrieNode()
	root.Insert("ab")
	root.Insert("abc")
	root.Insert("bcd")
	root.Insert("a")

	// Check root has correct children
	if _, ok := root.children['a']; !ok {
		t.Errorf("Expected child 'a' to exist in root")
	}
	if _, ok := root.children['b']; !ok {
		t.Errorf("Expected child 'b' to exist in root")
	}

	// Check specific node path
	node := root.children['a']
	if !node.Exists {
		t.Errorf("Expected node 'a' to have Exists = true")
	}

	if _, ok := node.children['b']; !ok {
		t.Errorf("Expected child 'b' under node 'a'")
	}
}

func TestIsPatternPossible(t *testing.T) {
	trie := buildTrie([]string{"a", "ab", "cd", "ef"})

	tests := []struct {
		input    string
		expected bool
	}{
		{"ab", true},     // "ab" is directly in trie
		{"abcdef", true}, // "ab" + "cd" + "ef"
		{"abc", false},   // "ab" + "c" not found, but "a" + "bc" is not possible -> "ab" + "c" => c not found
		{"abcd", true},   // "ab" + "cd"
		{"abcde", false},
		{"cdef", true}, // "cd" + "ef"
		{"gh", false},  // not in trie
		{"", true},     // empty string always returns true
	}

	for _, test := range tests {
		result := isPatternPossible(trie, test.input)
		if result != test.expected {
			t.Errorf("isPatternPossible(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
