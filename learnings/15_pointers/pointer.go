package main

import "fmt"
// // by value
// func changeNum(num int){
// 	num =5
// 	fmt.Println("In chnageNum",num)
// }

// func main() {
// 	var num int
// 	fmt.Println("Enter the number")
// 	fmt.Scan(&num)
// 	fmt.Println("Before changeNumber",&num)
// 	changeNum(num)

// 	fmt.Println("Ater changeNumber", num)
// }

// by refernce 

func changeNum(num *int){
	*num=5 // dereferncing to change the value at that address
	fmt.Println("In changeNum:",*num)
}

func main(){
	var num int
	fmt.Println("Enter the number:")
	fmt.Scan(&num)
	fmt.Println("Before change:",num)
	changeNum(&num)
	fmt.Println("After change:",num)
}