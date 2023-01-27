package main

import (
	"fmt"
	"strings"
)

func Accum(s string) string {
	var result []string
	for i, r := range s {
		result = append(result, strings.ToUpper(string(r))+strings.Repeat(strings.ToLower(string(r)), i))
	}
	return strings.Join(result, "-")
}

func main() {
	fmt.Println(Accum("ABC"))
	fmt.Println(Accum("RqaEzty"))
}
