package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Say how many numbers have to be fizzbuzzed:")
	fmt.Scanln(&n)
	fizzbuzz(n)
}

func fizzbuzz(n int) {
	for i := 0; i < n; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}