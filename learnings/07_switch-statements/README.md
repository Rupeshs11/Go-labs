# 07 — Switch Statements

A cleaner alternative to long `if/else if` chains. Go's switch has some nice differences from C/Java:

- **No `break` needed** — Go automatically breaks after each case
- **Multiple values** in one case — `case 1, 2, 3:`
- **Type switching** — check the runtime type of a variable

---

## `switch.go` — Basic Switch (commented reference)

A simple calculator: switches on `i` to decide which operation to perform.

```go
i := 2
n1, n2 := 10, 20

switch i {
case 1:
	fmt.Println(n1 + n2)    // 30
case 2:
	fmt.Println(n1 - n2)    // -10  ← this runs since i=2
case 3:
	fmt.Println(n1 * n2)    // 200
case 4:
	fmt.Println(n2 / n1)    // 2
default:
	fmt.Println("Wrong choice")
}
```

No `break` statement needed — Go handles it automatically.

---

## `multi-switch.go` — Multiple Values per Case

Checks if today is a weekday or weekend using the `time` package:

```go
switch time.Now().Weekday() {
case time.Saturday, time.Sunday:            // multiple values in one case
	fmt.Println("it's weekend")
case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
	fmt.Println("it's working day")
}
```

---

## `type-switch.go` — Type Switch (commented reference)

Used to check what **type** a value is at runtime. Works with `interface{}` (which accepts any type):

```go
whoAmI := func(i interface{}) {
	switch i.(type) {          // special syntax for type switching
	case int:
		fmt.Println("it is an integer")
	case string:
		fmt.Println("it is a string")
	case bool:
		fmt.Println("it is a boolean value")
	}
}

whoAmI(true)   // Output: it is a boolean value
```

- `interface{}` means "any type" (same as `any` in Go 1.18+)
- `i.(type)` is the special syntax to get the runtime type inside a switch
- Useful when dealing with JSON, APIs, or generic data

```bash
go run multi-switch.go
```
