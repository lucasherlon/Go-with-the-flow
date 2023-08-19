package main

import (
	"fmt"
)

func main() {
	levels := Beeramid(100, 100)
	fmt.Println(levels)
}


func Beeramid(bonus int, price float64) int {
      beers := float64(bonus) / price
      quant := int(beers)
      counter := quant
      levels := 0
      for i := 1; i <= quant; i++ {
        if counter - i*i < 0 {
          break
        }
        counter -= i * i
        levels++
       
      }
      return levels
}
