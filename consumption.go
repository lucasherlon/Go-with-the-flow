package main

import "fmt"

func main() {
	var km int64
	var fuel float64

	fmt.Scanln(&km)
	fmt.Scanln(&fuel)

	consumption := float64(km) / fuel

	fmt.Printf("%.3f km/l\n", consumption)
}
