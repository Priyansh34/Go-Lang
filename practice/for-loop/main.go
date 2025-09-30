package main

import "fmt"

func main() {
	for i := 0; i < 1; i++ {
		{
			fmt.Println(i)
		}
	}

	counter := 0
	for {
		fmt.Println("Infinite loop") // infinite for loop
		counter++

		if counter == 1 {
			break
		}
	}

	// range
	// this keywork simplifies the looping over elements of a collection like slices, arrays, maps, and strings
	// range keyword returns the index and value of the element in slice
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Indeex of numbers is %d, value of numbers is %d\n", index, value)

	}

	data := "Priyansh"
	for index, char := range data {
		fmt.Printf("index of char is %d, stored emelent is %c\n", index, char)
	}
}
