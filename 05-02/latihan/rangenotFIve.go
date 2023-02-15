package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DontGiveMeFive(start int, end int) int {
	ress := 0

	for x := start; x <= end; x++ {
		angka := strconv.Itoa(x)
		if !strings.Contains(angka, "5") {
			ress += 1
		}
		fmt.Println(ress)
	}
	return ress
}

func main() {
	DontGiveMeFive(1, 9)

	fmt.Println(DontGiveMeFive(1, 9)) //1,2,3,4,5,6,7,8,9
}
