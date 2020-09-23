package main

import "fmt"

var y = 42

func main() {

	fmt.Println(y)
	fmt.Printf("Type %T\n", y)                 // Type
	fmt.Printf("Binary %b\n", y)               // In Binary
	fmt.Printf("Octal %o\n", y)                // In Octal
	fmt.Printf("Hex %x\n", y)                  // In Hex
	fmt.Printf("Hex with leading 0x %#x\n", y) // Hex with leading 0x
	fmt.Printf("And now altogether %v %b %o %x %#x\n", y, y, y, y, y)

	s := fmt.Sprintf("And now altogether %v %b %o %x %#x\n", y, y, y, y, y) // Use sprintf to create formatted string
	fmt.Println(s)
}