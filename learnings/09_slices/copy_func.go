package main

import "fmt"

func main(){
	
// copy function in slices
var nums=make([]int,0,5)
nums=append(nums,2)
var nums2=make([]int,len(nums))
fmt.Println(nums,nums2)
//copy func
copy(nums2,nums)
fmt.Println(nums,nums2)

}