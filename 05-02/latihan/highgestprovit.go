package main

import "fmt"

func Minmax(arr []int) [2]int {
	result := [2]int{}

	max := arr[0]
	min := arr[0]

	for _, a := range arr {
		if max < a {
			max = a
		}
		if min > a {
			min = a
		}
	}
	result[0] = min
	result[1] = max

	return result
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(Minmax(a))

	// b := []int{2334454, 5}
	// fmt.Println(Minmax(b))
}
