# 🐹 Go Language — Learning Notes

A hands-on collection of Go (Golang) learning examples, progressing from basics to structs and interfaces. Each folder contains working Go code with a `README.md` explaining the topic.

---

## 📚 Topics Covered

| #   | Folder                                        | Topic                 | Key Concepts                                        |
| --- | --------------------------------------------- | --------------------- | --------------------------------------------------- |
| 01  | [Hello World](./01_Hello_world/)              | First Go program      | `package main`, `fmt.Println`, `go run`, `go build` |
| 02  | [Simple Values](./02_simple_values/)          | Basic data types      | `int`, `float64`, `string`, `bool`                  |
| 03  | [Variables](./03_variables/)                  | Variable declarations | `var`, `:=` shorthand, type inference, zero values  |
| 04  | [Constants](./04_constants/)                  | Constants             | `const`, grouped constants, compile-time values     |
| 05  | [Loops](./05_Loops/)                          | Loops                 | `for`, `range`, while-style, infinite loop          |
| 06  | [Conditionals](./06_Conditionals/)            | If / Else             | `if/else`, init statements, no ternary operator     |
| 07  | [Switch Statements](./07_switch-statements/)  | Switch                | Basic switch, multi-value cases, type switch        |
| 08  | [Arrays](./08_Arrays/)                        | Arrays                | Fixed-size arrays, 2D arrays, zero values           |
| 09  | [Slices](./09_slices/)                        | Slices                | `make`, `append`, `copy`, slice operator            |
| 10  | [Maps](./10_Maps/)                            | Maps                  | Key-value pairs, `delete`, comma-ok idiom           |
| 11  | [Range](./11_Range/)                          | Range                 | Iterating slices, maps, strings with `range`        |
| 12  | [Functions](./12_Functions/)                  | Functions             | Multiple returns, `fmt.Scan`, first-class functions |
| 13  | [Variadic Functions](./13_variadic_function/) | Variadic Functions    | `...` syntax, spreading slices                      |
| 14  | [Closures](./14_closures/)                    | Closures              | Functions with persistent state                     |
| 15  | [Pointers](./15_pointers/)                    | Pointers              | `&`, `*`, pass by value vs reference                |
| 16  | [Structs](./16_Structs/)                      | Structs               | Methods, constructors, embedding (composition)      |
| 17  | [Interfaces](./17_Interface/)                 | Interfaces            | Implicit implementation, polymorphism               |

---

## 🚀 Getting Started

### Prerequisites

- **Go** installed (version latest OR 1.22+ recommended) — [Download Go](https://go.dev/dl/)

### Running Examples

```bash
cd learnings/01_Hello_world
go run main.go
```

> **Note:** Some files have code commented out (used as reference/notes). Uncomment to run them.

---

## 📝 Structure

```
learnings/
├── 01_Hello_world/           # First program, go run & go build
├── 02_simple_values/         # int, float, string, bool
├── 03_variables/             # var, :=, type inference
├── 04_constants/             # const, grouped blocks
├── 05_Loops/                 # for, while-style, range
├── 06_Conditionals/          # if/else, init statements
├── 07_switch-statements/     # switch, multi-case, type switch
├── 08_Arrays/                # Fixed-size arrays, 2D arrays
├── 09_slices/                # Dynamic slices, append, copy
├── 10_Maps/                  # Key-value pairs, delete, compare
├── 11_Range/                 # Iterating with range
├── 12_Functions/             # Multiple returns, first-class funcs
├── 13_variadic_function/     # ... syntax, spread operator
├── 14_closures/              # Functions with persistent state
├── 15_pointers/              # & and * operators, pass by ref
├── 16_Structs/               # Methods, constructors, embedding
└── 17_Interface/             # Implicit interfaces, polymorphism
```
