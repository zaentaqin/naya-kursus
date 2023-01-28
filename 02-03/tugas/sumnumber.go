package main

import "fmt"

func GetSum(a, b int) int {
	if a == b {
		return a
	}
	var sum int
	if a < b {
		for i := a; i <= b; i++ {
			sum += i
		}
	} else {
		for i := b; i <= a; i++ {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(GetSum(1, 1))
}
