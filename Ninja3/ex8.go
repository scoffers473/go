package main

import (
	"fmt"
)

func main() {
	num := 16

	switch {
	case num%5 == 0:
		fmt.Println("The number can be divided by five")

	case num%3 == 0:
		fmt.Println("The number can be divided by three")
	case num%2 == 0:
		fmt.Println("The number can be divided by two")
	default:
		fmt.Println("Could be a prime?")

	}

}