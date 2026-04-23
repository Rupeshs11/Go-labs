package main

import "fmt"

func sumArray(nums []int) int {
    sum := 0

    for _, num := range nums {
        sum += num
    }

    return sum
}

func main() {
    nums := []int{1, 2, 3, 4}
    result := sumArray(nums)

    fmt.Println("Sum of array:", result)
}