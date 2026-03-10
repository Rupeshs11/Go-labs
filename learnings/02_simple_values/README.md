# 02 — Simple Values (Data Types)

Go is a **statically typed** language. These are the basic value types you'll use everywhere:

| Type      | Example           | Description                          |
| --------- | ----------------- | ------------------------------------ |
| `int`     | `50`, `200`       | Whole numbers                        |
| `float64` | `10.5`, `7.0/3.0` | Decimal numbers (default float type) |
| `string`  | `"hello golang"`  | Text, always in double quotes        |
| `bool`    | `true`, `false`   | Boolean values                       |

> ⚠️ **Integer vs Float division:** `7/3` gives `2` (integer division), but `7.0/3.0` gives `2.333...` (float division)

## Code — `main.go`

```go
fmt.Println(50 * 200)       // 10000       → integer math
fmt.Println("hello golang") // hello golang → string
fmt.Println(true)           // true         → boolean
fmt.Println(10.5)           // 10.5         → float
fmt.Println(7.0 / 3.0)     // 2.333...     → float division
```

```bash
go run main.go
```
