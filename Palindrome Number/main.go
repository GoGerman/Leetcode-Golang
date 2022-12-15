package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	_, _ = fmt.Scan(&num)
	var bl = isPalindrome(num)
	fmt.Println(bl)
}

func isPalindrome(x int) bool {
	var str string
	var num int
	if x < 0 {
		return false
	}
	for i := x; i != 0; i /= 10 {
		num = i % 10
		str += strconv.Itoa(num)
	}
	num1, _ := strconv.Atoi(str)
	if x == num1 {
		return true
	}
	return false
}
