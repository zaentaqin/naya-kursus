package main

import (
	"fmt"
)

func addDigits(num int) int {
	sum := 0
	for num >= 10 {
		for num > 0 {
			sum += num % 10
			num = num / 10
		}
		num = num + sum
		//fmt.Println(num)
	}
	return num
}

func main() {

	fmt.Println(addDigits(29))
}
