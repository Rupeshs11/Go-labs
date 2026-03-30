# 🩺 Web Health Checker

A lightning-fast Command Line Interface (CLI) tool written in **Go** to check if a specific website or domain is up and reachable.

---

## 🛠 Description

The **Web Health Checker** uses raw TCP connections to ping a domain and port, verifying if the target server is listening and actively accepting connections. It's built on top of the popular [`urfave/cli/v2`](https://github.com/urfave/cli) framework, making the command-line experience intuitive and clean.

---

## 💻 Code Explanations

### CLI Setup (`main.go`)

```go
app := &cli.App{
    Name:  "Healthchecker",
    Usage: "A tiny tool that checks the given domain is down.",
    Flags: []cli.Flag{
        &cli.StringFlag{
            Name:     "domain",
            Aliases:  []string{"d"},
            Usage:    "Domain name to check.",
            Required: true,
        },
        &cli.StringFlag{
            Name:     "port",
            Aliases:  []string{"p"},
            Required: false,
        },
    },
    // ...
}
```

- Defines required and optional flag arguments. The `--domain` (`-d`) flag is mandatory, while `--port` (`-p`) is optional (defaulting to port `80` if omitted).

### Core TCP Check Logic (`check.go`)

```go
func Check(destination string, port string) string {
    address := destination + ":" + port
    timeout := time.Duration(5 * time.Second)

    conn, err := net.DialTimeout("tcp", address, timeout)

    if err != nil {
        return fmt.Sprintf("[DOWN] %v is unreachable,\n Error: %v", destination, err)
    }

    return fmt.Sprintf("[UP] %v is reachable,\n From: %v\n To: %v", destination, conn.LocalAddr(), conn.RemoteAddr())
}
```

- **`net.DialTimeout`**: Initiates a TCP connection to the specified address. It waits up to 5 seconds before timing out, preventing the tool from hanging indefinitely on dead servers.
- Returns detailed results including routing IPs (`LocalAddr` and `RemoteAddr`) if successful.

---

## 🚀 How to Run

1. Clone the repository and navigate to the project directory:

   ```bash
   cd Projects/web_health_checker
   ```

2. Download dependencies (the `urfave/cli` framework):

   ```bash
   go mod tidy
   ```

3. Run the application passing the domain to check:

   **Check a standard website (Defaults to Port 80):**

   ```bash
   go run . --domain google.com
   # OR
   go run . -d google.com
   ```

   **Check a specific port (e.g., HTTPS Port 443):**

   ```bash
   go run . -d github.com --port 443
   # OR
   go run . -d github.com -p 443
   ```

---

## Dependencies

| Package                                          | Purpose                                                            |
| ------------------------------------------------ | ------------------------------------------------------------------ |
| [`urfave/cli/v2`](https://github.com/urfave/cli) | Fast, simple, and expressive framework for building CLI apps in Go |
