package main

import "fmt"

func reduceArray(n []int, k int) int {
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ {
			if n[i]*n[j] <= k {
				n[i] *= n[j]
				n = append(n[:j], n[j+1:]...)
				fmt.Println("Iterasi-", i+1, n)
				j--
			}
		}
	}
	return len(n)
}

func main() {
	n := []int{1, 2, 4, 2, 2}
	k := 9
	fmt.Println("Hasil:", reduceArray(n, k))
}
