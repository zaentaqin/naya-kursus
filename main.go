package main

import (
	"fmt"
	"math"
)

func OtherAngle(a int, b int) int {
	ress := a*a + b*b
	ress_2 := math.Sqrt(float64(ress))

	return int(ress_2)
}

func main() {
	fmt.Println(OtherAngle(30, 60))

}
