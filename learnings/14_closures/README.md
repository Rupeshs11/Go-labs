# 14 — Closures

A closure is a function that **remembers the variables from its outer scope** even after the outer function has returned. This lets you create functions with persistent state.

## Code — `closure.go`

```go
func counter() func() int {
	count := 0                    // this variable persists across calls
	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := counter()        // get a closure
	fmt.Println(increment())     // 1
	fmt.Println(increment())     // 2
	fmt.Println(increment())     // 3
}
```

### How it works

1. `counter()` creates a local `count` variable and returns an **anonymous function**
2. That inner function has access to `count` even after `counter()` finishes — this is the closure
3. Each call to `increment()` modifies the **same** `count` variable
4. If you call `counter()` again, you get a **new** closure with its own independent `count`

### When to use closures

- Creating counters/accumulators
- Wrapping functions with extra behavior (middleware pattern)
- Callback functions that need to remember state

```bash
go run closure.go
# Output: 1
```
