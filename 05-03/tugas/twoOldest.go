package main

import (
	"fmt"
	"sort"
)

func TwoOldestAges(ages []int) [2]int {
	sorted := ages
	sort.Ints(sorted[:])
	fmt.Println(sorted)

	return [2]int{sorted[len(ages)-2], sorted[len(ages)-1]}
}

func main() {
	fmt.Println(TwoOldestAges([]int{6, 5, 83, 5, 3, 18}))
}
