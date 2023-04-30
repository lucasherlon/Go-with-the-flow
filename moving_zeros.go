package main

import "fmt"

func MoveZeros(arr []int) []int {
	var zeros []int
	var newArr []int
	for _, num := range arr {
		if num == 0 {
			zeros = append(zeros, 0)
		} else {
			newArr = append(newArr, num)
		}
	}
	newArr = append(newArr, zeros...)
  	return newArr
}

func main() {
	test := []int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}
	newArr := MoveZeros(test)
	fmt.Println(newArr)
}