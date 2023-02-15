package main

import "fmt"

func Number(busStops [][2]int) int {
	result := 0

	for _, res := range busStops {
		result += res[0] - res[1]
	}
	return result
}

func main() {
	a := []int{3, 5}
	fmt.Println(a)
}
