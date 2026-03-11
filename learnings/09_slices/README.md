# 09 — Slices

Slices are the **most used data structure** in Go. Unlike arrays, they're **dynamic** — they can grow and shrink. Think of them as a flexible wrapper around arrays.

## Slice vs Array

|           | Array            | Slice             |
| --------- | ---------------- | ----------------- |
| Size      | Fixed (`[5]int`) | Dynamic (`[]int`) |
| Can grow? | ❌               | ✅ via `append()` |
| Usage     | Rare             | Very common       |

## Key Functions

| Function                  | What it does                            |
| ------------------------- | --------------------------------------- |
| `make([]int, len, cap)`   | Create a slice with length and capacity |
| `append(slice, elems...)` | Add elements (grows automatically)      |
| `copy(dst, src)`          | Deep copy elements into another slice   |
| `len(s)` / `cap(s)`       | Get length / capacity                   |
| `slices.Equal(s1, s2)`    | Compare two slices (Go 1.21+)           |

---

## `slices.go` — Slice Basics (commented reference)

```go
var nums []int                 // nil slice — uninitialized, but usable with append
fmt.Println(nums == nil)       // true

var nums = make([]int, 2, 5)   // length=2, capacity=5
nums = append(nums, 1)         // adds to the end, auto-grows if needed

nums := []int{2, 4, 5}        // shorthand — length and capacity = 3
nums[0] = 3                   // modify by index
```

- **`nil` slice** is safe to use — `append` works on it
- **Capacity** auto-doubles when exceeded

---

## `slice_operator.go` — Slicing & Comparison

```go
nums := []int{1, 2, 3}

nums[0:3]   // [1 2 3]  — full range
nums[:2]    // [1 2]    — from start to index 2 (exclusive)
nums[1:3]   // [2 3]    — index 1 to 3
nums[2:]    // [3]      — index 2 to end
```

The `high` index is always **exclusive** (not included).

**Comparing slices** — you can't use `==` (only works to check `nil`), use `slices.Equal()` instead:

```go
slices.Equal([]int{1, 2}, []int{1, 3})   // false
```

---

## `copy_func.go` — Copying Slices (commented reference)

```go
nums := make([]int, 0, 5)
nums = append(nums, 2)

nums2 := make([]int, len(nums))   // must pre-allocate with enough length
copy(nums2, nums)                 // deep copy — nums2 is now independent
```

> ⚠️ Simple assignment `nums2 = nums` makes both point to the **same** data. Use `copy()` when you need a truly independent duplicate.

```bash
go run slice_operator.go
# Output: false
```
