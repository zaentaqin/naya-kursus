package main

import "fmt"

func oddnumber(num int) int {
	if num%2 == 1 {
		return (num + 1) / 2
	}
	return num / 2
}

func main() {
	fmt.Println(oddnumber(3))
}
