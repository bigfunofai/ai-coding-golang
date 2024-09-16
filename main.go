package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    merged := mergeArrays(nums1, nums2)
    n := len(merged)
    if n%2 == 0 {
        return float64(merged[n/2-1]+merged[n/2]) / 2
    }
    return float64(merged[n/2])
}

func mergeArrays(nums1 []int, nums2 []int) []int {
    i, j := 0, 0
    var merged []int

    for i < len(nums1) && j < len(nums2) {
        if nums1[i] < nums2[j] {
            merged = append(merged, nums1[i])
            i++
        } else {
            merged = append(merged, nums2[j])
            j++
        }
    }

    for i < len(nums1) {
        merged = append(merged, nums1[i])
        i++
    }

    for j < len(nums2) {
        merged = append(merged, nums2[j])
        j++
    }

    return merged
}

func main() {
    nums1 := []int{1, 3}
    nums2 := []int{2}
    fmt.Println(findMedianSortedArrays(nums1, nums2)) // Output: 2

    nums1 = []int{1, 2}
    nums2 = []int{3, 4}
    fmt.Println(findMedianSortedArrays(nums1, nums2)) // Output: 2.5
}