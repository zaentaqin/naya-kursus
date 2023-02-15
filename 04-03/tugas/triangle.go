package main

import "fmt"

func IsTriangle(a, b, c int) bool {
	return a+b > c && b+c > a && a+c > b
}

func main() {
	fmt.Println(IsTriangle(5, 1, 2))
	fmt.Println(IsTriangle(4, 2, 3))
}
