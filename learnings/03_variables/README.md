# 03 — Variables

Go gives you multiple ways to declare variables. The key thing is **type inference** — Go can figure out the type from the value you assign.

## Ways to Declare

```go
var name string = "Knox"   // explicit type
var name = "Knox"          // type inferred (Go figures it out)
name := "Knox"             // shorthand — most common way ✅
```

> ⚠️ The `:=` shorthand **only works inside functions**, not at the package level.

## Zero Values

If you declare a variable without assigning anything, Go gives it a **zero value**:

| Type         | Zero Value |
| ------------ | ---------- |
| `int`        | `0`        |
| `float32/64` | `0.0`      |
| `string`     | `""`       |
| `bool`       | `false`    |

## Code — `variable.go`

```go
name := "Knox"         // string (inferred)
isAdult := true        // bool (inferred)
age := 21              // int (inferred)

var price float32      // declared without value → defaults to 0.0
price = 10             // assigned later
```

- The short-hand `:=` is the Go-idiomatic way to declare variables inside functions
- Use `var` when you need to declare without an initial value or need a specific type like `float32`

```bash
go run variable.go
# Output: Knox  true  21  10
```
