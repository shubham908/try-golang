package main

import "fmt"

var globalVariable string = "This is a global variable"

func main() {
	fmt.Println("Hello World!")
	fmt.Printf("Global variable: %s\n", globalVariable)

	globalVariable = "Global variable changed"
	fmt.Printf("Changed global variable: %s\n", globalVariable)

	printBoolean()
	printString()
	printInteger()
}
