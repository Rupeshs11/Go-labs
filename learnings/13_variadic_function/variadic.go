package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}
	return total
}

func main() {
	nums:=[]int{2,4,6,8,10}
	fmt.Println("the sum of the numbers :", sum(nums...))
}
