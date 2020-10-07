package main

import (
	"fmt"
)

func main() {

	// Define an array
	x := [5]int{1, 2, 4, 5, 5}

	for i,v := range x {
		fmt.Println("Position ", i, "has a value of ", v)
	}
	fmt.Printf("The array is of type %T\n", x)
}
