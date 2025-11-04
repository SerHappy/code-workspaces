# cw - VSCode Workspace Launcher

A fast CLI tool for opening VS Code workspaces by their relative path.

## What is it?

`cw` (code workspace) is a command-line utility that helps you quickly open VS Code workspace files without typing full paths. It scans your projects directory for `.code-workspace` files and lets you open them using just the relative directory name.

Instead of:
```bash
code ~/Projects/my-company/backend/api-service/api.code-workspace
```

You can simply type:
```bash
cw my-company/backend/api-service
```

## Why use it?

- **Faster than VS Code extensions**: VS Code workspace picker extensions can be slow, especially with many workspaces. Searching through a long list visually takes more time than simply typing the path you need.
- **Fast workspace switching**: Open any workspace with a short command
- **Shell autocompletion**: Tab-complete workspace paths for even faster navigation
- **Simple workflow**: No need to remember full paths or navigate through directories
- **Lightweight**: Single binary with no dependencies

## Installation

### Prerequisites

- Go 1.21 or higher
- VS Code with `code` command in PATH

### Install from source

```bash
git clone https://github.com/serhappy/code-workspaces.git
cd code-workspaces
make install
```

This will build and install the `cw` binary to `~/go/bin/`. Make sure this directory is in your `PATH`.

### Uninstall

```bash
make uninstall
```

## Setup

By default, `cw` scans for workspaces under `~/Projects`. The tool recursively searches for `.code-workspace` files in this directory.

### Directory structure example

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

With this structure, you can open workspaces using:
- `cw personal/blog`
- `cw work/frontend`
- `cw work/backend`

## Shell Autocompletion

Enable autocompletion for faster workspace navigation.

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

After adding the completion script, restart your shell or source the config file:
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
```

This opens the workspace file found in `~/Projects/work/backend/`.

### List all available workspaces

```bash
cw list
```

This displays all workspace keys (relative paths) that you can use with the `cw` command.

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
