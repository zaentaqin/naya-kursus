package main

import "fmt"

func century(year int) int {
	if year%100 == 0 {
		return year / 100
	}
	return (year / 100) + 1
}

func main() {
	fmt.Println(century(2023))
}
