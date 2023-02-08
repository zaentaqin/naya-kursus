package main

import (
	"fmt"
)

func romanToInt(s string) int {

	bil := 0

	for i := 0; i < len(s); i++ {

		if s[i] == 'M' {
			bil += 1000
		}
		if s[i] == 'D' {
			bil += 500
		}
		if s[i] == 'C' {

			if i+1 < len(s) && s[i+1] == 'D' {
				bil += 400
				i++

				continue
			}

			if i+1 < len(s) && s[i+1] == 'M' {
				bil += 900
				i++

				continue
			}
			bil += 100
		}
		if s[i] == 'L' {
			bil += 50
		}
		if s[i] == 'X' {

			if i+1 < len(s) && s[i+1] == 'L' {
				bil += 40
				i++

				continue
			}

			if i+1 < len(s) && s[i+1] == 'C' {
				bil += 90
				i++

				continue
			}
			bil += 10
		}
		if s[i] == 'V' {
			bil += 5
		}
		if s[i] == 'I' {

			if i+1 < len(s) && s[i+1] == 'V' {
				bil += 4
				i++

				continue
			}

			if i+1 < len(s) && s[i+1] == 'X' {
				bil += 9
				i++

				continue
			}
			bil += 1
		}
	}
	return bil
}

func main() {
	s := "XIX"
	var t string
	t = "MMII"
	fmt.Println(romanToInt(s))
	fmt.Println(romanToInt(t))
}
