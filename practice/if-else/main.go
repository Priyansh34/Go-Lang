package main

import "fmt"

func main() {
	fmt.Println("Let's start If-Else in Go Lang")
	x := 12
	if x > 5 {
		fmt.Println("X is greater than 5")
	} else {
		fmt.Println("X is smaller than 5")
	}

	a := 20
	if a > 20 {
		fmt.Println("A is greater than 20")
	} else if a > 10 {
		fmt.Println("A is greater than 20 but smaller than 20")
	} else {
		fmt.Println("A is smaller than 10")
	}

	y := 10
	if x > 5 && y > 5 {
		fmt.Println("Hey how are you?")
	} else {
		fmt.Println("Learn Go lang...")
	}

	z := 15
	if x < 5 && (y > 5 || z < 10) {
		fmt.Println("Hey what's up buddy?")
	} else {
		fmt.Println("Let's learn Go Lang together....")
	}
}
