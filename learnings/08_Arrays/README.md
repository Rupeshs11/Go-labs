# 08 — Arrays

Arrays in Go are **fixed-size** sequences of elements of the same type. Once declared, the size **cannot change**.

## Key Things to Know

- Size is part of the type — `[5]int` and `[3]int` are **different types**
- Unset elements get **zero values** (`0` for int, `""` for string, `false` for bool)
- Arrays are **value types** — assigning one array to another **copies** all elements
- In practice, you'll mostly use **slices** (next topic) instead of arrays

---

## `int-arrays.go` — Integer Arrays (commented reference)

```go
var nums [5]int             // [0 0 0 0 0] — all zeros by default
fmt.Println(len(nums))      // 5

nums[0] = 1                 // set first element
fmt.Println(nums)           // [1 0 0 0 0]

arr := [4]int{2, 4, 6, 8}  // declare + initialize in one line
fmt.Println(arr)            // [2 4 6 8]
```

## `string-array.go` — String Array (commented reference)

```go
var name [3]string
name[0] = "golang"
fmt.Println(name)           // [golang  ] — unset strings are ""
```

## `bool-array.go` — Boolean Array (commented reference)

```go
var vals [4]bool
vals[1] = true
vals[3] = true
fmt.Println(vals)           // [false true false true]
```

## `2D-array.go` — Two-Dimensional Array

```go
nums := [2][2]int{{1, 2}, {3, 4}}
fmt.Println(nums)           // [[1 2] [3 4]]
```

A 2×2 matrix — useful for grids or tabular data.

```bash
go run 2D-array.go
# Output: [[1 2] [3 4]]
```
