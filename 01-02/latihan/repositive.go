package main

import "fmt"

func possitive(numbers []int) int {
	result := 0
	for _, number := range numbers {
		if number > 0 {
			result += number
		}
	}
	return result
}

func main() {

	// number := []int{1, 5, -9, 3, 2}
	fmt.Println(possitive([]int{1, 5, -9, 3, 2}))
}
