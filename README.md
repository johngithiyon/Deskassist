<div align="center">

# DeskAssist

**A terminal-based AI assistant for Linux that helps users open files, launch applications, manage folders, and search the web using natural language commands.**

[![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![Linux](https://img.shields.io/badge/Linux-FCC624?logo=linux&logoColor=black)](https://kernel.org/)
[![VSCode](https://img.shields.io/badge/VSCode-007ACC?logo=visualstudiocode&logoColor=white)](https://code.visualstudio.com/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

</div>

---

## Overview

DeskAssist is a lightweight Linux desktop assistant built with Go that allows users to control their system using natural language commands directly from the terminal.

Instead of remembering complex Linux commands, users can simply type human-friendly instructions like:

```bash
deskassist open downloads

deskassist search kubernetes networking

deskassist open vscode

deskassist find my resume
```

DeskAssist converts user commands into Linux system operations and automates desktop workflows.

---

## Features

| Feature | Description |
|---|---|
| 📂 **Open Files & Folders** | Open files and directories directly from terminal commands |
| 🌐 **Google Search** | Search the web instantly using natural language |
| 💻 **Application Launcher** | Launch applications like Firefox, VS Code, Terminal, etc |
| ⚡ **Natural Language Commands** | Interact without memorizing Linux commands |
| 🔍 **File Finder** | Search files and folders quickly |
| 🐧 **Linux Automation** | Automate common desktop operations |
| 🧠 **Extensible Architecture** | Easily add new commands and features |

---

## Workflow

### 1️ User Runs DeskAssist

Users interact directly from the terminal using simple commands.

```bash
deskassist open downloads
```

---

### 2️ Parse Natural Language

DeskAssist parses the user input and detects the action and target.

```text
Action → open
Target → downloads
```

---

### 3️ Intent Detection

The assistant determines what operation needs to be performed.

Examples:
- Open Folder
- Open File
- Launch App
- Search Google

---

### 4️ Execute Linux Command

DeskAssist converts the command into Linux system execution.

Example:

```bash
xdg-open ~/Downloads
```

---

### 5️ Return Output

The requested application, folder, browser, or file is opened automatically.

---

## Example Commands

```bash
deskassist open downloads

deskassist open vscode

deskassist open github.com

deskassist search golang concurrency

deskassist find notes.txt

deskassist open nodefy backend
```

---

## Architecture

```text
                ┌─────────────────┐
                │   CLI Input     │
                └────────┬────────┘
                         │
                         ▼
                ┌─────────────────┐
                │ Command Parser  │
                └────────┬────────┘
                         │
                         ▼
                ┌─────────────────┐
                │ Intent Engine   │
                └────────┬────────┘
                         │
          ┌──────────────┼──────────────┐
          ▼              ▼              ▼
   ┌──────────┐   ┌──────────┐   ┌──────────┐
   │ File Ops │   │ Browser  │   │ App Exec │
   └────┬─────┘   └────┬─────┘   └────┬─────┘
        │              │              │
        ▼              ▼              ▼
 Linux Filesystem   Google/Open    VSCode/Apps
```

---

## Project Structure

```text
deskassist/
│
├── cmd/
│   └── deskassist/
│       └── main.go
│
├── internal/
│   ├── parser/
│   ├── intent/
│   ├── executor/
│   ├── commands/
│   ├── config/
│   └── logger/
│
├── pkg/
│
├── go.mod
├── go.sum
└── README.md
```

---

## Tech Stack

| Layer | Technology |
|---|---|
| **Language** | Go (Golang) |
| **Operating System** | Linux |
| **Process Execution** | os/exec |
| **Browser Integration** | xdg-open |
| **CLI Interface** | Terminal |

---

## Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install) (v1.22+)
- Linux Operating System
- `xdg-open`
- [Git](https://git-scm.com/downloads)

---

## Installation

### 1. Clone Repository

```bash
git clone https://github.com/yourusername/deskassist.git
```

```bash
cd deskassist
```

---

### 2. Install Dependencies

```bash
go mod tidy
```

---

### 3. Build the Application

```bash
go build -o deskassist ./cmd/deskassist
```

---

### 4. Move Binary Globally

```bash
sudo mv deskassist /usr/local/bin/
```

Now you can run:

```bash
deskassist
```

from anywhere in the terminal.

---

## Running the Application

```bash
deskassist open downloads
```

---

## Development Mode

```bash
go run ./cmd/deskassist/main.go
```

---

## Internal Execution Example

User Command:

```bash
deskassist search golang tutorial
```

Internally Executes:

```bash
xdg-open "https://google.com/search?q=golang+tutorial"
```

---

## Future Improvements

- 🎤 Voice Assistant
- 🤖 AI-powered command understanding
- 📂 Smart file indexing
- ⚡ Background daemon
- 🔍 Fuzzy search support
- 🧩 Plugin system
- 🌐 Multi-platform support
- 🧠 Local LLM integration

---

## Concepts Learned

DeskAssist demonstrates:

- Linux process management
- CLI application development
- Natural language parsing
- System automation
- Process execution
- File system interaction
- Browser automation
- Go project architecture

---

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

---

<div align="center">

### Built with Go ❤️ and Linux 🐧

</div>