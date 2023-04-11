package main

import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("Please enter your name: ")
    fmt.Scanln(&name)
    fmt.Print("Please enter your age: ")
    fmt.Scanln(&age)
    fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
    
    if age >= 21 {
        fmt.Println("You are old enough to drink alcohol!")
    } else {
        fmt.Println("Sorry, you are not old enough to drink alcohol.")
    }
}
