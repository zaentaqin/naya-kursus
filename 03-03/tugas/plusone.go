package main

import "fmt"

func plusOne(digits []int) []int {

	val := len(digits)
	for i := val - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits

		}
		digits[i] = 0

	}
	digits = append([]int{1}, digits...)
	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{1, 9}))
}
