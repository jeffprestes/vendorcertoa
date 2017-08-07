package main

import (
	"fmt"

	"github.com/jeffprestes/vendorcalc"
)

func main() {
	a := 2
	b := 2
	c := 2
	fmt.Println("Qual o delta? ", vendorcalc.Delta(a, b, c))
}
