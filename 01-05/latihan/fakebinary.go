package main

import "fmt"

func FakeBin(x string) string {
	var result string
	for _, r := range x {
		if r < '5' {
			result += "0"
		} else {
			result += "1"
		}
	}
	return result
}

func main() {
	fmt.Println(FakeBin("45385593107843568"))
}
