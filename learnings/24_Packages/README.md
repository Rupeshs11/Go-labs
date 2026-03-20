# 24 — Packages

Packages let you **organize your code** into reusable, modular pieces. Every Go project uses packages — from the standard library (`fmt`, `os`) to your own custom ones and third-party libraries.

## Key Rules

| Rule                                  | Example                                                         |
| ------------------------------------- | --------------------------------------------------------------- |
| **Uppercase = Public** (exported)     | `func LoginWithCredentials()` ✅ accessible from other packages |
| **Lowercase = Private** (unexported)  | `func helper()` ❌ only accessible within the same package      |
| **One package per folder**            | `auth/` folder = `package auth`                                 |
| **`package main`** is the entry point | Must have a `func main()`                                       |

## Essential Commands

```bash
go mod init github.com/username/project   # initialize a module
go get github.com/fatih/color             # install third-party package
go mod tidy                               # clean up unused dependencies
```

---

## Project Structure

```
24_Packages/
├── main.go          # entry point (package main)
├── go.mod           # module definition & dependencies
├── go.sum           # dependency checksums
├── auth/
│   ├── credentials.go   # LoginWithCredentials()
│   └── session.go       # GetSession()
└── user/
    └── user.go          # User struct
```

---

## `auth/credentials.go` — Exported Function

```go
package auth

func LoginWithCredentials(username string, password string) {
	fmt.Println("login user using", username, password)
}
```

- `LoginWithCredentials` starts with **uppercase** → exported (public)
- Belongs to `package auth` — imported as `"github.com/rupeshs11/Go-labs/learnings/auth"`

## `auth/session.go` — Another File in Same Package

```go
package auth

func GetSession() string {
	return "loggedin"
}
```

- Same package (`auth`) — multiple files can share one package

## `user/user.go` — Exported Struct

```go
package user

type User struct {
	Email string
	Name  string
}
```

- Struct fields are **uppercase** → accessible from outside the package
- If `Email` was `email`, it would be private to the `user` package

---

## `main.go` — Putting It All Together

```go
package main

import (
	"fmt"
	"github.com/rupeshs11/Go-labs/learnings/auth"
	"github.com/rupeshs11/Go-labs/learnings/user"
	"github.com/fatih/color"    // third-party package
)

func main() {
	auth.LoginWithCredentials("Knox", "Joyboy123")

	session := auth.GetSession()
	fmt.Println("session:", session)

	user := user.User{
		Email: "user@email.com",
		Name:  "Monkey D. Luffy",
	}
	color.Magenta(user.Email)    // prints in magenta color
	color.Green(user.Name)       // prints in green color
}
```

- **Local packages** are imported by their full module path
- **Third-party packages** (like `fatih/color`) are installed with `go get` and used the same way
- `color.Magenta()` and `color.Green()` print colored text in the terminal

```bash
cd 24_Packages
go run main.go
```
