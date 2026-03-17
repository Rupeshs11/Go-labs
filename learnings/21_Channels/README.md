# 21 — Channels

Channels are Go's way of **communicating between goroutines**. They let one goroutine send data to another safely.

## Syntax

```go
ch := make(chan int)       // unbuffered channel
ch := make(chan int, 100)  // buffered channel (holds 100 values)

ch <- 42                  // send value into channel
val := <-ch               // receive value from channel
close(ch)                 // close the channel
```

> ⚠️ **Unbuffered channels are blocking** — a send blocks until another goroutine receives, and vice versa.

---

## `channel.go` — Basic Channel Communication (commented reference)

```go
func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("Processing the number ...", num)
		time.Sleep(time.Second)
	}
}

func main() {
	numChan := make(chan int)
	go processNum(numChan)       // receiver goroutine

	for {
		numChan <- rand.IntN(100)  // sender (main) sends random numbers
	}
}
```

- `main` sends random numbers into the channel
- `processNum` goroutine receives and processes them one by one
- `range numChan` keeps receiving until the channel is closed

---

## `channel2.go` — Getting Results Back (commented reference)

```go
func sum(result chan int, num1 int, num2 int) {
	result <- num1 + num2       // send result back through channel
}

func main() {
	result := make(chan int)
	go sum(result, 4, 5)
	res := <-result              // blocks until result is ready
	fmt.Println(res)             // 9
}
```

Channels are great for getting **return values** from goroutines.

---

## `synchronization.go` — Using Channel as a Synchronizer (commented reference)

```go
func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("processing...")
}

func main() {
	done := make(chan bool)
	go task(done)
	<-done                       // blocks until task sends true
}
```

A `chan bool` can replace `WaitGroup` for simple cases — `<-done` blocks until the goroutine finishes.

---

## `buffered_chan.go` — Buffered Channels (commented reference)

```go
emailChan := make(chan string, 100)    // buffer of 100
done := make(chan bool)

go emailSender(emailChan, done)

for i := 0; i < 5; i++ {
	emailChan <- fmt.Sprintf("%d@gmail.com", i)
}
close(emailChan)     // must close to signal "no more data"
<-done               // wait for sender to finish
```

- Buffered channels don't block until the buffer is full
- **Always `close()` the channel** when done sending — otherwise `range` will deadlock waiting for more data

---

## `multi_chan.go` — Select Statement (Multiple Channels)

```go
chan1 := make(chan int)
chan2 := make(chan string)

go func() { chan1 <- 10 }()
go func() { chan2 <- "pong" }()

for i := 0; i < 2; i++ {
	select {
	case val := <-chan1:
		fmt.Println("received from chan1", val)
	case val := <-chan2:
		fmt.Println("received from chan2", val)
	}
}
```

- `select` listens on **multiple channels** at once — whichever is ready first gets handled
- Like a `switch` but for channels

```bash
go run multi_chan.go
```
