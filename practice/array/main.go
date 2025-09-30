package main

import "fmt"

func main() {
	// fmt.Println("We are not learning Array in Go Lang")
	// var name [3]string // array declaration
	// name[0] = "Priyansh"
	// name[1] = "Bhavesh"
	// name[2] = "Ayush"

	// var numbers = [5]int{1, 2, 3, 4, 5} // initialize the error while declaration
	// fmt.Println("name of persons are:", name)
	// fmt.Println("Numbers stored in array are:", numbers)

	// fmt.Println("Length of name array is:", len(name)) //to find the lenght of an array

	var name [5]string
	name[3] = "Priyansh"
	name[0] = "Ram"
	fmt.Println(name)
	fmt.Printf("Array is %q\n", name) //%q means, result will show in qoutes
	// strings result comes in double qoutes like "Priyansh" and integer comes in single qoutes like '2'
}
