package main

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"au", 2},
		{"dvdf", 3},
	}

	for _, test := range tests {
		result := lengthOfLongestSubstring(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %d, but got %d", test.input, test.expected, result)
		}
	}
}