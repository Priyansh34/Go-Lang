package main

import "fmt"

func main() {
	age := 24
	name := "Priyansh"
	height := 5.11
	fmt.Println("My name is", name, "I am", age, "years old and my height is", height)

	//printf is used to format the output
	// %s is used for string
	// %d is used for integer
	// %f is used for float
	// %.2f is used to print float with 2 decimal points
	// %T is used to print the type of variable
	fmt.Printf("name: %s\n", name)
	fmt.Printf("age: %d\n", age)
	fmt.Printf("height: %.2f\n", height)
	fmt.Printf("Type or name: %T", name)
}
