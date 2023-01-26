package main

import "fmt"

func CountBy(x, n int) []int {
	var res []int
	for i := 1; i <= n; i++ {
		res = append(res, x*i)
	}
	return res
}

func main() {
	fmt.Println(CountBy(1, 5))
	fmt.Println(CountBy(5, 10))
}
