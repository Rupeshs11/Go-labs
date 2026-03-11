# 🐹 Go Language — Learning Notes

A hands-on collection of Go (Golang) learning examples, progressing from the very basics to more advanced data structures. Each folder contains working Go code with a detailed `README.md` explaining the topic and code.

---

## 📚 Topics Covered

| #   | Folder                                       | Topic                     | Key Concepts                                             |
| --- | -------------------------------------------- | ------------------------- | -------------------------------------------------------- |
| 01  | [Hello World](./01_Hello_world/)             | First Go program          | `package main`, `fmt.Println`, `go run`, `go build`      |
| 02  | [Simple Values](./02_simple_values/)         | Basic data types          | `int`, `float64`, `string`, `bool`                       |
| 03  | [Variables](./03_variables/)                 | Variable declarations     | `var`, `:=` shorthand, type inference, zero values       |
| 04  | [Constants](./04_constants/)                 | Constants                 | `const`, grouped constants, compile-time values          |
| 05  | [Loops](./05_Loops/)                         | Loops (for, while, range) | `for`, `range`, while-style, infinite loop               |
| 06  | [Conditionals](./06_Conditionals/)           | If / Else                 | `if/else`, init statements, no ternary operator          |
| 07  | [Switch Statements](./07_switch-statements/) | Switch                    | Basic switch, multi-value cases, type switch             |
| 08  | [Arrays](./08_Arrays/)                       | Arrays                    | Fixed-size arrays, 2D arrays, zero values                |
| 09  | [Slices](./09_slices/)                       | Slices                    | `make`, `append`, `copy`, slice operator, `slices.Equal` |
---

## 🚀 Getting Started

### Prerequisites

- **Go** installed ( version latest OR version 1.22+ recommended) — [Download Go](https://go.dev/dl/)

### Running Examples

Navigate into any folder and run the Go file:

```bash
cd learnings/01_Hello_world
go run main.go
```

> **Note:** Some files have code commented out (used as reference/notes). Uncomment to run them.

---

## 📝 Structure

```
learnings/
├── 01_Hello_world/        # Your first Go program
├── 02_simple_values/      # Data types: int, float, string, bool
├── 03_variables/          # Variable declaration & type inference
├── 04_constants/          # Constants & grouped const blocks
├── 05_Loops/              # For loops, while-style, range
├── 06_Conditionals/       # If/else with init statements
├── 07_switch-statements/  # Switch, multi-case, type switch
├── 08_Arrays/             # Fixed-size arrays & 2D arrays
├── 09_slices/             # Dynamic slices, append, copy, compare
```
