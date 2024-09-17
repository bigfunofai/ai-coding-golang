package main

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"", 0},
		{"(", 0},
		{")", 0},
		{"()", 2},
		{"(()", 2},
		{")()())", 4},
		{"()(()", 2},
		{"()(())", 6},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := longestValidParentheses(tt.s); got != tt.want {
				t.Errorf("longestValidParentheses(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}