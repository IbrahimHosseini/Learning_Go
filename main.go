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
}
