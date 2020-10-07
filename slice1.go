package main

import (
	"fmt"
)

func main() {
	// Composite literal
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(x[2])

	for i, v := range x {
		fmt.Println(i, v)
	}

}
