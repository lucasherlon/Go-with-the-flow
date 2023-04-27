package main

import ("fmt"
		"strings"
)

func Map(slice []string, f func(string) string) []string {
    mapped := make([]string, len(slice))
    for i, v := range slice {
        mapped[i] = f(v)
    }
    return mapped
}

func ReverseWords(str string) string {
    var res string 
    chars := []byte(str)
    for i := len(chars) - 1; i >= 0; i-- {
      res += string(chars[i])
    }
    return res
}

func main() {
	var nome = "This is an example!"
	lista := strings.Split(nome, " ")
	novaLista := Map(lista, ReverseWords)
	novaString := strings.Join(novaLista, " ")
	fmt.Println(novaString)
}