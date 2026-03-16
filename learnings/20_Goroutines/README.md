# 20 — Goroutines

Goroutines are Go's way of doing **concurrency** — running multiple tasks at the same time. A goroutine is a **lightweight thread** managed by the Go runtime (not the OS), so you can spin up thousands of them cheaply.

## Syntax

```go
go functionName(args)    // that's it — just add "go" before the call
```

> ⚠️ **Problem:** When `main()` exits, all goroutines die — even if they haven't finished. You need a way to **wait** for them.

## Two Ways to Wait

| Method           | When to Use                                        |
| ---------------- | -------------------------------------------------- |
| `time.Sleep()`   | Quick & dirty — not reliable ❌                    |
| `sync.WaitGroup` | Proper way — waits for all goroutines to finish ✅ |

---

## `goroutine.go` — Basic Goroutine (commented reference)

```go
func task(id int) {
	fmt.Println("still Pending", id)
}

func main() {
	for i := 0; i <= 10; i++ {
		go task(i)               // launches 11 goroutines
	}
	time.Sleep(time.Second * 2)  // crude wait — not ideal
}
```

- `go task(i)` launches each call as a separate goroutine
- `time.Sleep` is unreliable — if tasks take longer than 2 seconds, they'll get killed

---

## `waitgroup.go` — WaitGroup (Proper Way) ✅

```go
func task(id int, w *sync.WaitGroup) {
	defer w.Done()                    // marks this goroutine as done when function returns
	fmt.Println("still pending", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)                     // tell WaitGroup: one more goroutine to wait for
		go task(i, &wg)               // pass WaitGroup by pointer
	}
	wg.Wait()                         // blocks until all goroutines call Done()
}
```

### WaitGroup cheat sheet

| Method      | What it does                                    |
| ----------- | ----------------------------------------------- |
| `wg.Add(1)` | Increment counter — "one more task to wait for" |
| `wg.Done()` | Decrement counter — "this task is finished"     |
| `wg.Wait()` | Block until counter reaches 0                   |

- `defer w.Done()` ensures `Done()` is called even if the function panics
- Always pass WaitGroup as a **pointer** (`*sync.WaitGroup`) — otherwise each goroutine gets a copy

```bash
go run waitgroup.go
```
