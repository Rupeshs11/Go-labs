package main

import "fmt"

func main(){
	subjects := map[string]int{
		"Math":98,
		"English":97,
		"science":92,
		"technical":96}
	fmt.Println(subjects)

	// to check whether the item is present in the map or not

	k, ok:=subjects["English"]
	fmt.Println(k)
	if ok{
		fmt.Println("present")
	}else{
		fmt.Println("not present")
	}
	
}

