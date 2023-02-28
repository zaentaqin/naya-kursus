package main

import "fmt"

func GroupArray(arr []int, target int) (ress [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i]+arr[j] == target && i != j {
				ress = append(ress, []int{arr[i], arr[j]})
			}
		}
	}
	return
}

func main() {
	fmt.Println(GroupArray([]int{1, 2, 3, 4, 5}, 6))
}
