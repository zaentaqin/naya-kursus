package main

import (
	"fmt"
	"strings"
)

func FindShort(s string) int {
	words := strings.Fields(s)
	minLength := len(words[0])
	for _, word := range words {
		if len(word) < minLength {
			minLength = len(word)
		}
	}
	return minLength
}

func main() {
	fmt.Println(FindShort("Muhammad Zaenal Muttaqin"))

}
