package main

import (
	"fmt"
)

func main() {

	m := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being Evil", "Ice Cream", "Sunsets"},
	}

	for k, v := range m {
		fmt.Println("Looking at ", k)
		for i, d := range v {
			fmt.Printf("\tIndex %d:\tValue: %s\n", i, d)
		}
	}
	m["scoffield_chris"] = []string{"Punk music", "Playing the guitar", "going to festivales"}
	for k, v := range m {
		fmt.Println("Looking at ", k)
		for i, d := range v {
			fmt.Printf("\tIndex %d:\tValue: %s\n", i, d)
		}
	}
}
