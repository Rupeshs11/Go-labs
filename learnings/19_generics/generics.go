package main

import "fmt"

// implementaion
func printSlice[T comparable](items []T) {
	// also in the place of any we can use interface{} both are same
	// also for best pratices with don't use any instead we use int | string | bool ...
	for _, item := range items {
		fmt.Println(item)
	}
}

// // tradition method and problem
// func printIntSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printBooleanSlice(items []bool) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

func main() {
	//printSlice([]int{1,2,3})
	nums := []int{11,12}
	names := []string{"Knox","luffy"}
	isTrue := []bool{true,false}

	printSlice(names)
	printSlice(nums)
	printSlice(isTrue)

	// // traditional method
	// printBooleanSlice(isTrue)
	// printIntSlice(nums)
	// printStringSlice(names)
}
