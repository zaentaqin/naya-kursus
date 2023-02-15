package main

import "fmt"

func solution(str, ending string) bool {
	a := len(str)
	b := len(ending)

	if a < b {
		return false
	}
	// Your code here!
	return str[a-b:] == ending

}

func main() {
	fmt.Println(solution("abc", "bc"))
}
