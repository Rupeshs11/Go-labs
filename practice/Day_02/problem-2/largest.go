package main

import "fmt"

func findMax(nums []int) int {
    max := nums[0]

    for i := 1; i < len(nums); i++ {
        if nums[i] > max {
            max = nums[i]
        }
    }

    return max
}

func main() {
    nums := []int{3, 7, 2, 9, 5}
    result := findMax(nums)

    fmt.Println("Largest element:", result)
}