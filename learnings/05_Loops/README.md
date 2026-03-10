# 05 — Loops

Go has **only one loop keyword: `for`**. But it's flexible enough to work as a classic for loop, a while loop, a range loop, and even an infinite loop.

> ⚠️ There is **no** `while`, `do-while`, or `foreach` keyword in Go — everything is `for`.

## All Loop Patterns

```go
// Classic for loop
for i := 0; i <= 5; i++ { }

// While-style (just a condition)
for i <= 3 { }

// Range-based (Go 1.22+)
for i := range 5 { }

// Infinite loop
for { }
```

## Code — `for.go`

```go
// Classic for — prints 0, 2, 4, 6, 8, 10
for i := 0; i <= 5; i++ {
	n := 2 * i
	fmt.Println(n)
}

// Range loop — prints 0, 1, 2, 3, 4
for i := range 5 {
	fmt.Println(i)
}
```

## Code — `while.go` (commented reference)

```go
// While-style loop — Go doesn't have "while", you just omit init & post statements
i := 1
for i <= 3 {
	fmt.Println(i)   // prints 1, 2, 3
	i++
}

// Infinite loop — runs forever, use break to exit
for { }
```

```bash
go run for.go
```
