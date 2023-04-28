package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Map(slice []string, f func(string) (int, error)) []int {
    mapped := make([]int, len(slice))
    for i, v := range slice {
        mapped[i], _ = f(v)
    }
    return mapped
}

func maxValue(slice []int) int {
	max := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

func minValue(slice []int) int {
	min := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}

func HighAndLow(in string) string {
	slice := strings.Split(in, " ")
	intSlice := Map(slice, strconv.Atoi)
	max := maxValue(intSlice)
	min := minValue(intSlice)
  
   return fmt.Sprintf("%d %d", max, min)
}

func main() {
	input := "0 -1"
	output := HighAndLow(input)
	fmt.Println(output)
}