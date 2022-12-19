package main

import "fmt"

func main() {
	heigth := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	hei := maxArea(heigth)
	fmt.Println(hei)
}

func maxArea(height []int) int {
	w, max, left, right := 0, 0, 0, len(height)-1
	for left < right {
		if hl, hr, dist := height[left], height[right], right-left; hl > hr {
			w = dist * hr
			right--
		} else {
			w = dist * hl
			left++
		}
		if w > max {
			max = w
		}
	}
	return max
}
