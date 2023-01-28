package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Printf("Nilai a = %v dan nilai b = %v\n", a, b)
	fmt.Println("Swab")

	b, a = a, b

	fmt.Printf("Nilai a = %d dan nilai b = %d\n", a, b)
}
