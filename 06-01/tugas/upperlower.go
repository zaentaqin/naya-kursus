package main

import (
	"fmt"
	"strings"
)

func solve(str string) string {
	// Your code here and happy coding!
	upp := 0
	low := 0

	for _, r := range str {
		if r < 'a' {
			upp++
			//fmt.Println(upp)
		} else {
			low++
		}
		//fmt.Println(low)
	}

	if low >= upp {
		return strings.ToLower(str)
	}
	return strings.ToUpper(str)
}

func main() {
	fmt.Println(solve("a"))
}
