package main

import "fmt"

func hero(bullets, dragon int) bool {
	return bullets >= 2*dragon
}
func main() {
	fmt.Println(hero(100, 60))
	fmt.Println(hero(100, 40))
}
