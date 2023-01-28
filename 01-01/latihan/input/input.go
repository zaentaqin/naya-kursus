package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama: ")

	input, _ := reader.ReadString('\n')

	fmt.Printf("Halo ! %s\n", input)
}
