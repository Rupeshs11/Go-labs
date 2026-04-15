# 💼 CRM (Customer Relationship Management) API

A fast, lightweight CRM backend built in **Go** utilizing the **Fiber** web framework and **GORM** (with a pure-Go SQLite driver). This API elegantly manages "Leads" via robust RESTful interfaces.

---

## 🛠 Features

- **Fiber Web Framework**: Handles HTTP routing gracefully with Express-like speeds using `github.com/gofiber/fiber`.
- **Pure-Go SQLite Integration**: Uses `github.com/glebarez/sqlite`—a CGO-free SQLite driver integrated flawlessly with GORM for zero-dependency local database management.
- **GORM Auto-Migrations**: Automatically spawns and syncs tables based on the Go structure (in this case, creating and managing a `leads.db` file).
- **CRUD Operations**: Secure endpoints explicitly defining mechanisms to create, read, target, and destroy client leads quickly.

---

## 💻 Code Architecture

The CRM logic is cleanly decoupled into specific domain packages (`database` and `lead`):

### 1. The Data Object (`lead.Lead`)

```go
type Lead struct {
    gorm.Model
    Name    string `json:"name"`
    Company string `json:"company"`
    Email   string `json:"email"`
    Phone   int    `json:"phone"`
}
```

- By embedding `gorm.Model`, GORM implicitly manages complex table features universally: `ID` (Primary Key), `CreatedAt`, `UpdatedAt`, and `DeletedAt` (Soft Deletes!).
- The JSON tags allow standard fiber `c.BodyParser()` techniques to map incoming payload bytes seamlessly to struct values.

### 2. GORM Auto-Migration (`initDatabase`)

```go
database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
database.DBConn.AutoMigrate(&lead.Lead{})
```

- GORM analyzes the `lead.Lead` struct and gracefully ensures the SQLite `leads.db` file contains parallel database constraints without writing raw SQL.

### 3. Fiber Controllers (`lead.go`)

```go
func NewLead(c *fiber.Ctx) {
    db := database.DBConn
    lead := new(Lead)
    if err := c.BodyParser(lead); err != nil {
        c.Status(503).Send(err)
        return
    }
    db.Create(&lead)
    c.JSON(lead)
}
```

- `BodyParser()` abstracts all JSON stream decoding.
- `db.Create(&lead)` binds directly into the database engine and safely builds the payload context before firing `c.JSON()` to deliver the finalized result back to the web client.

---

## 🚀 API Endpoints

| Method   | Route              | Description                                                                                                    |
| -------- | ------------------ | -------------------------------------------------------------------------------------------------------------- |
| `GET`    | `/api/v1/lead`     | Fetch an array payload of all Leads currently stored in the database.                                          |
| `GET`    | `/api/v1/lead/:id` | Fetch a specific Lead by their primary key `ID`.                                                               |
| `POST`   | `/api/v1/lead`     | Create a new lead. Send a formatted JSON body payload mapping the struct properties (`name`, `company`, etc.). |
| `DELETE` | `/api/v1/lead/:id` | Instructs the database to run a soft-delete on the target Lead via `ID`.                                       |

---

## 🏃 How to Run

1. Clone the repository and navigate to the project directory:

   ```bash
   cd Projects/CRM
   ```

2. Download all database and networking dependencies:

   ```bash
   go mod tidy
   ```

3. Fire up the backend:

   ```bash
   go run main.go
   # The server will acknowledge database connection and start listening on port 3000
   ```

4. Use Postman or `curl` to target `http://localhost:3000/api/v1/lead`!

---

## Dependencies

| Package                                                 | Purpose                                                                                 |
| ------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| [`gofiber/fiber`](https://github.com/gofiber/fiber)     | Express-inspired HTTP routing framework written natively in Go.                         |
| [`gorm.io/gorm`](https://gorm.io)                       | Fantastic Developer-Friendly ORM engine mapping struct structs to Relational databases. |
| [`glebarez/sqlite`](https://github.com/glebarez/sqlite) | Pure Go implementation of SQLite (no cumbersome CGO environment needed).                |
