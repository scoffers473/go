package main

import (
	"fmt"
)

func main() {

	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)

	mp := []string{"Miss", "Moneypenny", "Minty", "Scotch"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)

	fmt.Println(xp[0][1])
}
