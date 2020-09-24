package main

import (
	"fmt"
)

func main() {

	cnt := 10

	for cnt <= 100 {
		fmt.Printf("%d\t%d\n", cnt, cnt%4)
		cnt++
	}

}