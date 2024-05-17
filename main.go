package main

import (
	"fmt"
	"myproject/closures"
)

func main() {
	getUser := closures.NewUser("Ibrahim", "Hossinie")
	fmt.Println(getUser())
}
