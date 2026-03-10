# 04 — Constants

Constants are values that **cannot change** once defined. They're evaluated at **compile time**, making them ideal for config values, fixed numbers, etc.

## Key Rules

- Declared with `const` keyword (not `var`)
- **Cannot** be reassigned — `const x = 5; x = 10` → ❌ compile error
- **Cannot** use `:=` shorthand for constants
- Can be declared **inside or outside** functions
- Can be **grouped** together using `const ( ... )` block

## Code — `constants.go`

```go
// grouped constants — declared outside main (package level)
const (
	port  = 3000
	price = "33.3"
	host  = "localhost"
	os    = "Amazon-linux 2023"
)

func main() {
	const name = "golang"    // constant inside a function

	// name = "knox"         // ❌ would cause compile error — can't reassign constants

	fmt.Println(name)        // golang
	fmt.Println(port)        // 3000
	fmt.Println(price, host, os)  // 33.3 localhost Amazon-linux 2023
}
```

- Grouped `const ()` blocks are great for related configuration values
- `price` is `"33.3"` (a string), not a number — notice the quotes

```bash
go run constants.go
```
