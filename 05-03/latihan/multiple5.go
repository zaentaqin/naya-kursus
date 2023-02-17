package main

import "fmt"

func RoundToNext5(n int) int {

	if n%5 == 0 {
		return n
	}

	if n > 0 {
		return n/5*5 + 5
	}

	return n / 5 * 5
	/*
		 	if n < 0 {
				return n / 5 * 5
			}
			return 0
	*/
}

func main() {

	fmt.Println(RoundToNext5(-5))
	fmt.Println(RoundToNext5(15))
}
