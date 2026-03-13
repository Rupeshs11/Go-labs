# 15 — Pointers

Pointers let you **pass the memory address** of a variable instead of a copy. This allows functions to modify the original value.

## Quick Reference

| Symbol | Meaning                                           | Example                 |
| ------ | ------------------------------------------------- | ----------------------- |
| `*int` | Pointer type — "address of an int"                | `func change(num *int)` |
| `&num` | Get address — "give me the address of num"        | `change(&num)`          |
| `*num` | Dereference — "give me the value at this address" | `*num = 5`              |

## Code — `pointer.go`

### Pass by Value (doesn't modify original)

```go
func changeNum(num int) {
	num = 5                          // changes local copy only
}

num := 10
changeNum(num)
fmt.Println(num)                     // still 10 ❌
```

### Pass by Reference / Pointer (modifies original) ✅

```go
func changeNum(num *int) {           // accepts a pointer
	*num = 5                         // dereferences & changes the actual value
}

num := 10
changeNum(&num)                      // pass the address
fmt.Println(num)                     // now 5 ✅
```

### Key Takeaway

- **By value** = function gets a copy → original unchanged
- **By pointer** = function gets the address → can modify the original
- Use pointers when you want a function to **modify** the caller's variable, or to avoid copying large structs

```bash
go run pointer.go
```
