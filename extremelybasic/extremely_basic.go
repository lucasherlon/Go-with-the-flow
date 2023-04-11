package main

import (
	"fmt"
)

func main() {
	var x int
	var y int

	fmt.Scanln(&x)
	fmt.Scanln(&y)

	fmt.Printf("X = %v\n", x+y)
}