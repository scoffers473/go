package main

import (
	"fmt"
)

func main() {

	// Define an array
	var x [5] int
	arrlen  := len(x) -1

	for  i:=0; i<= arrlen; i++ {
		x[i]=i+1
	}
	fmt.Println(x,arrlen)
}
