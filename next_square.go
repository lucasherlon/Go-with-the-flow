package main 

import (
	"fmt"
	"math"
)

func isPerfectSquare(num int64) bool {
	return math.Sqrt(float64(num)) == math.Trunc(math.Sqrt(float64(num)))
}

func FindNextSquare(sq int64) int64 {
  if !isPerfectSquare(sq) {
  	return -1
  }

  number := int64(math.Sqrt(float64(sq))) + 1
  res := int64(math.Pow(float64(number), 2))
  return res
}



func main() {
	var num int64
	num = 114
	next := FindNextSquare(num)
	fmt.Println(next)

}