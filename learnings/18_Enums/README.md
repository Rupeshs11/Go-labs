# 18 — Enums

Go has **no built-in enum keyword**. Instead, enums are implemented using `const` blocks with a custom type. This gives you type safety — you can't pass a random string where an enum is expected.

## How to Create Enums in Go

```go
type OrderStatus string       // 1. Define a custom type

const (                       // 2. Define the enum values
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)
```

## Code — `enums.go`

```go
type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func main() {
	changeOrderStatus(Delivered)    // ✅ works
	// changeOrderStatus("random")  // ❌ won't work — type safety
}
```

- The function accepts `OrderStatus` type, not just any string — this prevents invalid values
- `iota` can be used for integer-based enums (auto-incrementing from 0):
  ```go
  const (
      Received OrderStatus = iota   // 0
      Confirmed                     // 1
      Prepared                      // 2
  )
  ```

```bash
go run enums.go
# Output: changing order status to delivered
```
