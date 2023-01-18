package main

import "fmt"

func makenegative(x int) int {
	if x > 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(makenegative(3))
	fmt.Println(makenegative(-3))

}
