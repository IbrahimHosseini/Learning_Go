package main

import (
	"fmt"
	"myproject/DSA"
	"myproject/closures"
)

func main() {

	callClosure()

	callLinkedList()
}

func callLinkedList() {
	DSA.PrintTest()
}

// The function `callClosure` creates a new user with the name "Ibrahim Hossinie" using a closure and
// then prints the user information.
func callClosure() {

	// `getUser := closures.NewUser("Ibrahim", "Hossinie")` is creating a new user with the name "Ibrahim
	// Hossinie" using a closure. The `NewUser` function in the `closures` package is likely a closure
	// that takes the first and last name of the user as parameters and returns a function that can be
	// called to retrieve the user information. In this case, the returned function is stored in the
	// `getUser` variable, which can then be called to get the user information and print it.
	getUser := closures.NewUser("Ibrahim", "Hossinie")
	fmt.Println(getUser())

	// The line `isDeleted := closures.DeleteUser("2213")` is likely creating a closure that deletes a
	// user with the ID "2213" and returns a function that can be called to check if the user was
	// successfully deleted. The `DeleteUser` function in the `closures` package is expected to be a
	// closure that takes the user ID as a parameter and returns a function that indicates whether the
	// user was deleted or not.
	isDeleted := closures.DeleteUser("2213")
	fmt.Println(isDeleted())

	fmt.Println("=============================================")

	// The line `deposit := closures.NewAccount(1000.0)` is creating a new account with an initial balance
	// of 1000.0 using a closure. The `NewAccount` function in the `closures` package is likely a closure
	// that takes the initial balance as a parameter and returns a function that can be called to deposit
	// or withdraw funds from the account.
	deposit := closures.NewAccount(1000.0)

	// `fmt.Println(deposit(100.0))` is calling the `deposit` function that was created using a closure in
	// the `closures` package. This closure function is likely designed to handle depositing or
	// withdrawing funds from an account.
	fmt.Println("Deposit:", deposit(100.0))

	// The line `fmt.Println(deposit(-50.0))` is calling the `deposit` function that was created using a
	// closure in the `closures` package with a parameter of -50.0.
	fmt.Println("New Deposit:", deposit(-50.0))

	fmt.Println("=============================================")

	// `service := closures.NewBanckService()` is creating a new bank service instance using a closure
	// function `NewBanckService` from the `closures` package. This closure function likely initializes
	// and returns a service object that provides banking functionalities such as transferring funds
	// between accounts. The `service` variable is then used to call the `Transfer` method on the bank
	// service instance to transfer funds between two accounts.
	service := closures.NewBanckService()

	// The line `fmt.Println(service.Transfer(1000.0, 400.0))` is calling the `Transfer` method on the
	// bank service instance `service` created using a closure in the `closures` package. This method is
	// likely designed to transfer funds between two accounts.
	fmt.Println("Balance:", service.Transfer(1000.0, 400.0))

	fmt.Println("=============================================")
}
