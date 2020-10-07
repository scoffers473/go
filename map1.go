package main

import (
	"fmt"
)

func main() {

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])

	// We get a zero value back if bad key
	// but this might be a problem

	fmt.Println(m["Scoffers"])

	// So use comma ok idiom to check
	v, ok := m["Scoffers"]
	fmt.Println(v)
	fmt.Println(ok)

	// So on a one-er
	if v, ok := m["Scoffers"]; ok {
		fmt.Println("This is it:", v)
	} else {
		fmt.Println("Nope!")
	}

	// Now add an element
	m["Scoffers"] = 53
	if v, ok := m["Scoffers"]; ok {
		fmt.Println("This is it:", v)
	} else {
		fmt.Println("Nope!")
	}


	// Now delete an entry
	delete(m,"Miss Moneypenny")
	// Range over te map
	for key,val := range m {
		fmt.Println(val,key)
	}

}
