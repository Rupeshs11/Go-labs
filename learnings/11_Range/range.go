package main

import "fmt"

func main() {
	// iterating over data structures
	nums:= []int {6,7,8}
	//sum:=0

	// for _,num:=range nums{
	// 	fmt.Println(num)
	// 	sum = sum+num
	// }
	// fmt.Println(sum)

	// printing index with number

	for i,num:=range nums{
		fmt.Println(i,num)
	}
}