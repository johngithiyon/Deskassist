````md
# DeskAssist

A lightweight terminal-based desktop assistant for Linux built with Go.

DeskAssist allows users to interact with their system using simple natural language commands instead of remembering Linux commands.

It can:
- Open files
- Open folders
- Launch applications
- Open websites
- Perform Google searches directly from the terminal

---

# Features

- Open folders in file manager
- Open files with default applications
- Launch applications like browser or VS Code
- Perform Google searches
- Simple natural language command handling
- Lightweight and fast
- Built entirely in Go

---

# Demo

```bash
deskassist open downloads
deskassist open vscode
deskassist open notes.txt
deskassist search kubernetes ingress
deskassist google golang tutorials
```

---

# Project Structure

```text
deskassist/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── commands/
│   ├── parser/
│   ├── executor/
│   └── config/
│
├── go.mod
├── go.sum
└── README.md
```

---

# Requirements

- Linux
- Go 1.22+
- Visual Studio Code (optional)
- xdg-open

---

# Installation

## Clone Repository

```bash
git clone <your-repository-url>
```

```bash
cd deskassist
```

---

# Build the Project

```bash
go build -o deskassist ./cmd
```

This creates a binary named:

```text
deskassist
```

---

# Run the Assistant

```bash
./deskassist
```

Or install globally:

```bash
sudo mv deskassist /usr/local/bin/
```

Now you can run:

```bash
deskassist
```

from anywhere in the terminal.

---

# Usage

## Open Folder

```bash
deskassist open downloads
```

---

## Open File

```bash
deskassist open resume.pdf
```

---

## Open Browser

```bash
deskassist open firefox
```

---

## Open VS Code Project

```bash
deskassist open nodefy
```

---

## Google Search

```bash
deskassist search kubernetes networking
```

or

```bash
deskassist google golang concurrency
```

---

# Example Commands

| Command | Action |
|----------|---------|
| `deskassist open downloads` | Opens Downloads folder |
| `deskassist open vscode` | Opens VS Code |
| `deskassist search docker` | Searches Google |
| `deskassist open github.com` | Opens website |
| `deskassist open notes.txt` | Opens file |

---

# How It Works

DeskAssist converts natural language terminal input into Linux system commands.

Example:

```text
deskassist search golang
```

Internally executes:

```bash
xdg-open "https://google.com/search?q=golang"
```

---

# Technologies Used

- Go
- Linux Process Execution
- os/exec
- xdg-open

---

# Core Concepts Learned

This project demonstrates:
- Process management
- Linux command execution
- CLI application development
- Natural language parsing
- System automation
- File handling
- Browser automation

---

# Future Improvements

- Voice commands
- AI-powered intent recognition
- File indexing
- Fuzzy search
- Background daemon
- Custom aliases
- Plugin system
- Multi-platform support

---

# Example Internal Flow

```text
User Command
      ↓
Parser
      ↓
Intent Detection
      ↓
Command Executor
      ↓
Linux System
```

---

# Run in Development

```bash
go run ./cmd/main.go
```

---

# Example Code

```go
cmd := exec.Command(
    "xdg-open",
    "https://google.com/search?q=golang",
)

err := cmd.Run()
```

---

# Contributing

Contributions are welcome.

You can improve:
- command parsing
- AI integration
- desktop integration
- plugin support
- Linux automation features

---

# License

MIT License

---

# Author

John Githiyon
````
