![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)
![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue)
![Platform](https://img.shields.io/badge/Platform-macOS%20%7C%20Linux-lightgrey)

# cw - VSCode Workspace Launcher

A fast and lightweight CLI tool for opening VS Code workspaces by their relative paths.  
Designed for developers who want to jump between projects instantly — without touching the mouse.

## Table of Contents

- [What is it?](#what-is-it)
- [Features](#features)
- [Why use it?](#why-use-it)
- [Installation](#installation)
  - [Prerequisites](#prerequisites)
  - [Install from source](#install-from-source)
  - [Install via Go](#install-via-go)
  - [Uninstall](#uninstall)
- [Setup](#setup)
- [Shell Autocompletion](#shell-autocompletion)
  - [Bash](#bash)
  - [Zsh](#zsh)
  - [Fish](#fish)
- [Usage](#usage)
  - [Open a workspace](#open-a-workspace)
  - [List all available workspaces](#list-all-available-workspaces)
  - [Get help](#get-help)
- [How it works](#how-it-works)
  - [Workspace detection logic](#workspace-detection-logic)
  - [Ignored directories](#ignored-directories)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)

## What is it?

`cw` (short for **code workspace**) is a command-line tool that lets you open VS Code workspaces using short, relative paths — no more typing full directories or navigating folders.

You can simply type:
```bash
cw my-company/backend/api-service
```

## Features

- ⚡ Instant workspace switching — open any project instantly with `cw <path>`
- 🔍 Smart recursive scanning — finds `.code-workspace` files under `~/Projects`
- 🧠 Leaf detection — stops scanning once a workspace is found inside a directory
- 🪄 Shell autocompletion — works with Bash, Zsh, and Fish
- 💻 Lightweight & dependency-free — just one self-contained Go binary

## Why use it?

- **Faster than VS Code extensions** — workspace picker extensions can lag when you have many projects. Typing is faster than scrolling.
- **Minimal friction** — jump to any workspace from the terminal.
- **Simple workflow** — no need to remember full paths or navigate through directories.

## Installation

### Prerequisites

- Go 1.21 or higher
- VS Code with the `code` command available in `PATH`

### Install from source

```bash
git clone https://github.com/serhappy/code-workspaces.git
cd code-workspaces
make install
```

This builds and installs the `cw` binary to `~/go/bin/`.
Ensure that `~/go/bin` is in your `PATH`:
```bash
export PATH="$HOME/go/bin:$PATH"
```

### Install via Go

You can also install directly using Go:
```bash
go install github.com/serhappy/code-workspaces/cmd/cw@latest
```

### Uninstall

```bash
make uninstall
```

## Setup

By default, `cw` scans for workspaces under `~/Projects`. The tool recursively searches for `.code-workspace` files in this directory.

### Example directory structure

```
~/Projects/
├── personal/
│   └── blog/
│       └── blog.code-workspace
├── work/
│   ├── frontend/
│   │   └── app.code-workspace
│   └── backend/
│       └── api.code-workspace
```

Usage examples:
- `cw personal/blog`
- `cw work/frontend`
- `cw work/backend`

## Shell Autocompletion

Enable autocompletion to quickly navigate between workspaces.

### Bash

Add to your `~/.bashrc`:
```bash
eval "$(cw completion bash)"
```

### Zsh

Add to your `~/.zshrc`:
```bash
eval "$(cw completion zsh)"
```

### Fish

Add to your `~/.config/fish/config.fish`:
```fish
cw completion fish | source
```

Then reload your shell:
```bash
source ~/.bashrc  # or ~/.zshrc
```

## Usage

### Open a workspace

```bash
cw <relative-directory>
```

Example:
```bash
cw work/backend
# Opens ~/Projects/work/backend/api.code-workspace in the current VS Code window
```

### List all available workspaces

```bash
cw list
```
This displays all relative workspace paths detected by `cw`.

### Get help

```bash
cw --help
cw list --help
```

## How it works

1. `cw` scans the `~/Projects` directory recursively
2. It finds all `.code-workspace` files
3. Each workspace is indexed by its relative directory path
4. When you run `cw <path>`, it opens the corresponding workspace in VS Code using `code -r` (reuse window)

### Workspace detection logic

**Important**: When `cw` finds a directory containing a `.code-workspace` file, it treats that directory as a **leaf node** and stops scanning deeper. This means:

- If `~/Projects/my-project/` contains `project.code-workspace`, the scan stops there
- Any subdirectories under `my-project/` are ignored, even if they contain their own workspace files
- This design assumes each workspace directory is a complete, self-contained project

Example:
```
~/Projects/
├── frontend/
│   ├── app.code-workspace          ← Found! Stop scanning here
│   └── nested/
│       └── ignored.code-workspace  ← Never scanned
└── backend/
    └── api.code-workspace          ← Found! Stop scanning here
```

Result: Only `frontend` and `backend` workspaces are available.

### Ignored directories

The following directories are automatically excluded from scanning:
- `.links`
- `python_wrappers`

## Roadmap

Planned features:

- [ ] Configurable root directory
- [ ] Custom exclude patterns
- [ ] Workspace aliases
- [ ] Configuration file support
- [ ] Multiple workspace roots

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

MIT License - see LICENSE file for details.
