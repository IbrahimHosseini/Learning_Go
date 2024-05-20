package DSA

import (
	"fmt"
)

func PrintTest() {
	fmt.Println("DSA package")
}

type Node struct {
	data int
	next *Node
}
