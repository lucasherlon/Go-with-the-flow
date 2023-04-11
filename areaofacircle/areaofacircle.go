package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = 3.14159
	var radius float64

	fmt.Scanln(&radius)

	area := pi * math.Pow(radius, 3)
	fmt.Printf("%.4f", area)
}