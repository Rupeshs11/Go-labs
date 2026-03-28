```text
██╗  ██╗███╗   ██╗██████╗ ██╗  ██╗██████╗ ██████╗
██║ ██╔╝████╗  ██║██╔═══██╗╚██╗██╔╝██╔══██╗██╔══██╗
█████╔╝ ██╔██╗ ██║██║   ██║ ╚███╔╝ ██║  ██║██████╔╝
██╔═██╗ ██║╚██╗██║██║   ██║ ██╔██╗ ██║  ██║██╔══██╗
██║  ██╗██║ ╚████║╚██████╔╝██╔╝ ██╗██████╔╝██████╔╝
╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚═╝  ╚═╝╚═════╝ ╚═════╝
```

# KnoxDB

A lightweight, thread-safe, JSON-file-based NoSQL database written in **Go**. KnoxDB stores data as structured JSON files inside directories (representing collections), offering a simple alternative to full-scale databases for small projects.

---

## 🛠 Features

- **Store Data as JSON**: Each record is saved as a pretty-printed `.json` file for easy readability.
- **Thread-safe**: Utilizes `sync.Mutex` at the collection level to prevent concurrent write collisions.
- **Collections & Resources**: Emulates a NoSQL structure (`Collection -> Resource`).
- **Standard DB Operations**: Supports `Write`, `Read`, `ReadAll`, and `Delete`.
- **Custom Logger**: Integrated with `github.com/jcelliott/lumber` for console logging.

---

## 🏗 Storage Architecture

KnoxDB uses the filesystem to structure data. When you insert a record, the folder structure looks like this:

```
YOUR_DB_DIR/
└── users/                     <-- Collection
    ├── Knox.json              <-- Resource
    ├── Luffy.json
    └── Tonystark.json
```

---

## 💻 Code Explanations

### Initializing the Database

```go
db, err := New("./", nil)
if err != nil {
    fmt.Println("Error", err)
}
```

- `New(dir, options)` checks if the database directory exists. If not, it creates it.
- Returns a `Driver` instance containing the base directory, a map of mutexes (for thread safety), and a logger.

### Writing a Record (`Write`)

```go
db.Write("users", "Knox", User{
    Name:    "Knox",
    Age:     "11",
    Contact: "1234567890",
    // ...
})
```

- Accepts a `collection` name {"users"}, a `resource` name {"Knox"}, and any struct/value `v`.
- Creates a temporary file first (`.tmp`), writes the marshaled JSON, and then safely renames it to `.json` to avoid corruption during crashes.
- Uses collection-level mutexes to lock writing operations.

### Reading Data (`Read` & `ReadAll`)

```go
records, err := db.ReadAll("users")
```

- `Read(collection, resource, &v)`: Checks if the `.json` file exists, reads it, and unmarshals it back into the provided struct pointer.
- `ReadAll(collection)`: Iterates over all files in the collection directory, reads them, and returns an array of JSON string records which can then be unmarshaled.

### Deleting Records (`Delete`)

```go
db.Delete("users", "Knox") // Deletes a specific resource
db.Delete("users", "")     // Deletes the entire collection
```

- `Delete` can seamlessly handle both `.json` files (removing a single resource) and directories (dropping an entire collection).

---

## 🚀 How to Run

1. Clone the repository and navigate to the project directory:

   ```bash
   cd Projects/KnoxDB
   ```

2. Download dependencies (the lumber logger):

   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

Once run, you will be presented with an interactive CLI menu to perform database operations:

```text
--- KnoxDB Operations Menu ---
1. Add initial dummy employees
2. Write a new user
3. Read a user
4. Read all users
5. Delete a user
6. Exit
Enter your choice:
```

The application handles your input using a `bufio.NewReader`, executing the selected database operation dynamically through a `switch` statement! After performing operations like adding a user, you can check the generated `users/` folder to see the JSON files showcasing the power of **KnoxDB**!
