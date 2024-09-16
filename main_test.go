package main

import "testing"

func TestIsPalindrome(t *testing.T) {
    tests := []struct {
        input    string
        expected bool
    }{
        {"Madam", true},
        {"Hello", false},
        {"Racecar", true},
        {"World", false},
        {"Level", true},
        {"A man a plan a canal Panama", false},
    }

    for _, tt := range tests {
        result := isPalindrome(tt.input)
        if result != tt.expected {
            t.Errorf("isPalindrome(%s) = %v; want %v", tt.input, result, tt.expected)
        }
    }
}