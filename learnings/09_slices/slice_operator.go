package main

import (
	"fmt"
	"slices"
)

func main() {
	//slice operator

	// var nums = []int{1, 2, 3}

	// fmt.Println(nums[0:3])
	// fmt.Println(nums[:2])
	// fmt.Println(nums[1:3])
	// fmt.Println(nums[2:])

	// to compare two slices

	var nums1 = []int{1, 2}
	var nums2 = []int{1, 3}
	fmt.Println(slices.Equal(nums1,nums2))
}
