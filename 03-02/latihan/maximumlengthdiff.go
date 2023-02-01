package main

import "fmt"

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	minLengtha1 := len(a1[0])
	maxLengtha1 := len(a1[0])
	for _, word := range a1 {
		if len(word) < minLengtha1 {
			minLengtha1 = len(word)
		}
		if len(word) > maxLengtha1 {
			maxLengtha1 = len(word)
		}
	}
	minLengtha2 := len(a2[0])
	maxLengtha2 := len(a2[0])
	for _, word := range a2 {
		if len(word) < minLengtha2 {
			minLengtha2 = len(word)
		}
		if len(word) > maxLengtha2 {
			maxLengtha2 = len(word)
		}
	}

	x := maxLengtha1 - minLengtha2
	y := maxLengtha2 - minLengtha1
	if x > y {
		return x
	} else {
		return y
	}

}

func main() {
	s1 := []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
	s2 := []string{"bbbaaayddqbbrrrv"}

	b1 := []string{}
	b2 := []string{}

	fmt.Println(MxDifLg(s1, s2))
	fmt.Println(MxDifLg(b1, b2))
}
