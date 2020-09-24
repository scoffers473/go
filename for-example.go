package main

import (
	"fmt"
)

func main() {

	// a 'C' like for
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Println(i, j)
		}
	}

	// sA While style for
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// Infinite loop
	for {

		if x > 20 {
			break
		}
		fmt.Println(x)
		x++
	}

}
