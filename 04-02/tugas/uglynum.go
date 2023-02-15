package main

import "fmt"

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	prima := [3]int{2, 3, 5}

	for _, x := range prima {
		for n%x == 0 {
			n = n / x
			// fmt.Println("Value x :", x)
			// fmt.Println("Value n :", n)
		}
	}
	return n == 1
}

func main() {
	fmt.Println(isUgly(30))
}
