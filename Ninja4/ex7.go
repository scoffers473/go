package main

import (
	"fmt"
)

func main() {

	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Monneypenny", "Heloooooo James"}
	xp := [][]string{jb, mp}

	fmt.Println(xp)
	for i, k := range xp {
		fmt.Println("Index: ", i)
		for j, y := range k {
			fmt.Printf("\t Index position:%d \t Value:%s\n", j, y)
		}
	}
}
