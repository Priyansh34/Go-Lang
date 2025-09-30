package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// personal data type
type person struct {
	FirstName string
	LastName  string
	Age       int
}

type address struct {
	house string
	city  string
	state string
}

type contact struct {
	email string
	phone string
}

type employee struct {
	employee_details person
	emplpyee_address address
	employee_contact contact
}

func main() {
	fmt.Println("Let's start learning Struct in GO Lang")

	var employee1 employee
	employee1.employee_contact.email = "priyansh3401@gmail.com"
	employee1.employee_details = person{
		FirstName: "Priyansh",
		LastName:  "Vashistha",
		Age:       24,
	}
	employee1.emplpyee_address = address{
		house: "Mera Ghar",
		city:  "Deeg",
		state: "Rajasthan",
	}

	fmt.Println("Employee details:", employee1)

	var user person
	fmt.Println("user details are:", user)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your first name:")
	userInput, _ := reader.ReadString('\n')
	user.FirstName = userInput

	fmt.Println("Enter your last name:")
	lastInput, _ := reader.ReadString('\n')
	user.LastName = lastInput

	fmt.Println("Enter your age:")
	var age int
	_, err := fmt.Scanln(&age)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	user.Age = age

	fmt.Println("Do you want to see the user details? Yes/No")
	for {
		ans, _ := reader.ReadString('\n')
		ans = strings.Replace(ans, "\r\n", "", -1)
		if ans == "Yes" || ans == "yes" || ans == "YES" {
			fmt.Println("Details of the user:", user)
			break
		} else if ans == "No" || ans == "no" || ans == "NO" {
			fmt.Println("Thanks for your input. Exitting from the program.....")
			break
		} else {
			fmt.Println("Please enter yes or no")
		}
	}

}
