package main

import "fmt"

var y = 42
var z string = "Hello Chris"
func main() {

	x := 42 // Use := for first declaration and assignment
	fmt.Println(x)
	x = 99 // No need for the colon as already been decalred
	fmt.Println(y)
	fmt.Printf("Variable %s is of type %T\n",z,z)
}
