package main

import (
	"fmt"
)

func main() {
	i := 2
	n1 := 10
	n2 := 20

	switch i {
	case 1:
		fmt.Println(n1+n2) // there is no need to write break in go handles it automatically
	case 2:
		fmt.Println(n1-n2)
	case 3:
		fmt.Println(n1*n2)
	case 4:
		fmt.Println(n2/n1)
	default:
		fmt.Println("Wrong choice")
	}

	
}
