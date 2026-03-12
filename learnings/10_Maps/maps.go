package main

import "fmt"

// maps -> hash , object , dict
func main() {
	// creating map
	m := make(map[string]string)
	// setting the elements

	m["name"] = "knox"
	m["area"] = "cloud"

	//get an element
	fmt.Println(m["name"], m["area"])
	fmt.Println(len(m))

	// if the keyvalue doesn't exist in the map then it returns zero or empty value

}
