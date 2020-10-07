package main

import (
	"fmt"
)

func main() {

	// Define a slice
	x := []int{1, 2, 3, 4, 5,6,7,8,9,10}

	for i,v := range x {
		fmt.Println("Position ", i, "has a value of ", v)
	}
	fmt.Printf("The slice is of type %T\n", x)
}
