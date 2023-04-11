package main

import (
	"fmt"
	"os"
	"strconv" // converts string to other types
)

func main() {
		if len(os.Args) == 1 {
			fmt.Println("Please give one or more floats.")
			os.Exit(1)
		}

		arguments := os.Args
		min, _ := strconv.ParseFloat(arguments[1], 64) // the zero is the nome of the program
		max, _ := strconv.ParseFloat(arguments[1], 64) // that's why we use the one

		for i := 2; i < len(arguments); i++ {
			n, _ := strconv.ParseFloat(arguments[i], 64) // the _ is the error handling return
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}

		fmt.Println("Min:", min)
		fmt.Println("Max:", max)

}