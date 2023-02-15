package main

import "fmt"

func NbYear(p0 int, percent float64, aug int, p int) (years int) {
	for p0 < p {
		p0 = p0 + int(float64(p0)*percent/100) + aug
		years += 1
	}
	return
}

func main() {
	fmt.Println(NbYear(1500, 5, 100, 5000))
	fmt.Println(NbYear(1500000, 2.5, 10000, 2000000))
}
