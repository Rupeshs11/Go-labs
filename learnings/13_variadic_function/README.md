# 13 — Variadic Functions

A variadic function accepts **any number of arguments** of the same type using `...` syntax. Internally, the arguments are received as a **slice**.

## Syntax

```go
func sum(nums ...int) int { }    // nums is []int inside the function
```

- Call with individual values: `sum(1, 2, 3)`
- Call with a slice using spread: `sum(nums...)`

## Code — `variadic.go`

```go
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	fmt.Println("the sum:", sum(nums...))    // 30
}
```

- `nums...` **spreads** the slice into individual arguments
- Without `...` you'd get a type error — `sum(nums)` won't work because `nums` is `[]int`, not individual `int`s
- Common real-world example: `fmt.Println()` itself is variadic — it accepts any number of arguments

```bash
go run variadic.go
# Output: the sum of the numbers : 30
```
