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

// Account represents a bank account
type Account struct {
	Name      string
	DOB       string
	AccNo     int
	Balance   float64
	CreatedAt time.Time
}

var accounts = make(map[int]*Account)
var reader = bufio.NewReader(os.Stdin)

func welcomeMessage() {
	fmt.Println("==================== Bank Management System ====================")
	fmt.Println("Welcome to the Bank Management System")
	fmt.Println("Please choose an option:")
	fmt.Println("1. Create account")
	fmt.Println("2. View account details")
	fmt.Println("3. Check bank balance")
	fmt.Println("4. Deposit amount")
	fmt.Println("5. Withdraw amount")
	fmt.Println("6. Delete account")
	fmt.Println("0. Exit")
	fmt.Println("================================================================")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		welcomeMessage()
		fmt.Print("Enter choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			createAccount()
		case "2":
			viewAccount()
		case "3":
			checkBalance()
		case "4":
			depositAmount()
		case "5":
			withdrawAmount()
		case "6":
			deleteAccount()
		case "0":
			fmt.Println("Exiting... Goodbye! Thanks for visiting our bank.")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}

		fmt.Println()
	}
}

// createAccount handles creation of a new account
func createAccount() {
	fmt.Println("--- Create Account ---")
	fmt.Print("Enter Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter DOB (dd-mm-yyyy): ")
	dob, _ := reader.ReadString('\n')
	dob = strings.TrimSpace(dob)

	// Amount must be >= 1000
	var amount float64
	for {
		fmt.Print("Enter initial amount (minimum 1000): ")
		amtStr, _ := reader.ReadString('\n')
		amtStr = strings.TrimSpace(amtStr)
		amt, err := strconv.ParseFloat(amtStr, 64)
		if err != nil {
			fmt.Println("Invalid amount. Please enter a numeric value.")
			continue
		}
		if amt < 1000 {
			fmt.Println("Initial amount must be at least 1000.")
			continue
		}
		amount = amt
		break
	}

	accNo := generateAccountNo()
	acct := &Account{
		Name:      name,
		DOB:       dob,
		AccNo:     accNo,
		Balance:   amount,
		CreatedAt: time.Now(),
	}

	accounts[accNo] = acct
	fmt.Printf("Account successfully created. Account No: %d\n", accNo)
}

// viewAccount shows details for an account
func viewAccount() {
	fmt.Println("--- View Account Details ---")
	accNo := promptAccountNo()
	acct, ok := accounts[accNo]
	if !ok {
		fmt.Println("Account not found.")
		return
	}
	fmt.Println("Account Details:")
	fmt.Printf("Name      : %s\n", acct.Name)
	fmt.Printf("DOB       : %s\n", acct.DOB)
	fmt.Printf("Account No: %d\n", acct.AccNo)
	fmt.Printf("Balance   : %.2f\n", acct.Balance)
	fmt.Printf("Created At: %s\n", acct.CreatedAt.Format("02-Jan-2006 15:04:05"))
}

// checkBalance prints the available balance
func checkBalance() {
	fmt.Println("--- Check Bank Balance ---")
	accNo := promptAccountNo()
	acct, ok := accounts[accNo]
	if !ok {
		fmt.Println("Account not found.")
		return
	}
	fmt.Printf("Available balance for account %d is: %.2f\n", accNo, acct.Balance)
}

// depositAmount deposits money into an account
func depositAmount() {
	fmt.Println("--- Deposit Amount ---")
	accNo := promptAccountNo()
	acct, ok := accounts[accNo]
	if !ok {
		fmt.Println("Account not found.")
		return
	}

	var amt float64
	for {
		fmt.Print("Enter amount to deposit: ")
		amtStr, _ := reader.ReadString('\n')
		amtStr = strings.TrimSpace(amtStr)
		amtParsed, err := strconv.ParseFloat(amtStr, 64)
		if err != nil || amtParsed <= 0 {
			fmt.Println("Enter a valid positive amount.")
			continue
		}
		amt = amtParsed
		break
	}

	acct.Balance += amt
	fmt.Printf("Amount deposited successfully. New balance: %.2f\n", acct.Balance)
}

// withdrawAmount withdraws money from an account
func withdrawAmount() {
	fmt.Println("--- Withdraw Amount ---")
	accNo := promptAccountNo()
	acct, ok := accounts[accNo]
	if !ok {
		fmt.Println("Account not found.")
		return
	}

	var amt float64
	for {
		fmt.Print("Enter amount to withdraw: ")
		amtStr, _ := reader.ReadString('\n')
		amtStr = strings.TrimSpace(amtStr)
		amtParsed, err := strconv.ParseFloat(amtStr, 64)
		if err != nil || amtParsed <= 0 {
			fmt.Println("Enter a valid positive amount.")
			continue
		}
		amt = amtParsed
		break
	}

	if amt > acct.Balance {
		fmt.Println("Insufficient balance. Withdrawal failed.")
		return
	}

	acct.Balance -= amt
	fmt.Printf("Withdrawal successful. New balance: %.2f\n", acct.Balance)
}

// deleteAccount deletes an account after confirmation
func deleteAccount() {
	fmt.Println("--- Delete Account ---")
	accNo := promptAccountNo()
	acct, ok := accounts[accNo]
	if !ok {
		fmt.Println("Account not found.")
		return
	}

	fmt.Printf("Are you sure you want to delete account %d (yes/no)? ", accNo)
	resp, _ := reader.ReadString('\n')
	resp = strings.TrimSpace(strings.ToLower(resp))
	if resp == "yes" {
		delete(accounts, accNo)
		fmt.Printf("Account %d deleted successfully.\n", accNo)
	} else {
		fmt.Println("Account deletion cancelled.")
	}

	_ = acct // just to avoid unused var warnings if needed
}

// promptAccountNo prompts the user and returns the account number as int
func promptAccountNo() int {
	for {
		fmt.Print("Enter account number: ")
		accStr, _ := reader.ReadString('\n')
		accStr = strings.TrimSpace(accStr)
		accNo, err := strconv.Atoi(accStr)
		if err != nil {
			fmt.Println("Invalid account number. Please enter numeric value.")
			continue
		}
		return accNo
	}
}

// generateAccountNo generates a random account number between 1 and 1000
// and ensures uniqueness within the current in-memory map
func generateAccountNo() int {
	for {
		n := rand.Intn(1000) + 1 // 1-1000
		if _, exists := accounts[n]; !exists {
			return n
		}
		// if exists, try again
	}
}
