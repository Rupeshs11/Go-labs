# 23 — File Handling

Go's `os` package gives you everything to work with files — create, read, write, delete, and get file info.

## Common Functions

| Function              | What it does                                                       |
| --------------------- | ------------------------------------------------------------------ |
| `os.Open("file")`     | Open file for reading                                              |
| `os.Create("file")`   | Create or truncate a file for writing                              |
| `os.ReadFile("file")` | Read entire file into memory (simple but not best for large files) |
| `os.Remove("file")`   | Delete a file                                                      |
| `f.Stat()`            | Get file info (name, size, permissions, modified time)             |
| `f.Close()`           | Close the file — always use `defer f.Close()`                      |

---

## `files.go` — Getting File Info (commented reference)

```go
f, err := os.Open("example.txt")
if err != nil {
	panic(err)
}

fileInfo, _ := f.Stat()
fmt.Println("File name:", fileInfo.Name())
fmt.Println("Is folder:", fileInfo.IsDir())
fmt.Println("Size:", fileInfo.Size())
fmt.Println("Permissions:", fileInfo.Mode())
fmt.Println("Modified at:", fileInfo.ModTime())
```

---

## `ReadFile.go` — Reading Files (commented reference)

**Method 1 — Manual read with buffer:**

```go
f, _ := os.Open("example.txt")
defer f.Close()

buf := make([]byte, 50)        // read 50 bytes at a time
f.Read(buf)
```

**Method 2 — Read entire file (simpler):**

```go
data, _ := os.ReadFile("example.txt")
fmt.Println(string(data))
```

> ⚠️ `os.ReadFile` loads the **entire file** into memory — not ideal for very large files.

---

## `ReadFolder.go` — Listing Directory Contents (commented reference)

```go
dir, _ := os.Open(".")
defer dir.Close()

files, _ := dir.ReadDir(3)     // read up to 3 entries
for _, fi := range files {
	fmt.Println(fi.Name())
}
```

---

## `Write.go` — Writing to Files (commented reference)

**Create and write strings:**

```go
f, _ := os.Create("example.txt")
defer f.Close()

f.WriteString("Hello Golang")

// or write bytes
bytes := []byte("Hello Golang")
f.Write(bytes)
```

**Copy file contents (streaming — good for large files):**

```go
sourceFile, _ := os.Open("source.txt")
destFile, _ := os.Create("destination.txt")
defer sourceFile.Close()
defer destFile.Close()

reader := bufio.NewReader(sourceFile)
writer := bufio.NewWriter(destFile)

for {
	b, err := reader.ReadByte()
	if err != nil {
		break                     // EOF reached
	}
	writer.WriteByte(b)
}
writer.Flush()                    // don't forget to flush!
```

- `bufio` reads/writes byte-by-byte — memory efficient for large files
- `writer.Flush()` is important — it pushes buffered data to the actual file

---

## `delete.go` — Deleting Files

```go
err := os.Remove("example2.txt")
if err != nil {
	panic(err)
}
fmt.Println("File deleted successfully...!")
```

```bash
go run delete.go
```
