package main

import (
	"fmt"
	"log"
	"myproject/DSA"
	"myproject/closures"
	"myproject/errorhandling"
	"myproject/methods"
)

func main() {

	seprator("Clouser")

	callClosure()

	seprator("Methods")

	callMethods()

	seprator("LinkedList")

	callLinkedList()

	seprator("Error handling")

	callDivision()

}

func seprator(name string) {
	fmt.Printf("======================[ %s ]======================\n", name)
}

func callDivision() {
	result, err := errorhandling.Divide(10, 5)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %f \n", result)

	_, err2 := errorhandling.Divide(3, 0)

	if err2 != nil {
		log.Printf("❌ %s\n", err2)
	}

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

}

func callMethods() {

	// The line `user := methods.User{Name: "Ibrahim"}` is creating a new instance of the `User` struct
	// defined in the `methods` package and initializing the `Name` field of that struct with the value
	// "Ibrahim". This line of code is essentially creating a new user object with the name "Ibrahim"
	// using the `User` struct constructor.
	user := methods.User{Name: "Ibrahim"}

	// `fmt.Println(user.Greet())` is calling the `Greet` method on the `user` object created from the
	// `User` struct in the `methods` package and then printing the result of that method call using
	// `fmt.Println()`.
	fmt.Println(user.Greet())

	// The `user.Uppername()` method is likely a method defined on the `User` struct in the `methods`
	// package. This method is expected to convert the name of the user to uppercase or perform some
	// operation that modifies the name in some way. By calling `user.Uppername()`, the method is invoked
	// on the `user` object, potentially changing the name stored in the `Name` field of the `User` struct
	// to uppercase or applying any other transformation specified within the method implementation.
	user.Uppername()

	fmt.Println(user.Name)

	// `userPointer := &user` is creating a new variable `userPointer` that stores a memory address
	// pointing to the `user` object. This operation is known as taking the address of the `user` object
	// using the `&` operator in Go. By creating a pointer to the `user` object, any modifications made to
	// the object through the `userPointer` variable will directly affect the original `user` object. This
	// allows for indirect manipulation of the `user` object's fields and methods through the pointer
	// variable.
	userPointer := &user

	// The line `userPointer.SetName("Nima")` is invoking the `SetName` method on the object that
	// `userPointer` is pointing to. In this case, `userPointer` is a pointer to a `User` object, and the
	// `SetName` method is likely a method defined on the `User` struct in the `methods` package.
	userPointer.SetName("Nima")

	fmt.Println(userPointer.Name)

}
