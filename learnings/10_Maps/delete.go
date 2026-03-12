 package main

import "fmt"

func main(){
	m := make(map[string]int)
	//to delete the values form the map

	m["Roll no"] = 25
	m["marks"] = 99


	fmt.Println(m)



	delete(m,"marks")
	fmt.Println(m)

	// to clear everything
	clear(m)
}