package main

import (
	"fmt"
)

func main() {

	x := make([]int, 10, 12) // Pre-allocate underlying array size for compiled code efficiency
	fmt.Println(x)
	fmt.Println(len(x)) // Size of array
	fmt.Println(cap(x)) // Underlying capacity

	x[0] = 43
	x[9] = 999
	fmt.Println(x)

	x = append(x, 2112)
	fmt.Println(x)
	fmt.Println(len(x)) // Size of array
	fmt.Println(cap(x)) // Underlying capacity

	x = append(x, 2541)
	fmt.Println(x)
	fmt.Println(len(x)) // Size of array
	fmt.Println(cap(x)) // Underlying capacity

	x = append(x, 90125)
	fmt.Println(x)
	fmt.Println(len(x)) // Size of array
	fmt.Println(cap(x)) // Underlying capacity - Crare new array double sie as we've maxed out

}