 package main

import (
	"fmt"
)

func add(a int, b int)int{
	c:= a+b
	return c

}

func getLanguages()(string,int,bool){
	// to return the multiple value we need the define it type as shown above
	return "knox",11,true
}
func main(){
	var num1 int
	var num2 int
	fmt.Println("Enter the first number:")
	fmt.Scan(&num1)
	fmt.Println("Enter the second number:")
	fmt.Scan(&num2)

	
	fmt.Println("the sum of the numbers:",add(num1,num2))
	fmt.Println(getLanguages())
}