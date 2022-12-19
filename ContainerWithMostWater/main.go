package main

import "fmt"

func main() {
	heigth := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	hei := maxArea(heigth)
	fmt.Println(hei)
}

func maxArea(height []int) int {
	var max, max1 int
	for i := 0; i < len(height); i++ {
		for j := i; j < len(height); j++ {
			if height[j] < height[i] {
				num := height[i] - (height[i] - height[j])
				max = num * j
				if max > max1 {
					max1 = max
				}
			} else {
				num := height[j] - (height[j] - height[i])
				max = num * j
				if max > max1 {
					max1 = max
				}
			}
		}
	}
	return max1
}
