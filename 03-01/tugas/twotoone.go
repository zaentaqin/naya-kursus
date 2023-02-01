package main

import (
	"fmt"
	"runtime"
	"strings"
)

func TwoToOne(a string, b string) string {
	var c []string
	str := a + b
	for r := 'a'; r <= 'z'; r++ {
		if strings.ContainsRune(str, r) {
			c = append(c, string(r))
			fmt.Println(c)
		}
	}
	return strings.Join(c, "")
}

func main() {
	fmt.Println(TwoToOne("qAwertyuioplkjh", "zaxmjkqwdfhiow"))
	fmt.Println(runtime.NumGoroutine())
}
