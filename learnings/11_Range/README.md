# 11 — Range

`range` is used to iterate over data structures like slices, arrays, maps, and strings. It returns **index + value** on each iteration.

## Syntax

```go
for index, value := range collection { }
```

- Use `_` to ignore the index if you don't need it
- If you only need the index, skip the second variable: `for i := range collection { }`

## Code — `range.go`

```go
nums := []int{6, 7, 8}

// sum all elements
sum := 0
for _, num := range nums {    // _ ignores the index
	sum = sum + num
}
fmt.Println(sum)              // 21

// print index with value
for i, num := range nums {
	fmt.Println(i, num)
}
// 0 6
// 1 7
// 2 8
```

> `range` also works on **maps** (returns key, value) and **strings** (returns index, rune).

```bash
go run range.go
```
