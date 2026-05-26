<div align="center">

# DeskAssist

**A terminal-based AI assistant for Linux that helps users open files, launch applications, manage folders, and search the web using natural language commands.**

[![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![Linux](https://img.shields.io/badge/Linux-FCC624?logo=linux&logoColor=black)](https://kernel.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

</div>

---

## Overview

DeskAssist is a lightweight Linux desktop assistant built with Go that allows users to control their system using natural language commands directly from the terminal.

Instead of remembering complex Linux commands, users can simply type human-friendly instructions like:

```bash
openfile filename

openfolder foldername

openbrowser browsername

google searchvalue
```

DeskAssist converts user commands into Linux system operations and automates desktop workflows.

---

## Features

| Feature | Description |
|---|---|
| 📂 **Open Files & Folders** | Open files and directories directly from terminal commands |
| 🌐 **Google Search** | Search the web instantly using natural language |
| 💻 **Application Launcher** | Launch applications like Firefox,Chrome,Edge etc |
| ⚡ **Natural Language Commands** | Interact without memorizing Linux commands |

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
- [Linux Operating System](https://ubuntu.com/download/desktop)
- [Git](https://git-scm.com/downloads)

---

## Installation

### 1. Clone Repository

```bash
git clone https://github.com/johngithiyon/deskassist.git
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


## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

---

<div align="center">


</div>