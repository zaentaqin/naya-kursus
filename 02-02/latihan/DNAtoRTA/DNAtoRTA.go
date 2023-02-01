package main

import (
	"fmt"
	"strings"
)

func DNAtoRNA(dna string) string {
	return strings.Replace(dna, "T", "U", -1)
}

func main() {
	fmt.Println(DNAtoRNA("GCAT"))
}
