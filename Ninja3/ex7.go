package main

import (
	"fmt"
)

func main() {
	x := "Miss"
	if x == "James" {
		fmt.Println("Bon;d")
	} else if x == "Miss" {
		fmt.Println("Monypenny")
	} else {
		fmt.Println(" Quire frainkly, some random ...")
	}
}