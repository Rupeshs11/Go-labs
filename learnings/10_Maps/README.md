# 10 — Maps

Maps are Go's built-in **key-value** data structure (like dictionaries in Python, objects in JS, or HashMaps in Java). Keys must be of the same type, and values must be of the same type.

## Key Things to Know

- Created with `make(map[keyType]valueType)` or literal syntax `map[string]int{"a": 1}`
- Accessing a missing key returns the **zero value** (no error)
- Use the **comma-ok** idiom to check if a key exists: `val, ok := m["key"]`
- `delete(m, "key")` removes a key, `clear(m)` removes everything
- Maps are **unordered** — iteration order is not guaranteed

---

## `maps.go` — Creating & Using Maps (commented reference)

```go
m := make(map[string]string)    // create an empty map

m["name"] = "knox"              // set values
m["area"] = "cloud"

fmt.Println(m["name"], m["area"])  // knox cloud
fmt.Println(len(m))                // 2

// accessing a key that doesn't exist returns "" (zero value for string)
```

---

## `marks.go` — Literal Syntax & Checking Key Existence (commented reference)

```go
subjects := map[string]int{
	"Math":      98,
	"English":   97,
	"science":   92,
	"technical": 96,
}

// comma-ok idiom — check if a key exists
k, ok := subjects["English"]    // k=97, ok=true
if ok {
	fmt.Println("present")
} else {
	fmt.Println("not present")
}
```

- `ok` is `true` if the key exists, `false` if it doesn't
- This is the safe way to check — since missing keys just return `0`, you can't tell if the value is actually `0` or missing without `ok`

---

## `delete.go` — Deleting from Maps (commented reference)

```go
m := make(map[string]int)
m["Roll no"] = 25
m["marks"] = 99

delete(m, "marks")    // removes the "marks" key
fmt.Println(m)        // map[Roll no:25]

clear(m)              // removes everything — map is now empty
```

---

## `packages.go` — Comparing Maps

```go
import "maps"

m1 := map[string]int{"price": 40, "phone": 2}
m2 := map[string]int{"price": 40, "phone": 2}
fmt.Println(maps.Equal(m1, m2))    // true

m3 := map[string]int{"price": 40, "phone": 2}
m4 := map[string]int{"price": 80, "phone": 2}
fmt.Println(maps.Equal(m3, m4))    // false
```

- You **can't** use `==` to compare maps — use `maps.Equal()` from the `maps` package (Go 1.21+)

```bash
go run packages.go
# Output:
# true
# false
```
