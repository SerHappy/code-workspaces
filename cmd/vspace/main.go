package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
)

const workspaceExt = ".code-workspace"

var ignoreDirs = []string{".links", "python_wrappers"}

type Workspace struct {
	DirAbs  string
	FileAbs string
	RelDir  string
	Name    string
}

func projectsRoot() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, "Projects"), nil
}

func getWorkspaces(root string) ([]Workspace, error) {
	return scanDir(root, root)
}

func scanDir(root string, dirAbsPath string) ([]Workspace, error) {
	entries, err := os.ReadDir(dirAbsPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if filepath.Ext(entry.Name()) == workspaceExt {
			relDir, err := filepath.Rel(root, dirAbsPath)
			if err != nil {
				return nil, err
			}
			return []Workspace{{
				DirAbs:  dirAbsPath,
				FileAbs: filepath.Join(dirAbsPath, entry.Name()),
				RelDir:  relDir,
				Name:    entry.Name(),
			},
			}, nil
		}
	}

	var result []Workspace
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		localName := entry.Name()
		if slices.Contains(ignoreDirs, localName) {
			continue
		}

		subdirPath := filepath.Join(dirAbsPath, localName)
		subResult, err := scanDir(root, subdirPath)
		if err != nil {
			return nil, err
		}
		if len(subResult) > 0 {
			result = append(result, subResult...)
		}
	}

	return result, nil
}

func buildIndexByRelPath(workspaces []Workspace) map[string]Workspace {
	index := make(map[string]Workspace, len(workspaces))
	for _, ws := range workspaces {
		index[ws.RelDir] = ws
	}

	return index
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vspace <relative-directory>")
		return
	}
	target := os.Args[1]

	root, err := projectsRoot()
	if err != nil {
		fmt.Println("Error determining projects root:", err)
		return
	}
	fmt.Printf("Searching workspaces in %s\n", root)

	workspaces, err := getWorkspaces(root)
	if err != nil {
		fmt.Println("Error while scanning workspaces:", err)
		return
	}

	if len(workspaces) == 0 {
		fmt.Println("No workspaces found.")
		return
	}

	index := buildIndexByRelPath(workspaces)
	ws, ok := index[target]
	if !ok {
		fmt.Printf("Workspace %q not found.\n", target)
		fmt.Println("Available keys:")
		for _, w := range workspaces {
			fmt.Println(" -", w.RelDir)
		}
		return
	}

	fmt.Printf("Opening workspace: %s\n", ws.FileAbs)

	cmd := exec.Command("code", "-r", ws.FileAbs)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error running code:", err)
		return
	}

}
