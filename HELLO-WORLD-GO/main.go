/*
package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
) // immporting multiple packages at the same time.

func main() {
	reader := bufio.NewReader(os.Stdin) // reading input from standard input (keyboard)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	for {
		fmt.Print("-> ")
		//taking input from user and reading it until newline character is encountered
		userInput, _ := reader.ReadString('\n') // ignoring the error value using _ (blank identifier)

		userInput = strings.Replace(userInput, "\r\n", "", -1) // to handle windows new line character
		userInput = strings.Replace(userInput, "\n", "", -1)   // to handle unix new line character
		//userInput = strings.Replace(userInput, "\r", "", -1)   // to handle mac new line character
		//fmt.Println("You typed: " + userInput)

		//fmt.Println(userInput)

		if userInput == "quit" {
			break
		} else {
			//response := doctor.Response(userInput)
			fmt.Println(doctor.Response(userInput))
		}

	}

}
*/

package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		// to handle windows new line character

		if userInput == "quit" {
			break
		} else {

			fmt.Println(doctor.Response(userInput))
		}

	}

}
