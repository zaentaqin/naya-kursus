package main

import "fmt"

func main() {
	n := []int{1, 2, 4, 7, 1}
	k := 9

	for i := 1; len(n) > 0; i++ {
		fmt.Printf("Iterasi-%d\n", i)
		var newSlice []int
		for j := 0; j < len(n); j++ {
			if n[j] <= k {
				k -= n[j]
				newSlice = append(newSlice, n[j])
			}
		}
		fmt.Println(newSlice)
		n = newSlice
	}

	fmt.Println("Hasil:", len(n))
}
