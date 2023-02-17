package main

import (
	"fmt"
)

func Gimme(array [3]int) (ress int) {
	arr := array
	fmt.Println(arr)
	fmt.Println(arr[1:1])
	// sort.Ints(arr[:0])
	// // ress := 0
	// for a, x := range array {
	// 	if x == arr[1] {
	// 		// return a
	// 		ress = a
	// 		break
	// 	}
	// }

	return
}

func main() {
	fmt.Println(Gimme([3]int{2, 3, 1}))
	// fmt.Println(Gimme([3]int{3, 2, 1}))
	// fmt.Println(Gimme([3]int{-1, -3, -2}))
}
