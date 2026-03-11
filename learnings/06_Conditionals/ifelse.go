package main

import "fmt"

func main() {
	// age := 7

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")

	// }else{
	// 	fmt.Println("Person is not adult")
	// }

	// we can also declare a variavle inside the if constructor

	if age:=7;age >=18{
		fmt.Println("Person is an adult",age)
	}else if age>=12{
		fmt.Println("Person is a teenager",age)
	}else{
		fmt.Println("person is a child",age)
	}
}


// note : there is no ternary operator in go so we need to use simple nested if-else instead