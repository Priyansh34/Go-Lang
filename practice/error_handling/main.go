package main

import "fmt"

func multiply(a, b int) int { // it will return the devision of two number in integer value
	return a * b
}

func devide(a, b float64) (float64, error) { //as we want to handle the error, so we return error as well
	// this function will return two values, one is float64 and other is error
	if b == 0 {
		return 0, fmt.Errorf("Cannot be devided by zero") // it will return 0 and error message
	}
	return a / b, nil // if there is no error, it will return nil

} // it will return the devision of two number in float value

// func devide(a, b float64) (float64, string) {
// 	if b == 0 {
// 		return 0, "Cannot be devided by zero"
// 	}
// 	return a / b, "nil"

// }

func main() {
	fmt.Println("we are learning error handling in Go lang") // This is a placeholder main function.

	multi := multiply(5, 6)
	fmt.Println("The multiplication of two numbers is:", multi)

	ans, _ := devide(10, 5) // we are ignoring the error message by using _
	fmt.Println("The devision of two numbers is:", ans)

}

// this is how we handled error in Go lang usin _
