# 17 — Interfaces

Interfaces define a **contract** — a set of method signatures that a type must implement. Any type that has all the methods automatically satisfies the interface. No `implements` keyword needed — it's **implicit**.

## Why Use Interfaces?

Without interfaces, if you want to switch from Razorpay to Stripe, you'd have to change code everywhere. With interfaces, you just swap the implementation — the rest of the code stays the same.

---

## Code — `interface.go`

### Step 1: Define the Interface

```go
type paymenter interface {
	pay(amount float32)
	refund(amount float32, id string)
}
```

Any type that has **both** `pay()` and `refund()` methods automatically implements `paymenter`.

### Step 2: Create Implementations

```go
type razorpay struct{}
func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount)
}

type stripe struct{}
func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)
}

type googlepay struct{}
func (g googlepay) pay(amount float32) {
	fmt.Println("making payment using Googlepay for", amount)
}
func (g googlepay) refund(amount float32, id string) {
	fmt.Println("initiating refund using googlepay of", amount, id)
}
```

> ⚠️ Only `googlepay` fully implements `paymenter` (has both `pay` and `refund`). The others only have `pay`.

### Step 3: Use the Interface

```go
type payment struct {
	gateway paymenter          // accepts ANY type that implements paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)      // calls whichever implementation is set
}

func (p payment) refundPayment(amount float32, id string) {
	p.gateway.refund(amount, id)
}
```

### Step 4: Swap Implementations Easily

```go
googlepayGw := googlepay{}
newPayment := payment{
	gateway: googlepayGw,      // swap to any gateway here
}
newPayment.makePayment(100)        // making payment using Googlepay for 100
newPayment.refundPayment(100, "11") // initiating refund using googlepay of 100 11
```

To switch payment providers, just change **one line** (`gateway: stripeGw`) — no other code changes needed.

### Key Takeaways

- Interfaces are **implicitly implemented** — no `implements` keyword
- A type satisfies an interface by having **all** its methods
- Use interfaces to make your code **flexible and testable** — notice `fakepayment` struct in the code for testing
- This is Go's version of **polymorphism**

```bash
go run interface.go
```
