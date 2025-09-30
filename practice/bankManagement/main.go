package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type user struct {
	FullName    string
	DateOfBirth string
	Amount      int
	Account     int
}

func welcomeMessage() {
	fmt.Println("How can we help you?")
	fmt.Println("1. Create Account")
	fmt.Println("2. View Account Details")
	fmt.Println("3. Check Bank Balance")
	fmt.Println("4. Deposit Amount")
	fmt.Println("5. Withdraw Amount")
	fmt.Println("6. Delete My Account")
	fmt.Print("Choose an option: ")

}

func createAccount() {
	fmt.Print("You have selected option 1: Create Account. Do you want to create an account? (Yes/No).\n")
	reader := bufio.NewReader(os.Stdin)

	for {
		i := 0
		response, _ := reader.ReadString('\n')
		response = strings.Replace(response, "\r\n", "", -1)
		if strings.ToLower(response) == "yes" {
			fmt.Println("Please provide the details to create your account.")
			// Name input
			fmt.Print("Enter your name: ")
			UserName, _ := reader.ReadString('\n')
			UserName = strings.Replace(UserName, "\r\n", "", -1)

			// DOB input
			fmt.Print("Enter your DOB (dd-mm-yyyy):")
			UserDOB, _ := reader.ReadString('\n')
			UserDOB = strings.Replace(UserDOB, "\r\n", "", -1)

			// Amount input
			fmt.Print("Enter amount to open an account:")
			UserAmount, _ := reader.ReadString('\n')
			UserAmount = strings.Replace(UserAmount, "\r\n", "", -1)
			AmountInt, _ := strconv.Atoi(UserAmount)
			if AmountInt < 1000 {
				fmt.Println("Minimum amount to open an account is 1000. Please enter valid amount.")
				UserAmount, _ := reader.ReadString('\n')
				UserAmount = strings.Replace(UserAmount, "\r\n", "", -1)
				amount, _ := strconv.Atoi(UserAmount)
				AmountInt = amount
				if amount < 1000 {
					fmt.Println("You have entered invalid amount again. Exitting from the program.....")
					return
				}
			}

			// Account no. generating
			// Generate an account number
			// accountNumber := "ACC" + strconv.Itoa(100000+amountInt) // simple account number generation logic
			// Generate an account number randomly

			rand.Seed(time.Now().UnixNano())
			AccNum := rand.Intn(1000)

			// Append new user to the global userDetails slice
			userDetails = append(userDetails, user{
				FullName:    UserName,
				DateOfBirth: UserDOB,
				Amount:      AmountInt,
				Account:     AccNum,
			})

			fmt.Println("Congratulations", UserName, "Account Created Successfully!")
			fmt.Println("Your Account Details:")
			fmt.Println("Account Number:", AccNum)
			fmt.Println("Name:", UserName)
			fmt.Println("DOB:", UserDOB)
			fmt.Println("Initial Deposit Amount:", AmountInt)
			// asking user if he wants to create another account
			fmt.Println("Do you want to create another account? (Yes/No)")
			reader := bufio.NewReader(os.Stdin)
			response, _ := reader.ReadString('\n')
			response = strings.Replace(response, "\r\n", "", -1)
			if strings.ToLower(response) == "yes" {
				i++
				continue
			} else {
				break
			}
		} else {
			break
		}

	}

}

var userDetails []user

func viewAccountDetails() {
	fmt.Println("Enter your account number:")
	reader := bufio.NewReader(os.Stdin)
	accNum, _ := reader.ReadString('\n')
	accNum = strings.Replace(accNum, "\r\n", "", -1)
	acc, _ := strconv.Atoi(accNum)

	found := false
	for _, u := range userDetails {
		if u.Account == acc {
			fmt.Println("Your account details are below:")
			fmt.Println("Your Name:", u.FullName)
			fmt.Println("Your DOB:", u.DateOfBirth)
			fmt.Println("Your Account Number:", u.Account)
			fmt.Println("Your Bank Balance:", u.Amount)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Account not found. Please enter correct account number.")
	}
}

func checkBankBalance() {
	fmt.Println("Please enter your account number:")
	reader := bufio.NewReader(os.Stdin)
	account, _ := reader.ReadString('\n')
	account = strings.Replace(account, "\r\n", "", -1)
	acc, _ := strconv.Atoi(account)
	if acc >= 0 {
		fmt.Println("Your account balance is: 3500")
		fmt.Println("Thank you for banking with us!")
	}
}

func depositAmount() {
	fmt.Println("Please enter your account number:")
	reader := bufio.NewReader(os.Stdin)
	account, _ := reader.ReadString('\n')
	account = strings.Replace(account, "\r\n", "", -1)
	acc, _ := strconv.Atoi(account)
	if acc >= 0 {
		fmt.Println("Please enter the amount to deposit:")
		amount, _ := reader.ReadString('\n')
		amount = strings.Replace(amount, "\r\n", "", -1)
		amt, _ := strconv.Atoi(amount)
		fmt.Println("Successfully deposited", amt, "to your account.")
		fmt.Println("Do you want to see current balance? (yes/no)")
		choice, _ := reader.ReadString('\n')
		choice = strings.Replace(choice, "\r\n", "", -1)
		if strings.ToLower(choice) == "yes" {
			fmt.Println("Your current balance is:", 3500+amt)
		} else if strings.ToLower(choice) == "no" {
			fmt.Println("Thank you for banking with us!")
		} else {
			fmt.Println("Invalid input. Please enter yes or no.")
		}
	}
}

func withdrawAmount() {
	fmt.Println("Please enter your account number:")
	reader := bufio.NewReader(os.Stdin)
	account, _ := reader.ReadString('\n')
	account = strings.Replace(account, "\r\n", "", -1)
	acc, _ := strconv.Atoi(account)
	if acc >= 0 {
		fmt.Println("Please enter the amount to withdraw:")
		amount, _ := reader.ReadString('\n')
		amount = strings.Replace(amount, "\r\n", "", -1)
		amt, _ := strconv.Atoi(amount)
		if amt > 3500 {
			fmt.Println("Insufficient balance. Your current balance is 3500")
		} else {
			fmt.Println("Successfully withdrawn", amt, "from your account.")
			fmt.Println("Your current balance is:", 3500-amt)
			fmt.Println("Thank you for banking with us!")
		}
	}
}

func deleteMyAccount() {
	fmt.Println("Do you really want to delete your account? (yes/no)")
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.Replace(choice, "\r\n", "", -1)

	if strings.ToLower(choice) == "yes" {
		fmt.Println("Please enter your account number:")
		reader := bufio.NewReader(os.Stdin)
		account, _ := reader.ReadString('\n')
		account = strings.Replace(account, "\r\n", "", -1)
		acc, _ := strconv.Atoi(account)
		if acc >= 0 {
			fmt.Println("Your account has been successfully deleted.")
		}
	} else if strings.ToLower(choice) == "no" {
		fmt.Println("Thank you for banking with us!")
	} else {
		fmt.Println("Invalid input. Please enter yes or no.")
	}
}

func operations() {
	fmt.Println("Do you want to go back to main menu? (yes/no)")
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.Replace(choice, "\r\n", "", -1)
	if strings.ToLower(choice) == "yes" {
		main()
	}
}

func main() {
	fmt.Println("Welcome to our bank management System.")
	welcomeMessage()
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while getting user input", err)
		return
	}
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInt, _ := strconv.Atoi(userInput)
	if userInt == 1 {
		createAccount()
	} else if userInt == 2 {
		viewAccountDetails()
	} else if userInt == 3 {
		checkBankBalance()
	} else if userInt == 4 {
		depositAmount()
	} else if userInt == 5 {
		withdrawAmount()
	} else if userInt == 6 {
		deleteMyAccount()
	} else {
		fmt.Println("Please enter a valid input")
	}

	operations()

}
