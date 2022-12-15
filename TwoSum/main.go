package main

import (
	"fmt"
)

func main() {
	var nums = []int{3, 2, 4}
	var targer int
	_, _ = fmt.Scan(&targer)
	nums1 := twoSum(nums, targer)
	fmt.Println(nums1)
}

func twoSum(nums []int, target int) []int {
	var num []int
	for j := 0; j < len(nums); j++ {
		for i := 0; i < len(nums); i++ {
			if i != j {
				if nums[i]+nums[j] == target {
					num = append(num, j)
					num = append(num, i)
					return num
				}
			}
		}
	}
	return num
}
