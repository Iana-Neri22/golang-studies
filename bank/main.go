package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to the Bank!")
	fmt.Println("What would you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	// wantsCheckBalance := choice == 1

	if choice == 1 {
		fmt.Println("Your balance is", accountBalance)
	} else if choice == 2 {
		var depositAmount float64
		fmt.Print("Enter deposit amount: ")
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid deposit amount")
			return
		}

		accountBalance += depositAmount
		fmt.Println("Your balance is", accountBalance)
	} else if choice == 3 {
		var withdrawAmount float64
		fmt.Print("Enter withdraw amount: ")
		fmt.Scan(&withdrawAmount)

		if withdrawAmount <= 0 {
			fmt.Println("Invalid withdraw amount")
			return
		} 

		if withdrawAmount > accountBalance {
			fmt.Println("Insufficient funds")
			return
		}

		accountBalance -= withdrawAmount
		fmt.Println("Your balance is", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}

}