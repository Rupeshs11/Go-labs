package main

import "fmt"

//numbered sequence of specific length
func main(){
	var nums [5]int // array declaration

	//array length
	fmt.Println(len(nums))
	fmt.Println(nums)

	//adding element to the array 
	nums[0]=1
	fmt.Println(nums[0])
	fmt.Println(nums)

	// to declare the array in a single line
	arr := [4]int{2,4,6,8}
	fmt.Println(arr)


}