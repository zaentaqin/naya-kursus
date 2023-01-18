package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Masukkan angka: ")
		input, _ := reader.ReadString('\n')
		//input = input[:len(input)-1]

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Print("Salah masukan")
			continue
		}

		result := factorial(num)
		fmt.Printf("Faktorial dari %d adalah %d\n", num, result)
		break
	}
}
