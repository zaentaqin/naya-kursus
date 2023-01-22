package main

import "fmt"

func multiplyArrayElements(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i-1] * arr[i]
		if arr[i] > 30 {
			break
		}
	}
	return arr
}

func main() {
	arr := []int{2, 3, 4, 5}
	fmt.Println("Original array: ", arr)
	result := multiplyArrayElements(arr)
	fmt.Println("Multiplied array: ", result)
}
