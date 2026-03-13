# 16 — Structs

Go has **no classes**. Instead, it uses **structs** to group related data, and **methods** to add behavior. Structs + methods give you everything you need for OOP-style programming.

---

## `struct.go` — Defining & Using Structs (commented reference)

```go
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

myorder := order{
	id:     "1",
	amount: 50,
	status: "received",
}
myorder.createdAt = time.Now()

fmt.Println(myorder)           // full struct
fmt.Println(myorder.status)    // access a specific field
```

- Fields are accessed with **dot notation**: `myorder.status`
- `time.Time` is from the `time` package — gives nanosecond precision

---

## `method.go` — Methods on Structs (commented reference)

Methods are functions with a **receiver** — they belong to a struct.

```go
// pointer receiver (*order) — can modify the struct
func (o *order) changeStatus(status string) {
	o.status = status
}

// pointer receiver — just reading, but still using pointer to avoid copy
func (o *order) getAmount() float32 {
	return o.amount
}

myOrder := order{id: "1", amount: 100, status: "received"}
myOrder.changeStatus("confirmed")
fmt.Println(myOrder.getAmount())    // 100
```

> **Rule of thumb:** Use `*` (pointer receiver) when you want to **modify** the struct. Use value receiver only for read-only operations on small structs.

---

## `constructor.go` — Constructor Pattern (commented reference)

Go has no built-in constructors, but the convention is a `newXxx` function that returns a pointer:

```go
func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder    // return pointer to the created struct
}

myOrder := newOrder("1", 30.50, "received")
fmt.Println(myOrder.getAmount())    // 30.5
```

- Returns `*order` (pointer) so the caller gets a reference, not a copy
- `&myOrder` takes the address of the local variable

---

## `embedding.go` — Struct Embedding (Composition)

Go uses **composition over inheritance**. You embed one struct inside another:

```go
type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer                // embedded — no field name, just the type
}

newOrder := order{
	id:     "1",
	status: "received",
	amount: 25.0,
	customer: customer{
		name:  "john",
		phone: "123456789",
	},
}

newOrder.customer.name = "knox"      // access embedded fields
fmt.Println(newOrder.customer)       // {knox 123456789}
```

- Embedding gives you direct access to the inner struct's fields
- This is Go's way of achieving **composition** (instead of inheritance)

```bash
go run embedding.go
```
