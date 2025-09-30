package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and don't type your number in, just press ENTER when you're ready"

func main() {
	/*
		// one way - declate, then asign (two steps)
		var firstNumber int
		firstNumber = 3
		// another way - declate and asign (one step)
		var secondNumber = 5
		//one step variable: declare name, asign value, and let Go figure out the type
		thirdNumber := 9
	*/

	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2 // random number between 2 and 9
	var secondNumber = rand.Intn(8) + 2
	var thirdNumber = rand.Intn(8) + 2
	//var answer int
	var answer = firstNumber*secondNumber - thirdNumber

	game(firstNumber, secondNumber, thirdNumber, answer)

}

func game(firstNumber, secondNumber, thirdNumber, answer int) {

	reader := bufio.NewReader(os.Stdin)

	//display a welcome message

	fmt.Println("Welcome to the Guess the number game!")
	fmt.Println(("--------------------------------"))
	fmt.Println("")

	//fmt.Println("Think a number between 1 and 10 and press ENTER when you're ready")

	fmt.Println("Think a number between 1 and 10", prompt)
	reader.ReadString('\n')
	//wait for them to press enter

	//take them through the game
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply that by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Devide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", thirdNumber, prompt)
	reader.ReadString('\n')

	// give them the answer
	//answer = firstNumber*secondNumber - thirdNumber
	fmt.Println("The answer is", answer)

}
