package main

import (
	"io" // Standard output library
	"os" // Used to read the arguments in the terminal
)

func main() {
	myString := ""
	arguments := os.Args // take the arguments passed in the execution of the file
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}
	
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}