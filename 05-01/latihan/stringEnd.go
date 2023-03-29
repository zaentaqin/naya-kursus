package main

import "fmt"

func solution(str, ending string) bool {
	a := len(str)
	b := len(ending)

	if a < b {
		return false
	}
	return str[a-b:] == ending
}

func main() {
	fmt.Println(solution("abc", "bc"))
}
