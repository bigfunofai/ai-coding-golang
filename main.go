package main

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	stack, maxLen := []int{-1}, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxLen = max(maxLen, i-stack[len(stack)-1])
			}
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestValidParentheses("(()"))    // should return 2
	fmt.Println(longestValidParentheses(")()())")) // should return 4
}