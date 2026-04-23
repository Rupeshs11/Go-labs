package main

import "fmt"

func countEvenOdd(nums []int) (int, int) {
    even := 0
    odd := 0

    for _, num := range nums {
        if num%2 == 0 {
            even++
        } else {
            odd++
        }
    }

    return even, odd
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    even, odd := countEvenOdd(nums)

    fmt.Println("Even:", even)
    fmt.Println("Odd:", odd)
}