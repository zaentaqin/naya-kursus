package main

import "fmt"

func isPowerOfTwo(n int) bool {

	for n != 1 {
		fmt.Println(n)
		if n%2 > 0 || n <= 0 {
			return false
		}
		n /= 2
	}
	return n == 1
}

func main() {
	fmt.Println(isPowerOfTwo(6))
}
