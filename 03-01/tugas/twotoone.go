package main

import (
	"fmt"
	"strings"
)

func TwoToOne(a string, b string) string {
	var c []string
	str := a + b
	for r := 'z'; r >= 'a'; r-- {
		if strings.ContainsRune(str, r) {
			c = append(c, string(r))
		}
	}
	return strings.Join(c, " ")
}

func main() {
	fmt.Println(TwoToOne("qwertyuioplkjh", "zxmjkqwdfhiow"))
}
