package main

import (
	"fmt"
)

func main() {
	// Composite literal
	x := []int{1, 2, 5, 7, 42}
	fmt.Println(x)

	x = append(x, 77, 88, 99, 112)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...) // ... unpack slice (variadic parameter)
	fmt.Println(x)

	// Use slice and append to delete from slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
