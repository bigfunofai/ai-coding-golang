package main

import "testing"

func TestMaxSubArray(t *testing.T) {
    tests := []struct {
        nums     []int
        expected int
    }{
        {[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
        {[]int{1}, 1},
        {[]int{0}, 0},
        {[]int{-1}, -1},
        {[]int{-100000}, -100000},
        {[]int{5, 4, -1, 7, 8}, 23},
    }

    for _, tt := range tests {
        actual := maxSubArray(tt.nums)
        if actual != tt.expected {
            t.Errorf("maxSubArray(%v) = %v; expected %v", tt.nums, actual, tt.expected)
        }
    }
}