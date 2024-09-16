package main

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
    testCases := []struct {
        nums1   []int
        nums2   []int
        want    float64
    }{
        {
            nums1: []int{1, 3},
            nums2: []int{2},
            want:  2,
        },
        {
            nums1: []int{1, 2},
            nums2: []int{3, 4},
            want:  2.5,
        },
        {
            nums1: []int{0, 0},
            nums2: []int{0, 0},
            want:  0,
        },
        {
            nums1: []int{},
            nums2: []int{1},
            want:  1,
        },
        {
            nums1: []int{2},
            nums2: []int{},
            want:  2,
        },
    }

    for _, tc := range testCases {
        got := findMedianSortedArrays(tc.nums1, tc.nums2)
        if got != tc.want {
            t.Errorf("findMedianSortedArrays(%v, %v) = %v; want %v", tc.nums1, tc.nums2, got, tc.want)
        }
    }
}