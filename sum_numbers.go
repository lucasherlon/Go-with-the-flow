package main

import "fmt"

func GetSum(a, b int) int {
	if a == b {
		return a
	}

	sum := 0

	if a > b {
		for i := b; i <= a; i++ {
			sum += i
		}
		return sum
	} else {
		for i := a; i <= b; i++ {
			sum += i
		}
		return sum
	}
}

func main() {
	sum := GetSum(-1, 2)
	fmt.Println(sum)
}