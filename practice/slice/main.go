package main

import "fmt"

func main() {
	// numbers := []int{1, 2, 3, 4, 5}                 //initialize the slice
	// numbers = append(numbers, 4, 28, 53, 34, 28, 3) //add elements in existing slice
	// fmt.Println("Numbers are:", numbers)
	// fmt.Printf("Numbers %T\n", numbers)  //type of numbers slice
	// fmt.Println("Length:", len(numbers)) //length of slice

	age := make([]int, 3, 5) // create slice using make function

	age = append(age, 24) //adding elements in slice
	age = append(age, 30)
	age = append(age, 19)
	age = append(age, 53)
	age = append(age, 3)
	fmt.Println("Age Slice:", age)
	fmt.Println("Length of age slice is:", len(age))
	fmt.Println("Capacity of age slice is:", cap(age))
	// if length match with capacity and we add one more item in slice
	// length increases by +1 but capacity increases by *2

	//empty slise declaration
	// slice := []int{}
	// slice := make([]int, 0)
}
