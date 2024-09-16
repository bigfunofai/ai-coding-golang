package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[rune]int)
	maxLength, start := 0, 0

	for i, c := range s {
		if idx, ok := charIndex[c]; ok && idx >= start {
			start = idx + 1
		}
		charIndex[c] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
}