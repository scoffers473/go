package main

import "fmt"

func main () {
	fmt.Println("Hello Chris")
	foo()

}

func foo() {
	var i,_ = fmt.Println("And back atcha")
	fmt.Println("We wrote ", i, "characters in funcion foo but ignored any errors")
}