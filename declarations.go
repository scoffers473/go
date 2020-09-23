package main

import "fmt"

var y = 42
var z string = "Hello Chris"
var a = `Chris said "Hello ..."` // Use backquotes for raw strings
func main() {

	x := 42 // Use := for first declaration and assignment
	fmt.Println(x)
	x = 99 // No need for the colon as already been decalred
	fmt.Println(a)
}
