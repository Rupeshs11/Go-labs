# 01 — Hello World 👋

The classic first program — prints "Hello world" to the console.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}
```

## How to Run

```bash
# Run directly (compiles + runs in one step, no binary created)
go run main.go

# Output: Hello world
```

## How to Build

```bash
# Compile into a binary executable
go build main.go

# This creates a binary file:
#   - On Linux/Mac → ./main
#   - On Windows   → main.exe

# Now run the binary directly
./main          # Linux/Mac
main.exe        # Windows

# Output: Hello world
```

> `go run` = compile + run (for development). `go build` = compile only (creates a binary you can distribute/deploy).
