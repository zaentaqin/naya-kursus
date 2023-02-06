package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	fmt.Println(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(containsDuplicate([]int{2, 1, 4, 5}))
	fmt.Println(containsDuplicate([]int{2, 1, 1, 5}))

}
