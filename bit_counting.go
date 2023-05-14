package main

import (
	"fmt"
	"strconv"
)

func CountBits(num uint) int {
	bin := strconv.FormatInt(int64(num), 2)
	var sum int
	for _, char := range bin {
		if char == '1' {
			sum++
		}
	}
	return sum
}

func main() {
	var example uint
	example = 1234
	sum := CountBits(example)
	fmt.Println(sum)
}