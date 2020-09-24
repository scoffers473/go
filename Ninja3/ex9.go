package main

import (
	"fmt"
)

func main() {
	favSport := "Cricket"

	switch favSport{
	case "Footie":
		fmt.Println("Up the Us")
	case "Rugby":
		fmt.Println("You have odd balls")
	case "Cricket":
		fmt.Println("Small and shiny")
	default:
		fmt.Println("Weirdo")

	}

}