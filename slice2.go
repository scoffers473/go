package main

import (
	"fmt"
)

func main() {
	// Composite literal
	x := []int{1, 2, 5, 7, 42}
	fmt.Println(x)

	// Slice a slice
	fmt.Println(x[1:])  // Pos 1 to end
	fmt.Println(x[:3])  // Start to but not including Pos 3
	fmt.Println(x[3:4]) // From pos 3 upto but not including pos 4 - Basically pos 3

	for i := 0; i <= 4; i++ {
		fmt.Println(x[i])
	}
}
