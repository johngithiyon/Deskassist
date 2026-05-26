````md id="gk8d21"
<div align="center">

# 🚀 DeskAssist

### AI-Powered Terminal Assistant for Linux

<img src="https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go">
<img src="https://img.shields.io/badge/Linux-Supported-FCC624?style=for-the-badge&logo=linux">
<img src="https://img.shields.io/badge/Open%20Source-❤-red?style=for-the-badge">

---

### Open Files • Launch Apps • Search Google • Control Linux

A lightweight terminal-based assistant built with Go that allows users to interact with Linux using natural language commands instead of memorizing shell commands.

</div>

---

# ✨ Features

- 📂 Open folders instantly
- 📄 Open files with default applications
- 🌐 Search Google directly from terminal
- 💻 Launch applications like VS Code or Firefox
- ⚡ Fast and lightweight
- 🧠 Natural language command support
- 🐧 Built for Linux
- 🔧 Easy to extend and customize

---

# 🎬 Demo

```bash
deskassist open downloads

deskassist open vscode

deskassist search kubernetes networking

deskassist google golang concurrency

deskassist open notes.txt
```

---

# 🏗 Architecture

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

# 📁 Project Structure

```text
deskassist/
│
├── cmd/
│   └── deskassist/
│       └── main.go
│
├── internal/
│   ├── parser/
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

# ⚙ Requirements

- Linux
- Go 1.22+
- `xdg-open`
- Visual Studio Code (optional)

---

# 📦 Installation

## Clone Repository

```bash
git clone <your-repository-url>
```

```bash
cd deskassist
```

---

# 🔨 Build

```bash
go build -o deskassist ./cmd/deskassist
```

---

# 🚀 Run

```bash
./deskassist
```

---

# 🌍 Global Installation

Move the binary to system path:

```bash
sudo mv deskassist /usr/local/bin/
```

Now run from anywhere:

```bash
deskassist
```

---

# 💡 Usage

## 📂 Open Folder

```bash
deskassist open downloads
```

---

## 📄 Open File

```bash
deskassist open resume.pdf
```

---

## 🌐 Google Search

```bash
deskassist search kubernetes ingress controller
```

---

## 💻 Launch VS Code

```bash
deskassist open vscode
```

---

## 🔥 Open Browser

```bash
deskassist open firefox
```

---

# 🧠 Example Internal Execution

User Command:

```bash
deskassist search golang tutorials
```

Internally Executes:

```bash
xdg-open "https://google.com/search?q=golang+tutorials"
```

---

# 🛠 Technologies Used

| Technology | Purpose |
|------------|----------|
| Go | Core language |
| os/exec | Process execution |
| xdg-open | Open browser/files |
| Linux | Operating system |
| CLI Parsing | Natural language commands |

---

# 📚 Concepts Learned

This project teaches:

- Linux process management
- CLI application development
- Natural language parsing
- System automation
- File handling
- Browser automation
- OS-level command execution
- Clean Go architecture

---

# 🔮 Future Improvements

- 🎤 Voice commands
- 🤖 AI-powered intent recognition
- ⚡ Background daemon
- 🧩 Plugin system
- 📑 File indexing
- 🔍 Fuzzy search
- 🌐 Multi-platform support
- 🧠 Local LLM integration

---

# 🧪 Development

Run directly:

```bash
go run ./cmd/deskassist/main.go
```

---

# 💻 Example Go Code

```go
package main

import (
	"net/url"
	"os/exec"
)

func main() {

	query := url.QueryEscape("golang tutorial")

	link := "https://google.com/search?q=" + query

	cmd := exec.Command("xdg-open", link)

	cmd.Run()
}
```

---

# 🤝 Contributing

Contributions are welcome.

Ideas:
- Better NLP parser
- Voice assistant support
- Plugin SDK
- Linux integrations
- AI features
- Smart file indexing

---

# 📜 License

MIT License

---

<div align="center">

# 👨‍💻 Author

### John Githiyon

Built with Go ❤️ and Linux 🐧

</div>
````
