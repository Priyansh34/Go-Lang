// in this program, we will take user input from the console/keyboard and print it back to the console
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What is your name?")
	fmt.Print("-> ")
	//var name string

	// taking user input
	// scan function is used to take input from the user
	// it takes the address of the variable where the input will be stored
	// &name means the address of the variable name

	// fmt.Scan(&name)
	// fmt.Println("Hello, Mr.", name)

	// scan function prints the input till the first space
	// eg. if I input Priyansh Vashistha, it will only take Priyansh as input

	//if we want to take whole line as input, there is a package called bufio which can be used to take whole line as input
	// we will use bufio.NewReader to create a new reader
	// then we will use reader.ReadString('\n') to read the input till the newline character

	reader := bufio.NewReader(os.Stdin)
	// stdin means standard input, which is the console/keyboard that reads by os package

	userInput, _ := reader.ReadString('\n')
	// ReadString returns two values, the input and an error
	// we are ignoring the error by using _
	// it takes input till the newline character or user hits enter key
	// userInput will contain the whole line of input including spaces

	fmt.Println("Hello, Mr.", userInput)

	// userInput will contain the newline character at the end, we can remove it using slicing
	//userInput = userInput[:len(userInput)-1]

}
