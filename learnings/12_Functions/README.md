# 12 — Functions

Functions in Go can **return multiple values**, take user input with `fmt.Scan`, and even return other functions.

## Syntax

```go
// single return
func add(a int, b int) int {
	return a + b
}

// multiple return values — specify types in parentheses
func getInfo() (string, int, bool) {
	return "knox", 11, true
}
```

> Go **requires** you to specify the return type(s). Multiple returns are very common — especially for returning `(result, error)`.

---

## `function.go` — Functions with Input & Multiple Returns

```go
func add(a int, b int) int {
	return a + b
}

func getLanguages() (string, int, bool) {
	return "knox", 11, true
}

func main() {
	var num1, num2 int
	fmt.Println("Enter the first number:")
	fmt.Scan(&num1)                          // reads user input into num1
	fmt.Println("Enter the second number:")
	fmt.Scan(&num2)

	fmt.Println("the sum:", add(num1, num2))
	fmt.Println(getLanguages())              // knox 11 true
}
```

- `fmt.Scan(&num1)` — reads from stdin. The `&` passes the **address** so Scan can write into the variable.
- Multiple return values are received directly: `name, num, ok := getLanguages()`

---

## `processIT.go` — Function Returning a Function (commented reference)

```go
func processIt() func(a int) int {
	return func(a int) int {
		return 11
	}
}

fn := processIt()    // fn is now a function
fn(12)               // calls the returned function
```

- Functions are **first-class** in Go — you can return them, pass them as arguments, and store them in variables.

```bash
go run function.go
```
