package main

import "fmt"

var outside = "global var" // Use var when defining outside a func
func main() {

	x := 42 // Use := for first declaration and assignment
	fmt.Println(x)
	x = 99 // No need for the colon as already been decalred
	fmt.Println(x)
	y := 100 + 70
	fmt.Println(y)
	z := "Chris S"
	fmt.Println(z)
	fmt.Println(outside)
}
