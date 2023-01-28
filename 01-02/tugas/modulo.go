package main

import "fmt"

func isDivisible(n, x, y int) bool {
	if (n%x == 0) && (n%y == 0) {
		return true
	}
	return false
}

func main() {
	fmt.Println(isDivisible(12, 4, 3))
	fmt.Println(isDivisible(12, 4, 5))
	fmt.Println(isDivisible(12, 7, 5))
}
