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
	getUser := closures.NewUser("Ibrahim", "Hossinie")
	fmt.Println(getUser())
}
