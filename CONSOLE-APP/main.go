/*package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("->")

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		if userInput == "quit" || userInput == "exit" {
			break

		} else {
			fmt.Println("You entered:", userInput)
		}
	}
}
*/

// Go lang doesn't have a method to read a sigle character and then do something
// for this, we can use third party packages/librbries like "github.com/eiannone/keyboard"

//-------------------------------------------------------------
/*
package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press any key on the keyboard. Press ESC to quit")

	for {
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if key == keyboard.KeyEsc {
			break
		}
		if key != 0 {
			fmt.Println("You entered:", char, key)
		} else {
			fmt.Println("You entered:", char)
		}

	}
	fmt.Println("Program exited")

}

*/

//-------------------------------------------------------------

package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1. Cappucino")
	fmt.Println("2. Latte")
	fmt.Println("3. Espresso")
	fmt.Println("4. Mocha")
	fmt.Println("5. Americano")
	fmt.Println("6. Macchiato")
	fmt.Println("Q. - Quit the program")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		t := fmt.Sprintf("You chose %d", char) // sprintf returns the formatted string
		fmt.Println(t)
		fmt.Println("You chose", char)
		if char == 'q' || char == 'Q' {
			break
		}

	}
	fmt.Println("Program exited")

}
