package main

import (
	"fmt"
	"bufio"
	"os"
	"math"
	"strconv"
	"strings"
)

func Greatest(a, b int64) int64 {
	great := (a + b + int64(math.Abs(float64(a-b)))) / 2
	return great
}

func main() {
	var numbers []int64
	var input string

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	vector := strings.Split(input, " ")

	for _, num := range vector {
		numb, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			fmt.Println("Error at parsing!")
		}
		numbers = append(numbers, numb)
	}

	a, b, c := numbers[0], numbers[1], numbers[2]

	intermediate := Greatest(a, b)
	biggest := Greatest(intermediate, c)

	fmt.Printf("%d eh o maior\n", biggest)
}
