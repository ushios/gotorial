package main

import (
	"fmt"

	"github.com/ushios/gotorial/calc"
)

func main() {
	var a, b int
	a = 10
	b = 3

	sum := calc.Add(a, b)
	fmt.Println(sum)
}
