# 06 — Conditionals (if / else)

Standard `if/else` with a few Go-specific twists:

- **No parentheses** around the condition — `if age >= 18 { }` not `if (age >= 18) { }`
- **Curly braces `{ }` are mandatory**, even for one-line blocks
- You can **declare a variable inside the `if`** — it's scoped to the if/else chain
- Go has **no ternary operator** (`? :`) — use `if/else` instead

## Code — `ifelse.go`

```go
// declaring "age" directly inside the if statement
// age is only accessible within this if/else chain
if age := 7; age >= 18 {
	fmt.Println("Person is an adult", age)
} else if age >= 12 {
	fmt.Println("Person is a teenager", age)
} else {
	fmt.Println("person is a child", age)   // ✅ this runs → age is 7
}
```

- `if age := 7; age >= 18` → declares `age = 7`, then checks the condition
- Since `7 < 12 < 18`, it falls to the `else` block → prints `"person is a child 7"`
- This init-statement pattern keeps the variable scoped tightly — `age` doesn't leak outside

> 💡 **No ternary:** In other languages you'd write `result = age >= 18 ? "adult" : "child"`. In Go, you must use a full `if/else`.

```bash
go run ifelse.go
# Output: person is a child 7
```
