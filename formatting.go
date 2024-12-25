package main

import "fmt"

// We use %t to print a boolean value
func printBoolean() {
	booleanValue := false
	fmt.Printf("Boolean value: %t\n", booleanValue)
}

// We use %s to print a string value
func printString() {
	stringValue := "This is a string value"
	fmt.Printf("String value: %s\n", stringValue)
}

// We use %d to print an integer value
func printInteger() {
	integerValue := 10
	fmt.Printf("Integer value: %d\n", integerValue)
}
