package main

import (
	"fmt"
	"strings"
)

func Map(slice []string, f func(string) string) []string {
    mapped := make([]string, len(slice))
    for i, v := range slice {
        mapped[i] = f(v)
    }
    return mapped
}

func CamelCase(s string) string {
    slice := strings.Split(s, " ")
    newSlice := Map(slice, strings.Title)
    newString := strings.Join(newSlice, "")
    return newString
}

func main() {
	test := "hello case"
	res := CamelCase(test)
	fmt.Println(res)
}