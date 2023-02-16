package main

import "fmt"

func BreakingChocolatee(n, m int) int {
	if n < 1 || m < 1 {
		return 0
	}
	return (n * m) - 1
}

func main() {
	a := BreakingChocolatee(5, 5)
	fmt.Println(a)

	b := BreakingChocolatee(1, 1)
	fmt.Println(b)
}
