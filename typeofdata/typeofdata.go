package main

import (
	"fmt"
)

func main() {
	number := 10
	numberfloat := 15.5
	text := "hello, world"
	boolean := true

	fmt.Printf("number: %v type: %T\nfloat: %v type: %T\ntext: %v type: %T\nboolean: %v type: %T\n", number, number, numberfloat, numberfloat,text, text, boolean, boolean )
}