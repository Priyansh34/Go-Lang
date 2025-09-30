package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Let's learn Defer in Go lang")
	fmt.Println("First line of the code")
	sum := add(34, 6)
	defer fmt.Println("Sum is:", sum)
	defer fmt.Println("Middle line of the code") // this line will be executed at last

	fmt.Println("Last line of the code")
}

// defer is used to execute a statement at last
// it is mostly used to close the file, close the database connection etc
// it is also used to free up the resources
// it is also used to print the statement at last
// it will be executed just before the main function is about to end

// if there are multiple defer statements then they will be executed in LIFO (Last In First Out) order
