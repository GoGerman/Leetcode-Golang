package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums1, nums2 []int = []int{1, 3, 4, 5}, []int{2, 6, 7, 8}
	num3 := findMedianSortedArrays(nums1, nums2)
	fmt.Println(num3)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var fl float64
	var nums3 []int
	for i := 0; i < len(nums1); i++ {
		nums3 = append(nums3, nums1[i])
	}
	for i := 0; i < len(nums2); i++ {
		nums3 = append(nums3, nums2[i])
	}
	sort.Ints(nums3)
	if len(nums3)%2 != 0 {
		fl = float64(nums3[len(nums3)/2])
		return fl
	} else {
		fl = (float64(nums3[len(nums3)/2]) + float64(nums3[len(nums3)/2-1])) / 2
		return fl
	}
}
