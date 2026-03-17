# 22 — Mutex

When multiple goroutines access the same variable, you get a **race condition** — the result becomes unpredictable. `sync.Mutex` solves this by letting only **one goroutine access the shared data at a time**.

## Without Mutex ❌

```go
// 100 goroutines all doing views += 1 at the same time
// Result could be anything: 97, 99, 100... unpredictable!
```

## With Mutex ✅

```go
mu.Lock()        // "I'm using this — everyone else wait"
views += 1
mu.Unlock()      // "I'm done — next goroutine can go"
```

---

## Code — `mutex.go`

A social media post where 100 goroutines increment the view count:

```go
type post struct {
	views int
	mu    sync.Mutex       // embed mutex inside the struct
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()      // always unlock when done
		wg.Done()
	}()

	p.mu.Lock()            // lock before modifying
	p.views += 1
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()
	fmt.Println(myPost.views)   // always 100 ✅
}
```

### Key Points

- `mu.Lock()` — only one goroutine can pass this point at a time
- `mu.Unlock()` — releases the lock so others can proceed
- `defer` ensures unlock happens even if the function panics
- Without the mutex, `views` might end up less than 100 due to race conditions
- You can test for races with `go run -race mutex.go`

```bash
go run mutex.go
# Output: 100
```
