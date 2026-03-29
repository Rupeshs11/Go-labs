# ⏰ CLI Reminder

A lightweight, cross-platform Command Line Interface (CLI) application written in **Go** that allows you to set reminders using natural language time formats. The app automatically forks into a background process so you can keep using your terminal!

---

## 🛠 Features

- **Natural Language Parsing**: Set times using phrases like "in 5 minutes", "tomorrow at 2pm", or "next friday".
- **Background Execution**: Once scheduled, the reminder automatically runs in the background. It won't block your current terminal session.
- **Cross-Platform Native Popups**: Uses the `beeep` library for Linux/macOS and raw `syscall` functionality to trigger native `MessageBoxW` alerts on Windows.
- **Sound Alerts**: Beeps accompanied by visual alerts so you never miss a reminder.

---

## 💻 Code Explanations

### Natural Language Time Parsing

```go
w := when.New(nil)
w.Add(en.All...)
w.Add(common.All...)

t, err := w.Parse(os.Args[1], now)
```

- We rely on `github.com/olebedev/when` to intelligently parse the human-readable string passed via `os.Args[1]`.
- Enforces validity to ensure the scheduled time is in the future.

### Background Spawning (Forking)

```go
if os.Getenv(markName) == markValue {
    // 2. Background process: wait for time, then trigger popup!
    time.Sleep(diff)
    // ... display alert logic ...
} else {
    // 1. Initial process: spawn identical command as a child, setting the Env variable
    cmd := exec.Command(os.Args[0], os.Args[1:]...)
    cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))
    if err := cmd.Start(); err != nil {
        os.Exit(5)
    }
    fmt.Println("Reminder will be displayed after", diff.Round(time.Second))
    os.Exit(0) // Exit terminal context without killing child
}
```

- The app checks for a specific environment variable `GOLANG_CLI_REMINDER`.
- If it isn't set, it uses `exec.Command` to rerun itself as a background process with the environment variable attached, and then exits the main terminal prompt.
- If it is set, it executes the sleep timer and waits to trigger the alert!

### Windows Native Popups

```go
if runtime.GOOS == "windows" {
    user32 := syscall.NewLazyDLL("user32.dll")
    messageBox := user32.NewProc("MessageBoxW")
    // ...
    messageBox.Call(0, uintptr(unsafe.Pointer(msgPtr)), uintptr(unsafe.Pointer(titlePtr)), 0x00010040)
}
```

- Direct interop with the Windows API (`user32.dll`) via `syscall` and `unsafe.Pointer`.
- Creates a genuine Windows UI modal `MessageBoxW`.

---

## 🚀 How to Run

1. Clone the repository and navigate to the project directory:

   ```bash
   cd Projects/CLI-Reminder
   ```

2. Generate the binary (optional, but convenient):

   ```bash
   go build -o reminder.exe main.go
   ```

3. Run the application passing the `<time>` and `<message>`:
   ```bash
   ./reminder.exe "in 5 seconds" "Check the oven!"
   ./reminder.exe "next friday at 5pm" "Weekly sync meeting"
   ```

_(Notice that the command immediately returns your terminal prompt, while the process is silently waiting in the background to alert you!)_

## Dependencies

| Package                                       | Purpose                                                                        |
| --------------------------------------------- | ------------------------------------------------------------------------------ |
| [`when`](https://github.com/olebedev/when)    | A natural language date/time parser with pluggable rules                       |
| [`beeep`](https://github.com/gen2brain/beeep) | Go cross-platform library for sending desktop notifications, alerts, and beeps |
