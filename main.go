package main

import (
    "fmt"
    "strings"
)

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    left := 0
    right := len(s) - 1
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}

func main() {
    testStrings := []string{"Madam", "Hello", "Racecar", "World", "Level"}
    for _, str := range testStrings {
        fmt.Printf("%s: %v\n", str, isPalindrome(str))
    }
}