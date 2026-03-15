# 19 — Generics

Generics (added in **Go 1.18**) let you write **one function or struct** that works with multiple types, instead of duplicating code for each type.

## The Problem (Without Generics)

```go
// You'd need a separate function for every type 😩
func printIntSlice(items []int) { ... }
func printStringSlice(items []string) { ... }
func printBooleanSlice(items []bool) { ... }
```

## The Solution (With Generics)

```go
// One function handles all types ✅
func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}
```

## Type Constraints

| Constraint              | Meaning                                                    |
| ----------------------- | ---------------------------------------------------------- |
| `any`                   | Accepts any type (same as `interface{}`)                   |
| `comparable`            | Types that support `==` and `!=` (int, string, bool, etc.) |
| `int \| string \| bool` | Only these specific types                                  |

> For best practice, be specific with constraints instead of just using `any`.

---

## `generics.go` — Generic Function

```go
func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	printSlice([]int{11, 12})               // works with int
	printSlice([]string{"Knox", "luffy"})   // works with string
	printSlice([]bool{true, false})          // works with bool
}
```

- `[T comparable]` declares a type parameter `T` that must be comparable
- The same function handles `[]int`, `[]string`, and `[]bool` — no duplication

---

## `stack.go` — Generic Struct (commented reference)

```go
type stack[T any] struct {
	elements []T
}

mystack := stack[int]{
	elements: []int{1, 2, 3},
}
fmt.Println(mystack)    // {[1 2 3]}
```

- `stack[T any]` — a generic struct where `T` can be any type
- `stack[int]` creates a stack of integers, `stack[string]` would create a stack of strings

```bash
go run generics.go
```
