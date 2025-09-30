package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

// this function can be defuned in such other ways
// func add(a int, b int) int {
// 	return a + b
// }

// other way to define function
// func add(a, b int) (sum int) { //we have already defined the return type as sum
// 	return a + b
//}

// another way to define function
// func add(a, b int) (sum int) { //we have already defined the return type as sum
// 	sum = a + b
// 	return // we can simply write return as we have already defined the return type as sum
//}

func multiply(a, b int) (multi int) {
	multi = a * b
	return
}

func main() {
	fmt.Println("we are learning functions in Go lang") // This is a placeholder main function.

	ans := add(5, 8)
	fmt.Println("The sum is:", ans)

	multi := multiply(5, 8)
	fmt.Println("The multiplication is:", multi)
}
