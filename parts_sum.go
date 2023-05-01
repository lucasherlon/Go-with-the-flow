package main

import 	"fmt"


func SumSlice(slice []uint64) uint64 {
	var sum uint64
	for _, val := range slice {
		sum += val
	}
	return sum
}

func PartsSums(ls []uint64) []uint64 {
	 if len(ls) == 0 {
	 	return []uint64{0}
	 }
	 var slice []uint64
	 sum := SumSlice(ls)
	 slice = append(slice, sum)
	 for i, _ := range ls {
	  	 sum -= ls[i]
	  	 slice = append(slice, sum)
	  }
	  return slice
}

func main() {
	test := []uint64{0, 1, 3, 6, 10}
	res := PartsSums(test)
	fmt.Println(res)
}