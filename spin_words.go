package main

import (
	"fmt"
	"strings"
)

func ReverseWord(str string) string {
    slice := []byte(str)
    for i := 0; i < len(slice)/2; i++ {
        j := len(slice) - i - 1
        slice[i], slice[j] = slice[j], slice[i]
    }
    return string(slice)
}

func SpinWords(str string) string {
	var spinned []string
	var res string
	list := strings.Split(str, " ")

	for _, word := range list {
		if len(word) >= 5 {
			spinned = append(spinned, ReverseWord(word))
		} else {
			spinned = append(spinned, word)
		}
	}
	res = strings.Join(spinned, " ")
	return res
}

func main() {
	test := "This is another test"
	res := SpinWords(test)
	fmt.Println(res)
}