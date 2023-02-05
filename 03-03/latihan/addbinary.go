package main

import "fmt"

func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	if la < lb {
		a, b = b, a
		la, lb = lb, la
	}
	res := make([]byte, la+1)
	carry := byte(0)
	for i := 0; i < lb; i++ {
		res[la-i] = a[la-i-1] + b[lb-i-1] + carry - '0'
		if res[la-i] > '1' {
			res[la-i] -= 2
			carry = 1
		} else {
			carry = 0
		}
	}
	for i := lb; i < la; i++ {
		res[la-i] = a[la-i-1] + carry
		if res[la-i] > '1' {
			res[la-i] -= 2
			carry = 1
		} else {
			carry = 0
		}
	}
	if carry == 1 {
		res[0] = '1'
		return string(res)
	}
	return string(res[1:])
}

func main() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))

	a = "1010"
	b = "1011"
	fmt.Println(addBinary(a, b))
}
