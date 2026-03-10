package main

import "fmt"

func main(){
	//classic for loop

	for i:=0; i<=5 ; i++{
		n:=2*i
		fmt.Println(n)
	}

	for i:= range 5{
		fmt.Println(i)
	}
}