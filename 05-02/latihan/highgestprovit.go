package main

import (
	"fmt"
	"sort"
)

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

func Soort(brr []int) [2]int {
	ress := [2]int{}
	sort.Ints(brr)

	ress[0] = brr[0]
	ress[1] = brr[len(brr)-1]

	return ress

}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(Minmax(a))
	fmt.Println(Soort(a))

	b := []int{2334454, 5}
	fmt.Println(Minmax(b))
	fmt.Println(Soort(b))
}
